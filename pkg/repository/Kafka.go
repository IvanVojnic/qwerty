package repository

import (
	"context"
	"fmt"
	"time"

	"EFpractic2/models"

	"github.com/segmentio/kafka-go"
)

// KafkaConn struct with kafka connection
type KafkaConn struct {
	writerConn *kafka.Conn
}

// NewKafkaConn used to init kafka connection
func NewKafkaConn(conn *kafka.Conn) *KafkaConn {
	return &KafkaConn{writerConn: conn}
}

// CreateBookKafka used to send book to another ms to create book
func (k *KafkaConn) CreateBookKafka(ctx context.Context, book *models.Book) error {
	k.writerConn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	a, err := k.writerConn.WriteMessages(
		kafka.Message{Value: []byte("hello world!")},
	)

	if err != nil {
		return fmt.Errorf("error while sending book to antother ms, %s", err)
	}
	fmt.Print(a)
	return nil
}
