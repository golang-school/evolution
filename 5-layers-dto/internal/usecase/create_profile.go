package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-school/evolution/5-layers-dto/internal/dto"

	"github.com/golang-school/evolution/5-layers-dto/internal/adapter/kafka_produce"
	"github.com/golang-school/evolution/5-layers-dto/internal/domain"
	"github.com/google/uuid"
)

func (p *Profile) CreateProfile(ctx context.Context, input dto.CreateProfileInput) (dto.CreateProfileOutput, error) {
	var output dto.CreateProfileOutput

	// Проверяем в Redis ключу идемпотентности
	if p.redis.IsExists(ctx, input.Email) {
		return output, domain.ErrAlreadyExists
	}

	// Создаём профиль
	profile := domain.Profile{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		Name:      input.Name,
		Age:       input.Age,
		Email:     input.Email,
	}

	// Валидируем
	err := profile.Validate()
	if err != nil {
		return output, fmt.Errorf("validate profile: %w", err)
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
