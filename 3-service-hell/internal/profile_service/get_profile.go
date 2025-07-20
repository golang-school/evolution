package profile_service

import (
	"context"
	"fmt"

	"github.com/golang-school/evolution/3-service-hell/internal/model"
	"github.com/google/uuid"
)

func (p *Profile) GetProfile(ctx context.Context, id string) (model.Profile, error) {
	// Валидируем ID
	profileID, err := uuid.Parse(id)
	if err != nil {
		return model.Profile{}, ErrUUIDInvalid
	}

	// Достаём профиль из БД
	profile, err := p.postgres.GetProfile(ctx, profileID)
	if err != nil {
		return model.Profile{}, fmt.Errorf("get profile from postgres: %w", err)
	}

	return profile, nil
}
