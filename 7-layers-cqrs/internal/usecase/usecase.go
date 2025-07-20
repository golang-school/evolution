package usecase

import (
	"context"

	"github.com/golang-school/evolution/7-layers-cqrs/internal/adapter/kafka_produce"
	"github.com/golang-school/evolution/7-layers-cqrs/internal/adapter/postgres"
	"github.com/golang-school/evolution/7-layers-cqrs/internal/adapter/redis"
	"github.com/golang-school/evolution/7-layers-cqrs/internal/domain"
	"github.com/google/uuid"
)

type Redis interface {
	IsExists(ctx context.Context, idempotencyKey string) bool
}

type Kafka interface {
	Produce(ctx context.Context, msgs ...kafka_produce.Message) error
}

type PostgresMaster interface {
	CreateProfile(ctx context.Context, profile domain.Profile) error
}

type PostgresReplica interface {
	GetProfile(ctx context.Context, id uuid.UUID) (domain.Profile, error)
}

type Profile struct {
	postgresMaster  PostgresMaster
	postgresReplica PostgresReplica
	kafka           Kafka
	redis           Redis
}

func NewProfile(master, replica *postgres.Pool, kafka *kafka_produce.Producer, redis *redis.Client) *Profile {
	return &Profile{
		postgresMaster:  master,
		postgresReplica: replica,
		kafka:           kafka,
		redis:           redis,
	}
}
