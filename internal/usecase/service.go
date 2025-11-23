package usecase

import "context"

type MessageService interface {
	SendMessage(ctx context.Context, telegramID int64, telegramUserName, message string) error
	GetMessages(ctx context.Context, messages chan<- []byte) error
}
