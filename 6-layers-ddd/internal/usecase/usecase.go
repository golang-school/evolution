package usecase

import (
	"context"

	"github.com/golang-school/evolution/6-layers-ddd/internal/adapter/kafka_produce"
	"github.com/golang-school/evolution/6-layers-ddd/internal/adapter/postgres"
	"github.com/golang-school/evolution/6-layers-ddd/internal/adapter/redis"

	"github.com/golang-school/evolution/6-layers-ddd/internal/domain"
	"github.com/google/uuid"
)

type Redis interface {
	IsExists(ctx context.Context, idempotencyKey string) bool
}

type Kafka interface {
	Produce(ctx context.Context, msgs ...kafka_produce.Message) error
}

type Postgres interface {
	CreateProfile(ctx context.Context, profile domain.Profile) error
	GetProfile(ctx context.Context, id uuid.UUID) (domain.Profile, error)
}

type Profile struct {
	postgres Postgres
	kafka    Kafka
	redis    Redis
}

func NewProfile(postgres *postgres.Pool, kafka *kafka_produce.Producer, redis *redis.Client) *Profile {
	return &Profile{
		postgres: postgres,
		kafka:    kafka,
		redis:    redis,
	}
}
