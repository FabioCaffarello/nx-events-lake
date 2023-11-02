package handlers

import (
	"apps/services-orchestration/services-input-handler/internal/entity"
	"apps/services-orchestration/services-input-handler/internal/usecase"
	"encoding/json"
	sharedDTO "libs/golang/services/dtos/services-input-handler/shared"
	"libs/golang/shared/go-events/events"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebInputStatusHandler struct {
	EventDispatcher         events.EventDispatcherInterface
	InputRepository         entity.InputInterface
	InputStatusUpdatedEvent events.EventInterface
}

func NewWebInputStatusHandler(
	EventDispatcher events.EventDispatcherInterface,
	InputRepository entity.InputInterface,
	InputStatusUpdatedEvent events.EventInterface,
) *WebInputStatusHandler {
	return &WebInputStatusHandler{
		EventDispatcher:         EventDispatcher,
		InputRepository:         InputRepository,
		InputStatusUpdatedEvent: InputStatusUpdatedEvent,
	}
}

func (h *WebInputStatusHandler) UpdateInputStatus(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")
	contextEnv := chi.URLParam(r, "context")
	id := chi.URLParam(r, "id")

	var status sharedDTO.Status
	err := json.NewDecoder(r.Body).Decode(&status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateInputStatus := usecase.NewUpdateStatusInputUseCase(
		h.InputRepository,
		h.InputStatusUpdatedEvent,
		h.EventDispatcher,
	)

	output, err := updateInputStatus.Execute(service, source, contextEnv, id, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
