package repository

import (
	"context"
)

type MessageRepository interface {
	SendMessage(ctx context.Context, message []byte) error
	GetMessages(ctx context.Context, messages chan<- []byte) error
}
