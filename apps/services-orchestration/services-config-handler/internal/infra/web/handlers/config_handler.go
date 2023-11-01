package handlers

import (
	"apps/services-orchestration/services-config-handler/internal/entity"
	"apps/services-orchestration/services-config-handler/internal/usecase"
	"encoding/json"
	inputDTO "libs/golang/services/dtos/services-config-handler/input"
	"libs/golang/shared/go-events/events"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebConfigHandler struct {
	EventDispatcher    events.EventDispatcherInterface
	ConfigRepository   entity.ConfigInterface
	ConfigCreatedEvent events.EventInterface
	// ConfigUpdatedEvent      events.EventInterface
	ConfigVersionRepository entity.ConfigVersionInterface
}

func NewWebConfigHandler(
	EventDispatcher events.EventDispatcherInterface,
	ConfigRepository entity.ConfigInterface,
	ConfigCreatedEvent events.EventInterface,
	// ConfigUpdatedEvent events.EventInterface,
	ConfigVersionRepository entity.ConfigVersionInterface,
) *WebConfigHandler {
	return &WebConfigHandler{
		EventDispatcher:    EventDispatcher,
		ConfigRepository:   ConfigRepository,
		ConfigCreatedEvent: ConfigCreatedEvent,
		// ConfigUpdatedEvent:      ConfigUpdatedEvent,
		ConfigVersionRepository: ConfigVersionRepository,
	}
}

func (h *WebConfigHandler) CreateConfig(w http.ResponseWriter, r *http.Request) {
	var dto inputDTO.ConfigDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	checkConfigExists := usecase.NewListOneConfigByServiceAndSourceAndContextUseCase(
		h.ConfigRepository,
	)

	_, err = checkConfigExists.Execute(dto.Service, dto.Source, dto.Context)
	if err == nil {
		http.Error(w, "Config already exists", http.StatusBadRequest)
		return
	}

	createConfig := usecase.NewCreateConfigUseCase(
		h.ConfigRepository,
		h.ConfigCreatedEvent,
		h.EventDispatcher,
	)

	output, err := createConfig.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createConfigVersion := usecase.NewCreateConfigVersionUseCase(
		h.ConfigVersionRepository,
	)

	_, err = createConfigVersion.Execute(output)
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

func (h *WebConfigHandler) UpdateConfig(w http.ResponseWriter, r *http.Request) {
	var dto inputDTO.ConfigDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	checkConfigExists := usecase.NewListOneConfigByServiceAndSourceAndContextUseCase(
		h.ConfigRepository,
	)

	_, err = checkConfigExists.Execute(dto.Service, dto.Source, dto.Context)
	if err != nil {
		http.Error(w, "Config not found", http.StatusBadRequest)
		return
	}

	updateConfig := usecase.NewUpdateConfigUseCase(
		h.ConfigRepository,
		h.ConfigCreatedEvent,
		h.EventDispatcher,
	)

	output, err := updateConfig.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updateConfigVersion := usecase.NewUpdateConfigVersionUseCase(
		h.ConfigVersionRepository,
	)

	_, err = updateConfigVersion.Execute(output)
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

func (h *WebConfigHandler) ListAllConfigs(w http.ResponseWriter, r *http.Request) {
	listAllConfigs := usecase.NewListAllConfigsUseCase(
		h.ConfigRepository,
	)

	output, err := listAllConfigs.Execute()
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

func (h *WebConfigHandler) ListOneConfigById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	listOneConfigById := usecase.NewListOneConfigByIdUseCase(
		h.ConfigRepository,
	)

	output, err := listOneConfigById.Execute(id)
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

func (h *WebConfigHandler) ListAllConfigsByService(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	listAllConfigsByService := usecase.NewListAllConfigsByServiceUseCase(
		h.ConfigRepository,
	)

	output, err := listAllConfigsByService.Execute(service)
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

func (h *WebConfigHandler) ListAllConfigsByServiceAndContext(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	contextEnv := chi.URLParam(r, "context")

	listAllConfigsByServiceAndContext := usecase.NewListAllConfigsByServiceAndContextUseCase(
		h.ConfigRepository,
	)

	output, err := listAllConfigsByServiceAndContext.Execute(service, contextEnv)
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

func (h *WebConfigHandler) ListAllConfigsByDependentJob(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")

	listAllConfigsByDependentJob := usecase.NewListAllConfigsByDependentJobUseCase(
		h.ConfigRepository,
	)

	output, err := listAllConfigsByDependentJob.Execute(service, source)
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

func (h *WebConfigHandler) ListOneConfigByServiceAndSourceAndContext(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")
	context := chi.URLParam(r, "context")

	listOneConfigByServiceAndSourceAndContext := usecase.NewListOneConfigByServiceAndSourceAndContextUseCase(
		h.ConfigRepository,
	)

	output, err := listOneConfigByServiceAndSourceAndContext.Execute(service, source, context)
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

func (h *WebConfigHandler) ListOneConfigByIdAndVersionId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	versionId := chi.URLParam(r, "versionId")

	listOneConfigByIdAndVersionId := usecase.NewListOneConfigVersionByIdAndVersionIdUseCase(
		h.ConfigVersionRepository,
	)

	output, err := listOneConfigByIdAndVersionId.Execute(id, versionId)
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

func (h *WebConfigHandler) ListOneConfigVersionById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	listOneConfigById := usecase.NewListOneConfigVersionByIdUseCase(
		h.ConfigVersionRepository,
	)

	output, err := listOneConfigById.Execute(id)
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

func (h *WebConfigHandler) ListAllConfigVersion(w http.ResponseWriter, r *http.Request) {
	listAllConfigVersion := usecase.NewListAllConfigsVersionUseCase(
		h.ConfigVersionRepository,
	)

	output, err := listAllConfigVersion.Execute()
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
