package kafka_produce

import (
	"context"
)

type Config struct {
	Addr  []string `envconfig:"KAFKA_WRITER_ADDR"  required:"true"`
	Topic string   `envconfig:"KAFKA_WRITER_TOPIC"`
}

type Producer struct {
	config Config
}

func NewProducer(c Config) *Producer {
	// Настройка топика, адресов и других параметров
	return &Producer{}
}

type Message struct{}

func (p *Producer) Produce(ctx context.Context, msgs ...Message) error {
	// Отправляем данные в Кафку
	return nil
}

func (p *Producer) Close() {
	// Shutdown
}
