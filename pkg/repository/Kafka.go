package repository

import (
	"context"
	"encoding/json"
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
	bookBytes, errMarsh := json.Marshal(book)
	if errMarsh != nil {
		return fmt.Errorf("invalid data, %s", errMarsh)
	}
	message := kafka.Message{
		Key:   []byte("book"),
		Value: bookBytes,
		Time:  time.Now(),
	}
	_, err := k.writerConn.WriteMessages(message)

	if err != nil {
		return fmt.Errorf("error while sending book to antother ms, %s", err)
	}
	return nil
}

func (k *KafkaConn) GetBookKafka(ctx context.Context, bookName string) (models.Book, error) {
	k.writerConn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	bookBytes, errMarsh := json.Marshal(bookName)
	if errMarsh != nil {
		return models.Book{}, fmt.Errorf("invalid data, %s", errMarsh)
	}
	message := kafka.Message{
		Key:   []byte("command"),
		Value: bookBytes,
		Time:  time.Now(),
	}
	_, err := k.writerConn.WriteMessages(message)
	if err != nil {
		return models.Book{}, fmt.Errorf("error while sending book to antother ms, %s", err)
	}
}
