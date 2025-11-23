package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func Open(ctx context.Context) (*kafka.Conn, error) {
	conn, err := kafka.DialContext(ctx, "tcp", "localhost:9092")
	if err != nil {
		return nil, err
	}

	return conn, nil
}
