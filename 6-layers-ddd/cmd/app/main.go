package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-school/evolution/6-layers-ddd/config"
	"github.com/golang-school/evolution/6-layers-ddd/internal/adapter/kafka_produce"
	"github.com/golang-school/evolution/6-layers-ddd/internal/adapter/postgres"
	"github.com/golang-school/evolution/6-layers-ddd/internal/adapter/redis"
	"github.com/golang-school/evolution/6-layers-ddd/internal/controller/http"
	"github.com/golang-school/evolution/6-layers-ddd/internal/controller/kafka_consume"
	"github.com/golang-school/evolution/6-layers-ddd/internal/usecase"
	"github.com/golang-school/evolution/6-layers-ddd/pkg/httpserver"
	"github.com/golang-school/evolution/6-layers-ddd/pkg/logger"
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
	// Postgres
	pgPool, err := postgres.New(ctx, c.Postgres)
	if err != nil {
		return fmt.Errorf("postgres.New: %w", err)
	}

	// Redis
	redisClient, err := redis.New(c.Redis)
	if err != nil {
		return fmt.Errorf("redis.New: %w", err)
	}

	// Kafka producer
	kafkaProducer := kafka_produce.NewProducer(c.KafkaProducer)

	// Usecase (Service)
	profileUsecase := usecase.NewProfile(pgPool, kafkaProducer, redisClient)

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
	pgPool.Close()

	return nil
}
