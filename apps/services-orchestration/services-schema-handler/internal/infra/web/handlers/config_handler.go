package handlers

import (
	"apps/services-orchestration/services-schema-handler/internal/entity"
	"apps/services-orchestration/services-schema-handler/internal/usecase"
	"encoding/json"
	inputDTO "libs/golang/services/dtos/services-schema-handler/input"
	"libs/golang/shared/go-events/events"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebSchemaHandler struct {
	EventDispatcher         events.EventDispatcherInterface
	SchemaRepository        entity.SchemaInterface
	SchemaCreatedEvent      events.EventInterface
	SchemaVersionRepository entity.SchemaVersionInterface
}

func NewWebSchemaHandler(
	EventDispatcher events.EventDispatcherInterface,
	SchemaRepository entity.SchemaInterface,
	SchemaCreatedEvent events.EventInterface,
	SchemaVersionRepository entity.SchemaVersionInterface,
) *WebSchemaHandler {
	return &WebSchemaHandler{
		EventDispatcher:         EventDispatcher,
		SchemaRepository:        SchemaRepository,
		SchemaCreatedEvent:      SchemaCreatedEvent,
		SchemaVersionRepository: SchemaVersionRepository,
	}
}

func (h *WebSchemaHandler) CreateSchema(w http.ResponseWriter, r *http.Request) {
	var dto inputDTO.SchemaDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	checkSchemaExists := usecase.NewListOneSchemaByServiceSourceAndContextAndSchemaTypeUseCase(
		h.SchemaRepository,
	)

	_, err = checkSchemaExists.Execute(dto.Service, dto.Source, dto.Context, dto.SchemaType)
	if err == nil {
		http.Error(w, "Schema already exists", http.StatusBadRequest)
		return
	}

	createSchema := usecase.NewCreateSchemaUseCase(
		h.SchemaRepository,
		h.SchemaCreatedEvent,
		h.EventDispatcher,
	)

	output, err := createSchema.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createSchemaVersion := usecase.NewCreateSchemaVersionUseCase(
		h.SchemaVersionRepository,
	)

	_, err = createSchemaVersion.Execute(output)
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

func (h *WebSchemaHandler) UpdateSchema(w http.ResponseWriter, r *http.Request) {
	var dto inputDTO.SchemaDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	checkSchemaExists := usecase.NewListOneSchemaByServiceSourceAndContextAndSchemaTypeUseCase(
		h.SchemaRepository,
	)

	_, err = checkSchemaExists.Execute(dto.Service, dto.Source, dto.Context, dto.SchemaType)
	if err != nil {
		http.Error(w, "Schema not found", http.StatusBadRequest)
		return
	}

	updateSchema := usecase.NewUpdateSchemaUseCase(
		h.SchemaRepository,
		h.SchemaCreatedEvent,
		h.EventDispatcher,
	)

	output, err := updateSchema.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updateSchemaVersion := usecase.NewUpdateSchemaVersionUseCase(
		h.SchemaVersionRepository,
	)

	_, err = updateSchemaVersion.Execute(output)
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

func (h *WebSchemaHandler) ListAllSchemas(w http.ResponseWriter, r *http.Request) {
	listAllSchemas := usecase.NewListAllSchemasUseCase(
		h.SchemaRepository,
	)

	output, err := listAllSchemas.Execute()
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

func (h *WebSchemaHandler) ListOneSchemaById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	listOneSchemaById := usecase.NewListOneSchemaByIdUseCase(
		h.SchemaRepository,
	)

	output, err := listOneSchemaById.Execute(id)
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

func (h *WebSchemaHandler) ListAllSchemasByService(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	listAllSchemasByService := usecase.NewListAllSchemasByServiceUseCase(
		h.SchemaRepository,
	)

	output, err := listAllSchemasByService.Execute(service)
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

func (h *WebSchemaHandler) ListAllSchemasByServiceAndContext(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	contextEnv := chi.URLParam(r, "context")

	listAllSchemasByServiceAndContext := usecase.NewListAllSchemasByServiceAndContextUseCase(
		h.SchemaRepository,
	)

	output, err := listAllSchemasByServiceAndContext.Execute(service, contextEnv)
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

func (h *WebSchemaHandler) ListOneSchemaByServiceSourceAndSchemaType(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")
	schemaType := chi.URLParam(r, "schemaType")
	listOneSchemaByServiceSourceAndSchemaType := usecase.NewListOneSchemaByServiceSourceAndSchemaTypeUseCase(
		h.SchemaRepository,
	)

	output, err := listOneSchemaByServiceSourceAndSchemaType.Execute(service, source, schemaType)
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

func (h *WebSchemaHandler) ListOneSchemaByServiceAndSourceAndContextAndSchemaType(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")
	context := chi.URLParam(r, "context")
	schemaType := chi.URLParam(r, "schemaType")

	listOneSchemaByServiceAndSourceAndContext := usecase.NewListOneSchemaByServiceSourceAndContextAndSchemaTypeUseCase(
		h.SchemaRepository,
	)

	output, err := listOneSchemaByServiceAndSourceAndContext.Execute(service, source, context, schemaType)
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

func (h *WebSchemaHandler) ListOneSchemaByIdAndVersionId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	versionId := chi.URLParam(r, "versionId")

	listOneSchemaByIdAndVersionId := usecase.NewListOneSchemaVersionByIdAndVersionIdUseCase(
		h.SchemaVersionRepository,
	)

	output, err := listOneSchemaByIdAndVersionId.Execute(id, versionId)
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

func (h *WebSchemaHandler) ListOneSchemaVersionById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	listOneSchemaById := usecase.NewListOneSchemaVersionByIdUseCase(
		h.SchemaVersionRepository,
	)

	output, err := listOneSchemaById.Execute(id)
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

func (h *WebSchemaHandler) ListAllSchemaVersion(w http.ResponseWriter, r *http.Request) {
	listAllSchemaVersion := usecase.NewListAllSchemasVersionUseCase(
		h.SchemaVersionRepository,
	)

	output, err := listAllSchemaVersion.Execute()
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
