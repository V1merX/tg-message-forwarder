package kafka

import (
	"github.com/segmentio/kafka-go"
)

func Open() (*kafka.Conn, error) {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		return nil, err
	}

	return conn, nil
}
