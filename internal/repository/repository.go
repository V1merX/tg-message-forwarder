package repository

import "context"

type MessageRepository interface {
	SendMessage(ctx context.Context) error
}
