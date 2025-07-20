package dto

import (
	"github.com/golang-school/evolution/5-layers-dto/internal/domain"
)

type GetProfileOutput struct {
	domain.Profile
}

type GetProfileInput struct {
	ID string
}
