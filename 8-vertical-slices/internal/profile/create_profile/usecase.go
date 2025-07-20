package create_profile

import (
	"context"
	"fmt"

	"github.com/golang-school/evolution/8-vertical-slices/internal/adapter/kafka_produce"
	"github.com/golang-school/evolution/8-vertical-slices/internal/adapter/postgres"
	"github.com/golang-school/evolution/8-vertical-slices/internal/adapter/redis"
	"github.com/golang-school/evolution/8-vertical-slices/internal/domain"
)

type Redis interface {
	IsExists(ctx context.Context, idempotencyKey string) bool
}

type Kafka interface {
	Produce(ctx context.Context, msgs ...kafka_produce.Message) error
}

type Postgres interface {
	CreateProfile(ctx context.Context, profile domain.Profile) error
}

type Usecase struct {
	postgres Postgres
	kafka    Kafka
	redis    Redis
}

func New(postgres *postgres.Pool, kafka *kafka_produce.Producer, redis *redis.Client) *Usecase {
	uc := &Usecase{
		postgres: postgres,
		kafka:    kafka,
		redis:    redis,
	}

	usecase = uc // global for handlers

	return uc
}

func (u *Usecase) CreateProfile(ctx context.Context, input Input) (Output, error) {
	// Проверяем в Redis ключу идемпотентности
	if u.redis.IsExists(ctx, input.Email) {
		return Output{}, domain.ErrAlreadyExists
	}

	// Создаём профиль
	profile, err := domain.NewProfile(input.Name, input.Age, input.Email)
	if err != nil {
		return Output{}, fmt.Errorf("new profile: %w", err)
	}

	// Сохраняем в БД
	err = u.postgres.CreateProfile(ctx, profile)
	if err != nil {
		return Output{}, fmt.Errorf("create profile in postgres: %w", err)
	}

	// Отправляем в Kafka событие создания профиля
	err = u.kafka.Produce(ctx, kafka_produce.Message{})
	if err != nil {
		return Output{}, fmt.Errorf("kafka produce: %w", err)
	}

	return Output{
		ID: profile.ID,
	}, nil
}
