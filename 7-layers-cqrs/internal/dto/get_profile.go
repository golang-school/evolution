package dto

import (
	"github.com/golang-school/evolution/7-layers-cqrs/internal/domain"
)

type GetProfileOutput struct {
	domain.Profile
}

type GetProfileInput struct {
	ID string
}
