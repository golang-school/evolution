package usecase

import (
	"context"
	"fmt"

	"github.com/golang-school/evolution/6-layers-ddd/internal/dto"

	"github.com/golang-school/evolution/6-layers-ddd/internal/adapter/kafka_produce"
	"github.com/golang-school/evolution/6-layers-ddd/internal/domain"
)

func (p *Profile) CreateProfile(ctx context.Context, input dto.CreateProfileInput) (dto.CreateProfileOutput, error) {
	var output dto.CreateProfileOutput

	// Проверяем в Redis ключу идемпотентности
	if p.redis.IsExists(ctx, input.Email) {
		return output, domain.ErrAlreadyExists
	}

	// Создаём профиль
	profile, err := domain.NewProfile(input.Name, input.Age, input.Email)
	if err != nil {
		return output, fmt.Errorf("new profile: %w", err)
	}

	// Сохраняем в БД
	err = p.postgres.CreateProfile(ctx, profile)
	if err != nil {
		return output, fmt.Errorf("create profile in postgres: %w", err)
	}

	// Отправляем в Kafka событие создания профиля
	err = p.kafka.Produce(ctx, kafka_produce.Message{})
	if err != nil {
		return output, fmt.Errorf("kafka produce: %w", err)
	}

	return dto.CreateProfileOutput{
		ID: profile.ID,
	}, nil
}
