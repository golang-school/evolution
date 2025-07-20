package create_profile

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

type Config struct {
	Addr  []string `envconfig:"KAFKA_CONSUMER_ADDR"     required:"true"`
	Topic string   `default:"my-topic"  envconfig:"KAFKA_CONSUMER_TOPIC"`
	Group string   `default:"my-group"  envconfig:"KAFKA_CONSUMER_GROUP"`
}

type Consumer struct {
	config Config
	reader *kafka.Reader
	stop   context.CancelFunc
	done   chan struct{}
}

func NewConsumer(cfg Config) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:          cfg.Addr,
		Topic:            cfg.Topic,
		GroupID:          cfg.Group,
		ReadBatchTimeout: time.Second,
		CommitInterval:   time.Second,
	})

	ctx, stop := context.WithCancel(context.Background())

	c := &Consumer{
		config: cfg,
		reader: r,
		stop:   stop,
		done:   make(chan struct{}),
	}

	go c.run(ctx)

	return c
}

func (c *Consumer) run(ctx context.Context) {
	log.Info().Msg("kafka consumer: started")

	for {
		m, err := c.reader.FetchMessage(ctx)
		if err != nil {
			log.Error().Err(err).Msg("kafka consumer: FetchMessage")

			if errors.Is(err, io.EOF) || errors.Is(err, context.Canceled) {
				break
			}
		}

		_, err = usecase.CreateProfile(ctx, Input{})
		if err != nil {
			log.Error().Err(err).Msg("kafka consumer: create profile usecase")
		}

		if err = c.reader.CommitMessages(ctx, m); err != nil {
			log.Error().Err(err).Msg("kafka consumer: CommitMessages")
		}
	}

	close(c.done)
}

func (c *Consumer) Close() {
	log.Info().Msg("kafka consumer: closing")

	c.stop()

	if err := c.reader.Close(); err != nil {
		log.Error().Err(err).Msg("kafka consumer: reader.Close")
	}

	<-c.done

	log.Info().Msg("kafka consumer: closed")
}
