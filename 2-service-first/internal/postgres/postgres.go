package postgres

import (
	"context"

	"github.com/golang-school/evolution/2-service-first/internal/model"
	"github.com/google/uuid"
)

type Config struct {
	User     string `envconfig:"POSTGRES_USER"     required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	Port     string `envconfig:"POSTGRES_PORT"     required:"true"`
	Host     string `envconfig:"POSTGRES_HOST"     required:"true"`
	DBName   string `envconfig:"POSTGRES_DB_NAME"  required:"true"`
}

type Pool struct{}

func New(ctx context.Context, c Config) (*Pool, error) {
	// Делаем настройки подключения и пингуем БД на доступность
	return &Pool{}, nil
}

func (p *Pool) CreateProfile(ctx context.Context, profile model.Profile) error {
	return nil
}

func (p *Pool) GetProfile(ctx context.Context, id uuid.UUID) (model.Profile, error) {
	return model.Profile{}, nil
}

func (p *Pool) Close() {
	// Shutdown
}
