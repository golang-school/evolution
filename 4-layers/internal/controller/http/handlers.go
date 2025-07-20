package http

import (
	"github.com/golang-school/evolution/4-layers/internal/usecase"
)

// Обработчики HTTP запросов
type Handlers struct {
	profileService *usecase.Profile
}

func NewHandlers(profileService *usecase.Profile) *Handlers {
	return &Handlers{
		profileService: profileService,
	}
}
