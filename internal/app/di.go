package app

import (
	"log/slog"
	"os"

	"github.com/V1merX/tg-message-forwarder/config"
	"github.com/V1merX/tg-message-forwarder/internal/api"
	"github.com/V1merX/tg-message-forwarder/internal/api/telegram"
	"github.com/V1merX/tg-message-forwarder/internal/api/telegram/start"
	"github.com/V1merX/tg-message-forwarder/internal/repository"
	kfk "github.com/V1merX/tg-message-forwarder/internal/repository/kafka"
	"github.com/V1merX/tg-message-forwarder/internal/usecase"
	"github.com/V1merX/tg-message-forwarder/internal/usecase/message"
	"github.com/segmentio/kafka-go"
)

type diContainer struct {
	kafkaConn      *kafka.Conn
	bot            *telegram.Bot
	log            *slog.Logger
	config         *config.Config
	messageSvc     usecase.MessageService
	messageHandler api.StartAPI
	messageRepo    repository.MessageRepository
}

func NewDIContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) KafkaConnection() *kafka.Conn {
	var err error
	if d.kafkaConn == nil {
		d.kafkaConn, err = kfk.Open()
		if err != nil {
			panic(err)
		}
	}

	return d.kafkaConn
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

func (d *diContainer) Logger() *slog.Logger {
	if d.log == nil {
		d.log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	}

	return d.log
}

func (d *diContainer) InitBot() (*telegram.Bot, error) {
	if d.bot == nil {
		config, err := d.InitConfig()
		if err != nil {
			return nil, err
		}
		d.bot = telegram.NewBot(d.Logger(), d.MessageHandler(), config.Bot.Token)
	}

	return d.bot, nil
}

func (d *diContainer) MessageHandler() api.StartAPI {
	if d.messageHandler == nil {
		d.messageHandler = start.NewHandler(d.MessageService())
	}

	return d.messageHandler
}

func (d *diContainer) MessageService() usecase.MessageService {
	if d.messageSvc == nil {
		d.messageSvc = message.NewService(d.MessageRepository())
	}

	return d.messageSvc
}

func (d *diContainer) MessageRepository() repository.MessageRepository {
	if d.messageRepo == nil {
		d.messageRepo = kfk.NewMessageRepository(d.KafkaConnection(), d.Logger())
	}

	return d.messageRepo
}
