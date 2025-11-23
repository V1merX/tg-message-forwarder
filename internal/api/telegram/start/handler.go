package start

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/V1merX/tg-message-forwarder/internal/entity"
	"github.com/V1merX/tg-message-forwarder/internal/usecase"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

type Handler struct {
	svc usecase.MessageService
	log *slog.Logger
}

func NewHandler(log *slog.Logger, svc usecase.MessageService) *Handler {
	return &Handler{
		log: log,
		svc: svc,
	}
}

func (h *Handler) GetMessage(ctx *th.Context, update telego.Update) error {
	stdCtx := ctx.Context()

	err := h.svc.SendMessage(ctx, update.Message.From.ID, update.Message.From.Username, update.Message.Text)
	if err != nil {
		return err
	}

	_, err = ctx.Bot().SendMessage(stdCtx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		"Message has been successfully sent",
	))
	return err
}

func (h *Handler) MessagePulling(ctx context.Context, bot *telego.Bot, adminID int64) error {
	h.log.Info("Start message pulling...")

	messagesCh := make(chan []byte, 1000)
	go func() {
		if err := h.svc.GetMessages(ctx, messagesCh); err != nil {
			h.log.Error("Failed to get message", slog.Any("error", err))
		}
	}()

	for msg := range messagesCh {
		var message entity.Message
		err := json.Unmarshal(msg, &message)
		if err != nil {
			h.log.Error("Failed to unmarshal message", slog.Any("err", err))
		}

		_, err = bot.SendMessage(ctx, tu.Message(
			tu.ID(adminID),
			entity.FormatMessage(message),
		).WithParseMode(telego.ModeMarkdownV2))
		if err != nil {
			h.log.Error("Failed to send message log to admin", slog.Any("error", err))
		}
	}
	return nil
}
