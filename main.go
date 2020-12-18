package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beop13/notification-service/configs"
	"github.com/beop13/notification-service/logger"
	"github.com/beop13/notification-service/notificators"
	"github.com/beop13/notification-service/notificators/model"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
	"time"
)


/* TODO
tls
 */
func main() {
	logger.L.Print("Welcome to notification-service")

	r := httprouter.New()
	r.POST("/notification", NotificationHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    configs.Cfg.ServiceSettings.Host + ":" + configs.Cfg.ServiceSettings.Port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.L.Printf("Running server. Listening %s:%s\n", configs.Cfg.ServiceSettings.Host, configs.Cfg.ServiceSettings.Port)
	logger.L.Fatal(srv.ListenAndServe())

}

func NotificationHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logger.L.Printf("Notification handler: %v", r)

	sendVia := r.URL.Query().Get("send_via")
	to := r.URL.Query().Get("to")
	subject := r.URL.Query().Get("subject")
	messageText := r.URL.Query().Get("message_text")

	toFormatted := strings.Split(to, ",")

	m := model.Message{
		Body:    messageText,
		Subject: subject,
		To:      toFormatted,
	}

	logger.L.Printf("Send via: %v \n\nMessage is: %+v\n", sendVia, m)

	sendViaMap := make(map[string]bool)

	if sendVia != "" {
		logger.L.Print("Parsing send_via to map")
		sendViaArr := strings.Split(sendVia, ",")
		if len(sendViaArr) != 0 {
			for _, v := range sendViaArr {
				sendViaMap[v] = true
			}
		}
	}

	err := validateRequest(sendViaMap, m)
	if err != nil {
		logger.L.Print("Validating failed")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	errs := make([]map[string]string, 0)
	if _, ok := sendViaMap["email"]; ok {
		err := notificators.NM.Notificators["email"].SendNotification(m)
		if err != nil {
			errs = append(errs, map[string]string{
				"send_by": "email",
				"error":   fmt.Sprint(err),
			})
		}
	}
	if _, ok := sendViaMap["telegram"]; ok {
		err := notificators.NM.Notificators["telegram"].SendNotification(m)
		if err != nil {
			errs = append(errs, map[string]string{
				"send_by": "telegram",
				"error":   fmt.Sprint(err),
			})
		}
	}

	if len(errs) > 0 {
		w.WriteHeader(http.StatusInternalServerError)
		b, _ := json.Marshal(errs)
		fmt.Fprintf(w, "%s", b)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func validateRequest(sendBy map[string]bool, m model.Message) error {
	logger.L.Print("Validating fields")
	if len(sendBy) == 0 {
		return errors.New("you have to fill parameter [send_via] with at least one value. [telegram, email]")
	}

	if m.Subject == "" && m.Body == "" {
		return errors.New("you have to fill at least one parameter [message_text], [subject]")
	}

	if _, ok := sendBy["email"]; len(m.To) == 0 && ok {
		return errors.New("you have to fill parameter [to]")
	}
	logger.L.Print("Validation")
	return nil
}
