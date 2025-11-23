package usecase

import "context"

type MessageService interface {
	SendMessage(ctx context.Context) error
}
