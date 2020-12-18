package notificators

import (
	"github.com/beop13/notification-service/configs"
	"github.com/beop13/notification-service/logger"
	"github.com/beop13/notification-service/notificators/email"
	"github.com/beop13/notification-service/notificators/model"
	"github.com/beop13/notification-service/notificators/telegram"
)

var NM NotificatorManager

type NotificatorManager struct {
	Notificators map[string]Notificator
}

func init() {
	NM.Notificators = make(map[string]Notificator)

	NM.Notificators["email"] = email.Email{
		Login:    configs.Cfg.EmailSettings.Login,
		Password: configs.Cfg.EmailSettings.Password,
		Host:     configs.Cfg.EmailSettings.Host,
		Port:     configs.Cfg.EmailSettings.Port,
	}

	NM.Notificators["telegram"] = telegram.Telegram{
		BotToken: configs.Cfg.TelegramSettings.BotToken,
		ChatId:   configs.Cfg.TelegramSettings.ChatId,
	}

	logger.L.Println("Initialized notificators:")
	logger.L.Println(NM.Notificators)
}

type Notificator interface {
	SendNotification(message model.Message) error
}
