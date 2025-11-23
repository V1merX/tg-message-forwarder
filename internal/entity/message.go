package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID               uuid.UUID `json:"id"`
	TelegramID       int64     `json:"telegram_id"`
	TelegramUserName string    `json:"telegram_username"`
	Text             string    `json:"text"`
	CreatedAt        time.Time `json:"created_at"`
}

func FormatMessage(message Message) string {
	str := fmt.Sprintf(`
{
		"id": %v
		"telegram_id": %d
		"telegram_username": %s
		"text": %s
		"created_at": %v
}
	`, message.ID, message.TelegramID, message.TelegramUserName, message.Text, message.CreatedAt)

	return fmt.Sprintf("```json%s```", str)
}
