package http

import (
	"encoding/json"
	"net/http"

	"github.com/golang-school/evolution/7-layers-cqrs/internal/dto"

	"github.com/golang-school/evolution/7-layers-cqrs/pkg/render"
)

func (h *Handlers) CreateProfile(w http.ResponseWriter, r *http.Request) {
	input := dto.CreateProfileInput{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	output, err := h.profileService.CreateProfile(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	render.JSON(w, output, http.StatusOK)
}
