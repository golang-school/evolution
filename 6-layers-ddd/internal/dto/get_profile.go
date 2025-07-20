package dto

import (
	"github.com/golang-school/evolution/6-layers-ddd/internal/domain"
)

type GetProfileOutput struct {
	domain.Profile
}

type GetProfileInput struct {
	ID string
}
