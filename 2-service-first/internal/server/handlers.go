package server

import (
	"github.com/golang-school/evolution/2-service-first/internal/profile_service"
)

// Обработчики HTTP запросов
type Handlers struct {
	profileService *profile_service.Profile
}

func NewHandlers(profileService *profile_service.Profile) *Handlers {
	return &Handlers{
		profileService: profileService,
	}
}
