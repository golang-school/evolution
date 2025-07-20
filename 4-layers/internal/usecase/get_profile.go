package usecase

import (
	"context"
	"fmt"

	"github.com/golang-school/evolution/4-layers/internal/domain"

	"github.com/google/uuid"
)

func (p *Profile) GetProfile(ctx context.Context, id string) (domain.Profile, error) {
	// Валидируем ID
	profileID, err := uuid.Parse(id)
	if err != nil {
		return domain.Profile{}, domain.ErrUUIDInvalid
	}

	// Достаём профиль из БД
	profile, err := p.postgres.GetProfile(ctx, profileID)
	if err != nil {
		return domain.Profile{}, fmt.Errorf("get profile from postgres: %w", err)
	}

	return profile, nil
}
