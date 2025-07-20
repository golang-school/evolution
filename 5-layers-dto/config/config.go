package config

import (
	"github.com/golang-school/evolution/5-layers-dto/internal/adapter/kafka_produce"
	"github.com/golang-school/evolution/5-layers-dto/internal/adapter/postgres"
	"github.com/golang-school/evolution/5-layers-dto/internal/adapter/redis"
	"github.com/golang-school/evolution/5-layers-dto/internal/controller/kafka_consume"
	"github.com/golang-school/evolution/5-layers-dto/pkg/httpserver"
	"github.com/golang-school/evolution/5-layers-dto/pkg/logger"
	"github.com/golang-school/evolution/5-layers-dto/pkg/otel"
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
