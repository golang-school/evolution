package message_service

import (
	"context"

	"github.com/golang-school/evolution/3-service-hell/internal/kafka_produce"
)

type Kafka interface {
	Produce(ctx context.Context, msgs ...kafka_produce.Message) error
}

type MessageService struct {
	kafka Kafka
}

func NewMessageService(kafka *kafka_produce.Producer) *MessageService {
	return &MessageService{
		kafka: kafka,
	}
}
