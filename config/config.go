package config

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Bot TelegramBot
}

type TelegramBot struct {
	Token string
}

func MustLoad() (*Config, error) {
	var cfg Config

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	cfg.Bot.Token = telegramBotToken

	return &cfg, nil
}
