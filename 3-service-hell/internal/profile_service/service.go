package profile_service

import (
	"context"
	"errors"

	"github.com/golang-school/evolution/3-service-hell/internal/kafka_produce"
	"github.com/golang-school/evolution/3-service-hell/internal/model"
	"github.com/golang-school/evolution/3-service-hell/internal/postgres"
	"github.com/golang-school/evolution/3-service-hell/internal/redis"
	"github.com/google/uuid"
)

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrUUIDInvalid   = errors.New("invalid UUID format")
)

type Redis interface {
	IsExists(ctx context.Context, idempotencyKey string) bool
}

type Postgres interface {
	CreateProfile(ctx context.Context, profile model.Profile) error
	GetProfile(ctx context.Context, id uuid.UUID) (model.Profile, error)
}

type MessageService interface {
	SendMessage(ctx context.Context, msgs ...kafka_produce.Message) error
}

type Profile struct {
	postgres       Postgres
	redis          Redis
	messageService MessageService
}

func NewProfile(postgres *postgres.Pool, messageService MessageService, redis *redis.Client) *Profile {
	return &Profile{
		postgres:       postgres,
		messageService: messageService,
		redis:          redis,
	}
}
