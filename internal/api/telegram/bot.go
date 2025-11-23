package telegram

import (
	"context"
	"log/slog"

	"github.com/V1merX/tg-message-forwarder/internal/api"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type Bot struct {
	token string
	log   *slog.Logger
	mh    api.StartAPI
}

func NewBot(log *slog.Logger, messageHandler api.StartAPI, token string) *Bot {
	return &Bot{
		mh:    messageHandler,
		log:   log,
		token: token,
	}
}

func (b *Bot) Start(ctx context.Context) error {
	bot, err := telego.NewBot(b.token, telego.WithDefaultLogger(false, false))
	if err != nil {
		b.log.Error("Failed to start telegram bot", slog.Any("err", err))
		return err
	}

	updates, err := bot.UpdatesViaLongPolling(ctx, nil)
	if err != nil {
		b.log.Error("Failed to get updates via long polling", slog.Any("err", err))
		return err
	}

	bh, err := th.NewBotHandler(bot, updates)
	if err != nil {
		b.log.Error("Failed to create bot handler", slog.Any("err", err))
		return err
	}

	defer func() {
		if err := bh.Stop(); err != nil {
			b.log.Error("Failed to stop bot handler", slog.Any("err", err))
			return
		}
	}()

	bh.Handle(b.mh.GetMessage, th.AnyCommand())

	b.log.Debug("Start handling updates")
	if err := bh.Start(); err != nil {
		return err
	}

	return nil
}
