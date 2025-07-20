package server

import (
	"errors"

	"github.com/golang-school/evolution/1-handler-first/internal/kafka_produce"
	"github.com/golang-school/evolution/1-handler-first/internal/postgres"
	"github.com/golang-school/evolution/1-handler-first/internal/redis"
)

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrUUIDInvalid   = errors.New("invalid UUID format")
)

type Handlers struct {
	postgres *postgres.Pool
	kafka    *kafka_produce.Producer
	redis    *redis.Client
}

func NewHandlers(postgres *postgres.Pool, kafka *kafka_produce.Producer, redis *redis.Client) *Handlers {
	return &Handlers{
		postgres: postgres,
		kafka:    kafka,
		redis:    redis,
	}
}
