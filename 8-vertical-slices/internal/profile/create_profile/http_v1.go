package create_profile

import (
	"encoding/json"
	"net/http"

	"github.com/golang-school/evolution/8-vertical-slices/pkg/render"
)

var usecase *Usecase

func HTTPv1(w http.ResponseWriter, r *http.Request) {
	input := Input{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	output, err := usecase.CreateProfile(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	render.JSON(w, output, http.StatusOK)
}
