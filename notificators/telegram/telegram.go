package telegram

import (
	"fmt"
	"github.com/beop13/notification-service/logger"
	"github.com/beop13/notification-service/notificators/model"
	"net/http"
	"net/url"
)

type Telegram struct {
	BotToken string
	ChatId   string
}

func (t Telegram) SendNotification(message model.Message) error {
	logger.L.Printf("Sending notification with message %v by telegram", message)
	base, err := url.Parse(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.BotToken))
	if err != nil {
		logger.L.Printf("%s", err)
		return err
	}

	params := url.Values{}
	params.Add("chat_id", t.ChatId)
	params.Add("text", message.Subject+"\n\n"+message.Body)
	base.RawQuery = params.Encode()

	res, err := http.Get(base.String())
	if err != nil {
		logger.L.Printf("%s", err)
		return err
	}

	logger.L.Print("Message send", res)
	return nil
}
