package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-school/evolution/1-handler-first/pkg/render"

	"github.com/golang-school/evolution/1-handler-first/internal/kafka_produce"
	"github.com/golang-school/evolution/1-handler-first/internal/model"
	"github.com/google/uuid"
)

func (h *Handlers) CreateProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

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

	// Проверяем в Redis ключу идемпотентности
	if h.redis.IsExists(ctx, input.Email) {
		http.Error(w, ErrAlreadyExists.Error(), http.StatusBadRequest)

		return
	}

	// Создаём профиль
	profile := model.Profile{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		Name:      input.Name,
		Age:       input.Age,
		Email:     input.Email,
	}

	// Валидируем
	err = profile.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	// Сохраняем в БД
	err = h.postgres.CreateProfile(ctx, profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	// Отправляем в Kafka событие создания профиля
	err = h.kafka.Produce(ctx, kafka_produce.Message{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	render.JSON(w, profile.ID, http.StatusOK)
}
