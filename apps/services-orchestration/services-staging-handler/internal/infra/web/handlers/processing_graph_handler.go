package handlers

import (
	"apps/services-orchestration/services-staging-handler/internal/entity"
	"apps/services-orchestration/services-staging-handler/internal/usecase"
	"encoding/json"
	inputDTO "libs/golang/services/dtos/services-staging-handler/input"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WebProcessingGraphHandler struct {
	ProcessingGraphRepository entity.ProcessingGraphInterface
}

func NewWebProcessingGraphHandler(
	ProcessingGraphRepository entity.ProcessingGraphInterface,
) *WebProcessingGraphHandler {
	return &WebProcessingGraphHandler{
		ProcessingGraphRepository: ProcessingGraphRepository,
	}
}

func (h *WebProcessingGraphHandler) CreateProcessingGraphHandler(w http.ResponseWriter, r *http.Request) {
	var inputDTO inputDTO.ProcessingGraphDTO
	err := json.NewDecoder(r.Body).Decode(&inputDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	useCase := usecase.NewCreateProcessingGraphUseCase(
		h.ProcessingGraphRepository,
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

func (h *WebProcessingGraphHandler) ListOneProcessingGraphBySourceAndStartProcessingIdUseCase(w http.ResponseWriter, r *http.Request) {
	source := chi.URLParam(r, "source")
	startProcessingId := chi.URLParam(r, "start_processing_id")

	useCase := usecase.NewListOneProcessingGraphBySourceAndStartProcessingIdUseCase(
		h.ProcessingGraphRepository,
	)

	output, err := useCase.Execute(source, startProcessingId)
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

func (h *WebProcessingGraphHandler) ListOneProcessingGraphByTaskSourceAndParentProcessingIdUseCase(w http.ResponseWriter, r *http.Request) {
	source := chi.URLParam(r, "source")
	parentProcessingId := chi.URLParam(r, "parent_processing_id")

	useCase := usecase.NewListOneProcessingGraphByTaskSourceAndParentProcessingIdUseCase(
		h.ProcessingGraphRepository,
	)

	output, err := useCase.Execute(source, parentProcessingId)
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

func (h *WebProcessingGraphHandler) CreateTaskToProcessingGraphHandler(w http.ResponseWriter, r *http.Request) {
	source := chi.URLParam(r, "source")
	startProcessingId := chi.URLParam(r, "start_processing_id")

	var inputDTO inputDTO.Task
	err := json.NewDecoder(r.Body).Decode(&inputDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	useCase := usecase.NewCreateTaskToProcessingGraphUseCase(
		h.ProcessingGraphRepository,
	)

	output, err := useCase.Execute(source, startProcessingId, inputDTO)
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

func (h *WebProcessingGraphHandler) UpdateTaskStatusProcessingGraphHandler(w http.ResponseWriter, r *http.Request) {
	source := chi.URLParam(r, "source")
	processingId := chi.URLParam(r, "processing_id")
	statusStr := chi.URLParam(r, "status")
    processingTimestamp := chi.URLParam(r, "processing_timestamp")
	statusCode, err := strconv.Atoi(statusStr)
	if err != nil {
		http.Error(w, "Invalid status parameter", http.StatusBadRequest)
		return
	}

	useCase := usecase.NewUpdateTaskProcessingStatusGraphUseCase(
		h.ProcessingGraphRepository,
	)

	output, err := useCase.Execute(source, processingId, statusCode, processingTimestamp)
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

func (h *WebProcessingGraphHandler) UpdateTaskOutputProcessingGraphHandler(w http.ResponseWriter, r *http.Request) {
    source := chi.URLParam(r, "source")
    processingId := chi.URLParam(r, "processing_id")
    outputId := chi.URLParam(r, "output_id")

    useCase := usecase.NewUpdateProcessingGraphTaskOutputUseCase(
        h.ProcessingGraphRepository,
    )

    output, err := useCase.Execute(source, processingId, outputId)
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
