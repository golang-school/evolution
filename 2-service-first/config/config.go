package config

import (
	"github.com/golang-school/evolution/2-service-first/internal/kafka_consume"
	"github.com/golang-school/evolution/2-service-first/internal/kafka_produce"
	"github.com/golang-school/evolution/2-service-first/internal/postgres"
	"github.com/golang-school/evolution/2-service-first/internal/redis"
	"github.com/golang-school/evolution/2-service-first/pkg/httpserver"
	"github.com/golang-school/evolution/2-service-first/pkg/logger"
	"github.com/golang-school/evolution/2-service-first/pkg/otel"
)

// Конфиг приложения
type App struct {
	Name    string `envconfig:"APP_NAME"    required:"true"`
	Version string `envconfig:"APP_VERSION" required:"true"`
}

type Config struct {
	App           App
	HTTP          httpserver.Config
	Logger        logger.Config
	OTEL          otel.Config
	Postgres      postgres.Config
	Redis         redis.Config
	KafkaProducer kafka_produce.Config
	KafkaConsumer kafka_consume.Config
}

func InitConfig() (Config, error) {
	// Парсим и валидируем конфиг
	return Config{}, nil
}
