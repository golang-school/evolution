package get_profile

import (
	"github.com/golang-school/evolution/8-vertical-slices/internal/domain"
)

type Output struct {
	domain.Profile
}

type Input struct {
	ID string
}
