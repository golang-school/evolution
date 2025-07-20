package kafka_consume

import "github.com/golang-school/evolution/3-service-hell/internal/profile_service"

type Config struct {
	Addr  []string `envconfig:"KAFKA_CONSUMER_ADDR"     required:"true"`
	Topic string   `default:"mnepryakhin-my-app-topic"  envconfig:"KAFKA_CONSUMER_TOPIC"`
	Group string   `default:"mnepryakhin-my-app-group"  envconfig:"KAFKA_CONSUMER_GROUP"`
}

// Кафка консьюмер
type Consumer struct{}

func New(cfg Config, profile *profile_service.Profile) *Consumer {
	// Настройки
	// Запуск горутины для приёма сообщений и передачи их в сервис
	return &Consumer{}
}

func (c *Consumer) Close() {
	// Shutdown
}
