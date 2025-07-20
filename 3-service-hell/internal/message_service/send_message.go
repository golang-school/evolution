package message_service

import (
	"context"
	"fmt"

	"github.com/golang-school/evolution/3-service-hell/internal/kafka_produce"
)

func (m *MessageService) SendMessage(ctx context.Context, msgs ...kafka_produce.Message) error {
	// Что-то делаем и...

	// отправляем в Kafka событие создания профиля
	err := m.kafka.Produce(ctx, msgs...)
	if err != nil {
		return fmt.Errorf("kafka produce: %w", err)
	}

	return nil
}
