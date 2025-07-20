package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-school/evolution/2-service-first/config"
	"github.com/golang-school/evolution/2-service-first/internal/kafka_consume"
	"github.com/golang-school/evolution/2-service-first/internal/kafka_produce"
	"github.com/golang-school/evolution/2-service-first/internal/postgres"
	"github.com/golang-school/evolution/2-service-first/internal/profile_service"
	"github.com/golang-school/evolution/2-service-first/internal/redis"
	"github.com/golang-school/evolution/2-service-first/internal/server"
	"github.com/golang-school/evolution/2-service-first/pkg/httpserver"
	"github.com/golang-school/evolution/2-service-first/pkg/logger"
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

	// Service
	profileService := profile_service.NewProfile(pgPool, kafkaProducer, redisClient)

	// Kafka consumer
	kafkaConsumer := kafka_consume.New(c.KafkaConsumer, profileService)

	// HTTP сервер
	router := server.Router(profileService)
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
