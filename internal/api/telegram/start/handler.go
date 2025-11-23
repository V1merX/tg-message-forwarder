package start

import (
	"github.com/V1merX/tg-message-forwarder/internal/usecase"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

type Handler struct {
	svc usecase.MessageService
}

func NewHandler(svc usecase.MessageService) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) GetMessage(ctx *th.Context, update telego.Update) error {
	stdCtx := ctx.Context()

	err := h.svc.SendMessage(ctx)
	if err != nil {
		return err
	}

	_, err = ctx.Bot().SendMessage(stdCtx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		"Message has been successfully sent",
	))
	return err
}
