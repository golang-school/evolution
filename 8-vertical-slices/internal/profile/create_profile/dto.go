package create_profile

import (
	"github.com/google/uuid"
)

type Output struct {
	ID uuid.UUID `json:"id"`
}

type Input struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}
