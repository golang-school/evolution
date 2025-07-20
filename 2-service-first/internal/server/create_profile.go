package server

import (
	"encoding/json"
	"net/http"

	"github.com/golang-school/evolution/2-service-first/pkg/render"
)

func (h *Handlers) CreateProfile(w http.ResponseWriter, r *http.Request) {
	type Input struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	input := Input{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	output, err := h.profileService.CreateProfile(r.Context(), input.Name, input.Age, input.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	render.JSON(w, output, http.StatusOK)
}
