package transaction

import (
	"context"
)

func Wrap(ctx context.Context, fn func(context.Context) error) error {
	// Обобщённая функция для обёртки транзакций

	return nil
}
