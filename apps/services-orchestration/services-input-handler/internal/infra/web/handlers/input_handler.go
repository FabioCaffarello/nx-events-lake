package handlers

import (
	"apps/services-orchestration/services-input-handler/internal/entity"
	"apps/services-orchestration/services-input-handler/internal/usecase"
	"encoding/json"
	inputDTO "libs/golang/services/dtos/services-input-handler/input"
	"libs/golang/shared/go-events/events"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WebInputHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	InputRepository   entity.InputInterface
	InputCreatedEvent events.EventInterface
}

func NewWebInputHandler(
	EventDispatcher events.EventDispatcherInterface,
	InputRepository entity.InputInterface,
	InputCreatedEvent events.EventInterface,
) *WebInputHandler {
	return &WebInputHandler{
		EventDispatcher:   EventDispatcher,
		InputRepository:   InputRepository,
		InputCreatedEvent: InputCreatedEvent,
	}
}

func (h *WebInputHandler) CreateInput(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")
	contextEnv := chi.URLParam(r, "context")

	var dto inputDTO.InputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createInput := usecase.NewCreateInputUseCase(
		h.InputRepository,
		h.InputCreatedEvent,
		h.EventDispatcher,
	)

	output, err := createInput.Execute(dto, service, source, contextEnv)
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

func (h *WebInputHandler) ListAllByServiceAndSource(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")
	listInputs := usecase.NewListAllByServiceAndSourceUseCase(
		h.InputRepository,
	)
	output, err := listInputs.Execute(service, source)
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

func (h *WebInputHandler) ListAllByServiceAndSourceAndStatus(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")
	statusStr := chi.URLParam(r, "status")
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		http.Error(w, "Invalid status parameter", http.StatusBadRequest)
		return
	}

	listInputs := usecase.NewListAllByServiceAndSourceAndStatusUseCase(
		h.InputRepository,
	)
	output, err := listInputs.Execute(service, source, status)
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

func (h *WebInputHandler) ListAllByService(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	listInputs := usecase.NewListAllByServiceUseCase(
		h.InputRepository,
	)
	output, err := listInputs.Execute(service)
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

func (h *WebInputHandler) ListOneByIdAndService(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	id := chi.URLParam(r, "id")
	listInputs := usecase.NewListOneByIdAndServiceUseCase(
		h.InputRepository,
	)
	output, err := listInputs.Execute(service, id)
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
