package kafka

import (
	"context"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

const messageTopic string = "messages"

type MessageRepository struct {
	conn        *kafka.Conn
	log         *slog.Logger
	kafkaWriter *kafka.Writer
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

	msgRepo.kafkaWriter = &kafka.Writer{
		Addr:         msgRepo.conn.RemoteAddr(),
		Topic:        messageTopic,
		RequiredAcks: -1,
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

func (r *MessageRepository) SendMessage(ctx context.Context, message []byte) error {
	err := r.kafkaWriter.WriteMessages(ctx, kafka.Message{
		Value: message,
	})
	if err != nil {
		r.log.Error("Failed to send message to kafka", slog.Any("err", err))
		return err
	}

	r.log.Info("Successfull sending message to kafka")

	return nil
}

func (r *MessageRepository) GetMessages(ctx context.Context, messages chan<- []byte) error {
	defer close(messages)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{r.conn.RemoteAddr().String()},
		Topic:   messageTopic,
		GroupID: "message-puller",
	})
	defer func() {
		if err := reader.Close(); err != nil {
			r.log.Error("Failed to close reader", slog.Any("err", err))
			return
		}
	}()

	r.log.Info("Start consuming messages...")
	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			r.log.Error("Failed to read message from 'message' topic", slog.Any("err", err))
			return err
		}

		messages <- m.Value
	}
}
