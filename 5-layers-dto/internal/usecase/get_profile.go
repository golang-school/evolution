package usecase

import (
	"context"
	"fmt"

	"github.com/golang-school/evolution/5-layers-dto/internal/dto"

	"github.com/golang-school/evolution/5-layers-dto/internal/domain"

	"github.com/google/uuid"
)

func (p *Profile) GetProfile(ctx context.Context, input dto.GetProfileInput) (dto.GetProfileOutput, error) {
	// Валидируем ID
	profileID, err := uuid.Parse(input.ID)
	if err != nil {
		return dto.GetProfileOutput{}, domain.ErrUUIDInvalid
	}

	// Достаём профиль из БД
	profile, err := p.postgres.GetProfile(ctx, profileID)
	if err != nil {
		return dto.GetProfileOutput{}, fmt.Errorf("get profile from postgres: %w", err)
	}

	return dto.GetProfileOutput{
		Profile: profile,
	}, nil
}
