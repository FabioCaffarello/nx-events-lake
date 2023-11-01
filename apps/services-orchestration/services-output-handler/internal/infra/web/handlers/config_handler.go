package handlers

import (
	"apps/services-orchestration/services-output-handler/internal/entity"
	"apps/services-orchestration/services-output-handler/internal/usecase"
	"encoding/json"
	inputDTO "libs/golang/services/dtos/services-output-handler/input"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebServiceOutputHandler struct {
	ServiceOutputRepository entity.ServiceOutputInterface
}

func NewWebServiceOutputHandler(
	ServiceOutputRepository entity.ServiceOutputInterface,
) *WebServiceOutputHandler {
	return &WebServiceOutputHandler{
		ServiceOutputRepository: ServiceOutputRepository,
	}
}

func (h *WebServiceOutputHandler) CreateServiceOutput(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	var dto inputDTO.ServiceOutputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createServiceOutput := usecase.NewCreateServiceOutputUseCase(
		h.ServiceOutputRepository,
	)

	output, err := createServiceOutput.Execute(dto, service)
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

func (h *WebServiceOutputHandler) ListAllServiceOutputsByServiceAndId(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	id := chi.URLParam(r, "id")

	listOneServiceOutputByServiceAndId := usecase.NewListOneServiceOutputByServiceAndIdUseCase(
		h.ServiceOutputRepository,
	)

	output, err := listOneServiceOutputByServiceAndId.Execute(service, id)
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

func (h *WebServiceOutputHandler) ListAllServiceOutputsByService(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")

	listAllServiceOutputByService := usecase.NewListAllServiceOutputByServiceUseCase(
		h.ServiceOutputRepository,
	)

	output, err := listAllServiceOutputByService.Execute(service)
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

func (h *WebServiceOutputHandler) ListAllServiceOutputsByServiceAndSource(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")

	listAllServiceOutputByServiceAndSource := usecase.NewListAllServiceOutputByServiceAndSourceUseCase(
		h.ServiceOutputRepository,
	)

	output, err := listAllServiceOutputByServiceAndSource.Execute(service, source)
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
