package server

import (
	"github.com/golang-school/evolution/3-service-hell/internal/profile_service"
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
