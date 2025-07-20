package get_profile

import (
	"context"
	"fmt"

	"github.com/golang-school/evolution/8-vertical-slices/internal/adapter/postgres"
	"github.com/golang-school/evolution/8-vertical-slices/internal/domain"
	"github.com/google/uuid"
)

type Postgres interface {
	GetProfile(ctx context.Context, id uuid.UUID) (domain.Profile, error)
}

type Usecase struct {
	postgres Postgres
}

func New(postgres *postgres.Pool) *Usecase {
	uc := &Usecase{
		postgres: postgres,
	}

	usecase = uc // global for handlers

	return uc
}

func (u *Usecase) GetProfile(ctx context.Context, input Input) (Output, error) {
	// Валидируем ID
	profileID, err := uuid.Parse(input.ID)
	if err != nil {
		return Output{}, domain.ErrUUIDInvalid
	}

	// Достаём профиль из БД
	profile, err := u.postgres.GetProfile(ctx, profileID)
	if err != nil {
		return Output{}, fmt.Errorf("get profile from postgres: %w", err)
	}

	return Output{
		Profile: profile,
	}, nil
}
