package message

import (
	"context"

	"github.com/V1merX/tg-message-forwarder/internal/repository"
)

type Service struct {
	repo repository.MessageRepository
}

func NewService(repo repository.MessageRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) SendMessage(ctx context.Context) error {
	return s.repo.SendMessage(ctx)
}
