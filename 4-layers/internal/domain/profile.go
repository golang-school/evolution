package domain

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"  validate:"required,min=3,max=64"`
	Age       int       `json:"age"   validate:"required,min=18,max=120"`
	Email     string    `json:"email" validate:"email"`
}

func (p Profile) Validate() error {
	// валидация
	return nil
}
