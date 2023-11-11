package handlers

import (
	"encoding/json"
	"net/http"

	"apps/services-orchestration/services-file-catalog-handler/internal/entity"
	"apps/services-orchestration/services-file-catalog-handler/internal/usecase"
	inputDTO "libs/golang/services/dtos/services-file-catalog-handler/input"
	"github.com/go-chi/chi/v5"
)

type WebFileCatalogHandler struct {
	FileCatalogRepository entity.FileCatalogInterface
}

func NewWebFileCatalogHandler(
	repository entity.FileCatalogInterface,
) *WebFileCatalogHandler {
	return &WebFileCatalogHandler{
		FileCatalogRepository: repository,
	}
}

func (wsh *WebFileCatalogHandler) CreateFileCatalog(w http.ResponseWriter, r *http.Request) {
	var dto inputDTO.FileCatalogDTO
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createFileCatalog := usecase.NewCreateFileCatalogUseCase(
		wsh.FileCatalogRepository,
	)

	output, err := createFileCatalog.Execute(dto)
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

func (wsh *WebFileCatalogHandler) ListAllFileCatalogs(w http.ResponseWriter, r *http.Request) {
     listAllFileCatalogs := usecase.NewListAllFileCatalogsUseCase(
          wsh.FileCatalogRepository,
     )

     output, err := listAllFileCatalogs.Execute()
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

func (wsh *WebFileCatalogHandler) ListOneFileCatalogById(w http.ResponseWriter, r *http.Request) {
     id := chi.URLParam(r, "id")

     listOneFileCatalogById := usecase.NewListOneFileCatalogByIdUseCase(
          wsh.FileCatalogRepository,
     )

     output, err := listOneFileCatalogById.Execute(id)
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

func (wsh *WebFileCatalogHandler) ListAllFileCatalogsByService(w http.ResponseWriter, r *http.Request) {
     service := chi.URLParam(r, "service")

     listAllFileCatalogsByService := usecase.NewListAllFileCatalogsByServiceUseCase(
          wsh.FileCatalogRepository,
     )

     output, err := listAllFileCatalogsByService.Execute(service)
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

func (wsh *WebFileCatalogHandler) ListOneFileCatalogByServiceAndSource(w http.ResponseWriter, r *http.Request) {
     service := chi.URLParam(r, "service")
     source := chi.URLParam(r, "source")

     listOneFileCatalogByServiceAndSource := usecase.NewListOneFileCatalogByServiceAndSourceUseCase(
          wsh.FileCatalogRepository,
     )

     output, err := listOneFileCatalogByServiceAndSource.Execute(service, source)
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
