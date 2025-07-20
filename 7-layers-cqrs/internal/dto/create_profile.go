package dto

import (
	"github.com/google/uuid"
)

type CreateProfileOutput struct {
	ID uuid.UUID `json:"id"`
}

type CreateProfileInput struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}
