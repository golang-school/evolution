package otel

import (
	"context"
)

type Config struct {
	AppName    string  `envconfig:"APP_NAME"`
	AppVersion string  `envconfig:"APP_VERSION"`
	Endpoint   string  `envconfig:"OTEL_ENDPOINT"`
	Namespace  string  `envconfig:"OTEL_NAMESPACE"`
	InstanceID string  `envconfig:"OTEL_INSTANCE_ID"`
	Ratio      float64 `default:"1.0"  envconfig:"OTEL_RATIO"`
}

func Init(ctx context.Context, c Config) error {
	// Делаем настройки

	return nil
}

func Close() {
	// Shutdown
}
