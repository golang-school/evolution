package profile_service

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-school/evolution/2-service-first/internal/kafka_produce"

	"github.com/golang-school/evolution/2-service-first/internal/model"
	"github.com/google/uuid"
)

func (p *Profile) CreateProfile(ctx context.Context, name string, age int, email string) (uuid.UUID, error) {
	// Проверяем в Redis ключу идемпотентности
	if p.redis.IsExists(ctx, name+email) {
		return uuid.Nil, ErrAlreadyExists
	}

	// Создаём профиль
	profile := model.Profile{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		Name:      name,
		Age:       age,
		Email:     email,
	}

	// Валидируем
	err := profile.Validate()
	if err != nil {
		return uuid.Nil, fmt.Errorf("validate profile: %w", err)
	}

	// Сохраняем в БД
	err = p.postgres.CreateProfile(ctx, profile)
	if err != nil {
		return uuid.Nil, fmt.Errorf("create profile in postgres: %w", err)
	}

	// Отправляем в Kafka событие создания профиля
	err = p.kafka.Produce(ctx, kafka_produce.Message{})
	if err != nil {
		return uuid.Nil, fmt.Errorf("kafka produce: %w", err)
	}

	return profile.ID, nil
}
