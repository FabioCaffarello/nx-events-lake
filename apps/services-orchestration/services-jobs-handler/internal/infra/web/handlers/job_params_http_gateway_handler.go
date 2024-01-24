package handlers

import (
	"apps/services-orchestration/services-jobs-handler/internal/entity"
	"apps/services-orchestration/services-jobs-handler/internal/usecase"
	"encoding/json"
	inputDTO "libs/golang/services/dtos/services-jobs-handler/input"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebJobParamsHttpGatewayHandler struct {
	JobParamsHttpGatewayRepository        entity.HttpGatewayParamsInterface
	JobParamsHttpGatewayVersionRepository entity.HttpGatewayParamsVersionInterface
}

func NewWebJobParamsHttpGatewayHandler(
	JobParamsHttpGatewayRepository entity.HttpGatewayParamsInterface,
	JobParamsHttpGatewayVersionRepository entity.HttpGatewayParamsVersionInterface,
) *WebJobParamsHttpGatewayHandler {
	return &WebJobParamsHttpGatewayHandler{
		JobParamsHttpGatewayRepository:        JobParamsHttpGatewayRepository,
		JobParamsHttpGatewayVersionRepository: JobParamsHttpGatewayVersionRepository,
	}
}

func (h *WebJobParamsHttpGatewayHandler) CreateJobParamsHttpGateway(w http.ResponseWriter, r *http.Request) {
	var dto inputDTO.HttpGatewayParamsDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	checkJobParamsHttpGatewayExists := usecase.NewListOneJobParamsHttpGatewayByServiceAndSourceAndContextUseCase(
		h.JobParamsHttpGatewayRepository,
	)

	_, err = checkJobParamsHttpGatewayExists.Execute(dto.Service, dto.Source, dto.Context)
	if err == nil {
		http.Error(w, "JobParamsHttpGateway already exists", http.StatusInternalServerError)
		return
	}

	createJobParamsHttpGateway := usecase.NewCreateJobParamsHttpGatewayUseCase(
		h.JobParamsHttpGatewayRepository,
	)

	output, err := createJobParamsHttpGateway.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createJobParamsHttpGatewayVersion := usecase.NewCreateJobParamsHttpGatewayVersionUseCase(
		h.JobParamsHttpGatewayVersionRepository,
	)

	_, err = createJobParamsHttpGatewayVersion.Execute(output)
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

func (h *WebJobParamsHttpGatewayHandler) UpdateJobParamsHttpGateway(w http.ResponseWriter, r *http.Request) {
	var dto inputDTO.HttpGatewayParamsDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	checkJobParamsHttpGatewayExists := usecase.NewListOneJobParamsHttpGatewayByServiceAndSourceAndContextUseCase(
		h.JobParamsHttpGatewayRepository,
	)

	_, err = checkJobParamsHttpGatewayExists.Execute(dto.Service, dto.Source, dto.Context)
	if err != nil {
		http.Error(w, "JobParamsHttpGateway not exists", http.StatusInternalServerError)
		return
	}

	updateJobParamsHttpGateway := usecase.NewUpdateJobParamsHttpGatewayUseCase(
		h.JobParamsHttpGatewayRepository,
	)

	output, err := updateJobParamsHttpGateway.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updateJobParamsHttpGatewayVersion := usecase.NewUpdateJobParamsHttpGatewayVersionUseCase(
		h.JobParamsHttpGatewayVersionRepository,
	)

	_, err = updateJobParamsHttpGatewayVersion.Execute(output)
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

func (h *WebJobParamsHttpGatewayHandler) ListOneJobParamsHttpGatewayByServiceAndSourceAndContext(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")
	context := chi.URLParam(r, "context")

	log.Println("service", service)
	log.Println("source", source)
	log.Println("context", context)

	checkJobParamsHttpGatewayExists := usecase.NewListOneJobParamsHttpGatewayByServiceAndSourceAndContextUseCase(
		h.JobParamsHttpGatewayRepository,
	)

	output, err := checkJobParamsHttpGatewayExists.Execute(service, source, context)
	log.Println("output", output)
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
