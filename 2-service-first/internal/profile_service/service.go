package profile_service

import (
	"context"
	"errors"

	"github.com/golang-school/evolution/2-service-first/internal/kafka_produce"
	"github.com/golang-school/evolution/2-service-first/internal/model"
	"github.com/golang-school/evolution/2-service-first/internal/postgres"
	"github.com/golang-school/evolution/2-service-first/internal/redis"
	"github.com/google/uuid"
)

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrUUIDInvalid   = errors.New("invalid UUID format")
)

type Redis interface {
	IsExists(ctx context.Context, idempotencyKey string) bool
}

type Kafka interface {
	Produce(ctx context.Context, msgs ...kafka_produce.Message) error
}

type Postgres interface {
	CreateProfile(ctx context.Context, profile model.Profile) error
	GetProfile(ctx context.Context, id uuid.UUID) (model.Profile, error)
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
