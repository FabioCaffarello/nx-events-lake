package handlers

import (
	"apps/services-orchestration/services-staging-handler/internal/entity"
	"apps/services-orchestration/services-staging-handler/internal/usecase"
	"encoding/json"
	inputDTO "libs/golang/services/dtos/services-staging-handler/input"
	sharedDTO "libs/golang/services/dtos/services-staging-handler/shared"
	"libs/golang/shared/go-events/events"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebProcessingJobDependenciesHandler struct {
	EventDispatcher                       events.EventDispatcherInterface
	ProcessingJobDependenciesRepository   entity.ProcessingJobDependenciesInterface
	ProcessingJobDependenciesCreatedEvent events.EventInterface
}

func NewWebProcessingJobDependenciesHandler(
    EventDispatcher events.EventDispatcherInterface,
	ProcessingJobDependenciesRepository entity.ProcessingJobDependenciesInterface,
    ProcessingJobDependenciesCreatedEvent events.EventInterface,
) *WebProcessingJobDependenciesHandler {
	return &WebProcessingJobDependenciesHandler{
        EventDispatcher:                       EventDispatcher,
		ProcessingJobDependenciesRepository: ProcessingJobDependenciesRepository,
        ProcessingJobDependenciesCreatedEvent: ProcessingJobDependenciesCreatedEvent,
	}
}

func (h *WebProcessingJobDependenciesHandler) CreateProcessingJobDependenciesHandler(w http.ResponseWriter, r *http.Request) {
	var inputDTO inputDTO.ProcessingJobDependenciesDTO
	err := json.NewDecoder(r.Body).Decode(&inputDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	useCase := usecase.NewCreateProcessingJobDependenciesUseCase(
        h.ProcessingJobDependenciesRepository,
        h.ProcessingJobDependenciesCreatedEvent,
        h.EventDispatcher,
    )

	output, err := useCase.Execute(inputDTO)
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

func (h *WebProcessingJobDependenciesHandler) ListOneProcessingJobDependenciesByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	useCase := usecase.NewListOneProcessingJobDependenciesByIdUseCase(
		h.ProcessingJobDependenciesRepository,
	)

	output, err := useCase.Execute(id)
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

func (h *WebProcessingJobDependenciesHandler) RemoveProcessingJobDependenciesHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	useCase := usecase.NewRemoveProcessingJobDependenciesUseCase(
		h.ProcessingJobDependenciesRepository,
	)

	err := useCase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebProcessingJobDependenciesHandler) UpdateProcessingJobDependenciesHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var jobDepsDTO sharedDTO.ProcessingJobDependencies
	err := json.NewDecoder(r.Body).Decode(&jobDepsDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("UpdateProcessingJobDependenciesHandler: id=%s, jobDepsDTO=%v", id, jobDepsDTO)

	useCase := usecase.NewUpdateProcessingJobDependenciesUseCase(
		h.ProcessingJobDependenciesRepository,
	)

	err = useCase.Execute(jobDepsDTO, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updatedDoc, err := h.ProcessingJobDependenciesRepository.FindOneById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(updatedDoc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
