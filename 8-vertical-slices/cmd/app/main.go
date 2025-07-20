package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-school/evolution/8-vertical-slices/internal/profile/get_profile"

	"github.com/golang-school/evolution/8-vertical-slices/internal/profile/create_profile"

	"github.com/golang-school/evolution/8-vertical-slices/config"
	"github.com/golang-school/evolution/8-vertical-slices/internal/adapter/kafka_produce"
	"github.com/golang-school/evolution/8-vertical-slices/internal/adapter/postgres"
	"github.com/golang-school/evolution/8-vertical-slices/internal/adapter/redis"
	"github.com/golang-school/evolution/8-vertical-slices/internal/controller/http"
	"github.com/golang-school/evolution/8-vertical-slices/pkg/httpserver"
	"github.com/golang-school/evolution/8-vertical-slices/pkg/logger"
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

	// Create profile usecase
	create_profile.New(pgPool, kafkaProducer, redisClient)
	createProfileConsumer := create_profile.NewConsumer(c.KafkaConsumer)

	// Get profile usecase
	get_profile.New(pgPool)

	// HTTP сервер
	router := http.Router()
	httpServer := httpserver.New(router, c.HTTP)

	// Приложение запущено и готово к работе

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	<-sig // ждём здесь сигнала (Ctrl+C или SIGTERM)

	// Закрываем ресурсы
	createProfileConsumer.Close()
	httpServer.Close()
	redisClient.Close()
	kafkaProducer.Close()
	pgPool.Close()

	return nil
}
