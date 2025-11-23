package message

import (
	"context"
	"encoding/json"
	"time"

	"github.com/V1merX/tg-message-forwarder/internal/entity"
	"github.com/V1merX/tg-message-forwarder/internal/repository"
	"github.com/google/uuid"
)

type Service struct {
	repo repository.MessageRepository
}

func NewService(repo repository.MessageRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) SendMessage(ctx context.Context, telegramID int64, telegramUserName string, text string) error {
	message := entity.Message{
		ID:               uuid.New(),
		TelegramID:       telegramID,
		TelegramUserName: telegramUserName,
		Text:             text,
		CreatedAt:        time.Now(),
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return s.repo.SendMessage(ctx, messageBytes)
}

func (s *Service) GetMessages(ctx context.Context, messages chan<- []byte) error {
	return s.repo.GetMessages(ctx, messages)
}
