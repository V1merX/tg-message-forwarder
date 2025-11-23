package app

import (
	"context"
	"log/slog"
	"os"

	"github.com/V1merX/tg-message-forwarder/config"
	"github.com/V1merX/tg-message-forwarder/internal/api/telegram"
	kfk "github.com/V1merX/tg-message-forwarder/internal/repository/kafka"
	"github.com/segmentio/kafka-go"
)

type diContainer struct {
	kafkaConn *kafka.Conn
	bot       *telegram.Bot
	log       *slog.Logger
	config    *config.Config
}

func NewDIContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) InitKafka(ctx context.Context) (*kafka.Conn, error) {
	var err error
	if d.kafkaConn == nil {
		d.kafkaConn, err = kfk.Open(ctx)
		if err != nil {
			return nil, err
		}
	}

	return d.kafkaConn, nil
}

func (d *diContainer) InitConfig() (*config.Config, error) {
	var err error
	if d.config == nil {
		d.config, err = config.MustLoad()
		if err != nil {
			return nil, err
		}
	}

	return d.config, nil
}

func (d *diContainer) InitLogger() (*slog.Logger, error) {
	if d.log == nil {
		d.log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	}

	return d.log, nil
}

func (d *diContainer) InitBot() (*telegram.Bot, error) {
	if d.bot == nil {
		logger, err := d.InitLogger()
		if err != nil {
			return nil, err
		}
		config, err := d.InitConfig()
		if err != nil {
			return nil, err
		}
		d.bot = telegram.NewBot(logger, config.Bot.Token)
	}

	return d.bot, nil
}
