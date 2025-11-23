package telegram

import (
	"context"
	"log/slog"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

type Bot struct {
	token string
	log   *slog.Logger
}

func NewBot(log *slog.Logger, token string) *Bot {
	return &Bot{
		log:   log,
		token: token,
	}
}

func (b *Bot) Start(ctx context.Context) error {
	bot, err := telego.NewBot(b.token, telego.WithDefaultDebugLogger())
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

	bh.Handle(func(ctx *th.Context, update telego.Update) error {
		stdCtx := ctx.Context()

		_, err := ctx.Bot().SendMessage(stdCtx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			"Unknown command, use /start",
		))
		return err
	}, th.AnyCommand())

	b.log.Debug("Start handling updates")
	if err := bh.Start(); err != nil {
		return err
	}

	return nil
}
