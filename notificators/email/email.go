package email

import (
	"fmt"
	"github.com/beop13/notification-service/logger"
	"github.com/beop13/notification-service/notificators/model"
	"net/smtp"
	"strings"
)

type Email struct {
	Login    string
	Password string
	Host     string
	Port     string
}

func (e Email) SendNotification(message model.Message) error {
	logger.L.Printf("Sending notification with message %v by email", message)
	from := e.Login
	pass := e.Password
	//body := fmt.Sprintf("<html><body><h1>%s</h1></body></html>", message.Body)
	toHeader := strings.Join(message.To, ",")

	header := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n", e.Login, toHeader, message.Subject)
	//mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(fmt.Sprintf("%s%s", header, message.Body))

	err := smtp.SendMail(e.Host+":"+e.Port,
		smtp.PlainAuth("", from, pass, e.Host),
		from, message.To, []byte(msg))

	if err != nil {
		logger.L.Printf("%s", err)
		return err
	}

	logger.L.Print("Email sent")
	return nil
}
