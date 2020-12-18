package configs

import (
	"github.com/beop13/notification-service/logger"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

var Cfg Config

type Config struct {
	ServiceSettings struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"service_settings"`
	EmailSettings struct {
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	} `yaml:"email_settings"`
	TelegramSettings struct {
		BotToken string `yaml:"bot_token"`
		ChatId   string `yaml:"chat_id"`
	} `yaml:"telegram_settings"`
}

func init() {
	args := os.Args[0:]
	if len(args) < 2 {
		panic("\nMissing argument filePath.\n\nUsage: go run main ./config.yml (relative or absolute path)")
	}

	filePath := args[1]

	logger.L.Printf("Creating config from yaml at path %s", filePath)

	err := cleanenv.ReadConfig(filePath, &Cfg)
	if err != nil {
		logger.L.Print("Failed to unmarshal yml")
		panic(err)
	}
}
