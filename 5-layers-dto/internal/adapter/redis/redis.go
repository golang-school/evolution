package redis

import "context"

type Config struct {
	Addr     string `envconfig:"REDIS_ADDR"  required:"true"`
	Password string `envconfig:"REDIS_PASSWORD"`
	DB       int    `default:"0"  envconfig:"REDIS_DB"`
}

type Client struct{}

func New(c Config) (*Client, error) {
	// Делаем подключение к Redis
	return &Client{}, nil
}

func (c *Client) IsExists(ctx context.Context, idempotencyKey string) bool {
	// Проверяем, существует ли ключ в Redis
	return false
}

func (c *Client) Close() {
	// Shutdown
}
