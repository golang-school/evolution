package domain

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Name string

type Age int

type Email string

type Profile struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      Name      `json:"name"  validate:"required,min=3,max=64"`
	Age       Age       `json:"age"   validate:"required,min=18,max=120"`
	Email     Email     `json:"email" validate:"email"`
}

var validate = validator.New(validator.WithRequiredStructEnabled())

func NewProfile(name string, age int, email string) (Profile, error) {
	p := Profile{
		ID:    uuid.New(),
		Name:  Name(name),
		Age:   Age(age),
		Email: Email(email),
	}

	if err := p.Validate(); err != nil {
		return Profile{}, fmt.Errorf("p.Validate: %w", err)
	}

	return p, nil
}

func (p Profile) Validate() error {
	err := validate.Struct(p)
	if err != nil {
		return fmt.Errorf("validate.Struct: %w", err)
	}

	return nil
}
