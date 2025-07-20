package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-school/evolution/7-layers-cqrs/config"
	"github.com/golang-school/evolution/7-layers-cqrs/internal/adapter/kafka_produce"
	"github.com/golang-school/evolution/7-layers-cqrs/internal/adapter/postgres"
	"github.com/golang-school/evolution/7-layers-cqrs/internal/adapter/redis"
	"github.com/golang-school/evolution/7-layers-cqrs/internal/controller/http"
	"github.com/golang-school/evolution/7-layers-cqrs/internal/controller/kafka_consume"
	"github.com/golang-school/evolution/7-layers-cqrs/internal/usecase"
	"github.com/golang-school/evolution/7-layers-cqrs/pkg/httpserver"
	"github.com/golang-school/evolution/7-layers-cqrs/pkg/logger"
)

// Main функция приложения
func main() {
	// Конфиг
	c, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	// Логгер
	logger.Init(c.Logger)

	// Билдим и запускаем приложение
	err = AppRun(context.Background(), c)
	if err != nil {
		panic(err)
	}
}

// Запускаем приложение
func AppRun(ctx context.Context, c config.Config) error {
	// Postgres master
	master, err := postgres.New(ctx, c.Postgres)
	if err != nil {
		return fmt.Errorf("postgres.New master: %w", err)
	}

	// Postgres replica
	replica, err := postgres.New(ctx, c.Postgres)
	if err != nil {
		return fmt.Errorf("postgres.New replica:  %w", err)
	}

	// Redis
	redisClient, err := redis.New(c.Redis)
	if err != nil {
		return fmt.Errorf("redis.New: %w", err)
	}

	// Kafka producer
	kafkaProducer := kafka_produce.NewProducer(c.KafkaProducer)

	// Usecase
	profileUsecase := usecase.NewProfile(master, replica, kafkaProducer, redisClient)

	// Kafka consumer
	kafkaConsumer := kafka_consume.New(c.KafkaConsumer, profileUsecase)

	// HTTP сервер
	router := http.Router(profileUsecase)
	httpServer := httpserver.New(router, c.HTTP)

	// Приложение запущено и готово к работе

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	<-sig // ждём здесь сигнала (Ctrl+C или SIGTERM)

	// Закрываем ресурсы
	kafkaConsumer.Close()
	httpServer.Close()
	redisClient.Close()
	kafkaProducer.Close()
	master.Close()
	replica.Close()

	return nil
}
