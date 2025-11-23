package kafka

import (
	"context"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

const messageTopic string = "messages"

type MessageRepository struct {
	conn *kafka.Conn
	log  *slog.Logger
}

func NewMessageRepository(conn *kafka.Conn, log *slog.Logger) *MessageRepository {
	msgRepo := &MessageRepository{
		conn: conn,
		log:  log,
	}

	if err := msgRepo.createMessageTopic(); err != nil {
		msgRepo.log.Error("Failed to create topic")
		panic(err)
	}

	return msgRepo
}

func (r MessageRepository) createMessageTopic() error {
	topicConfig := kafka.TopicConfig{
		Topic:             messageTopic,
		ReplicationFactor: 1,
		NumPartitions:     1,
	}

	err := r.conn.CreateTopics(topicConfig)
	if err != nil {
		return err
	}

	return err
}

func (r *MessageRepository) SendMessage(ctx context.Context) error {
	writer := kafka.Writer{
		Addr:         r.conn.RemoteAddr(),
		Topic:        messageTopic,
		RequiredAcks: -1,
	}

	err := writer.WriteMessages(ctx, kafka.Message{
		Value: []byte("new message"),
	})
	if err != nil {
		r.log.Error("Failed to send message to kafka", slog.Any("err", err))
		return err
	}

	r.log.Info("Successfull sending message to kafka")

	return nil
}
