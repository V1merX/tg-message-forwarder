package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Bot TelegramBot
}

type TelegramBot struct {
	Token           string
	TelegramAdminID int64
}

func MustLoad() (*Config, error) {
	var cfg Config

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	telegramAdminID := os.Getenv("TELEGRAM_ADMIN_ID")

	cfg.Bot.Token = telegramBotToken

	telegramAdminIDInt, err := strconv.Atoi(telegramAdminID)
	if err != nil {
		return nil, err
	}

	cfg.Bot.TelegramAdminID = int64(telegramAdminIDInt)

	return &cfg, nil
}
