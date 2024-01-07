package usecase

import (
	"apps/services-orchestration/services-config-handler/internal/entity"
	"fmt"
	inputDTO "libs/golang/services/dtos/services-config-handler/input"
	outputDTO "libs/golang/services/dtos/services-config-handler/output"
	"libs/golang/shared/go-events/events"
)

type UpdateConfigUseCase struct {
	ConfigRepository entity.ConfigInterface
	ConfigUpdated    events.EventInterface
	EventDispatcher  events.EventDispatcherInterface
}

func NewUpdateConfigUseCase(
	repository entity.ConfigInterface,
	ConfigUpdated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *UpdateConfigUseCase {
	return &UpdateConfigUseCase{
		ConfigRepository: repository,
		ConfigUpdated:    ConfigUpdated,
		EventDispatcher:  EventDispatcher,
	}
}

func (ccu *UpdateConfigUseCase) Execute(config inputDTO.ConfigDTO) (outputDTO.ConfigDTO, error) {
	// FIXME:

	configEntity, err := entity.NewConfig(
		config.Name,
		config.Active,
		config.Frequency,
		config.Service,
		config.Source,
		config.Context,
        config.OutputMethod,
		ConvertDependsOnDTOToEntity(config.DependsOn),
		config.JobParameters,
		config.ServiceParameters,
	)
	if err != nil {
		return outputDTO.ConfigDTO{}, err
	}

	err = ccu.ConfigRepository.Save(configEntity)
	if err != nil {
		return outputDTO.ConfigDTO{}, err
	}

	usecaseDeps := ConvertEntityToUseCaseDependencies(configEntity.DependsOn)

	dto := outputDTO.ConfigDTO{
		ID:                string(configEntity.ID),
		Name:              configEntity.Name,
		Active:            configEntity.Active,
		Frequency:         configEntity.Frequency,
		Service:           configEntity.Service,
		Source:            configEntity.Source,
		Context:           configEntity.Context,
        OutputMethod:      configEntity.OutputMethod,
		DependsOn:         usecaseDeps,
		ConfigID:          configEntity.ConfigID,
		ServiceParameters: configEntity.ServiceParameters,
		JobParameters:     configEntity.JobParameters,
		CreatedAt:         configEntity.CreatedAt,
		UpdatedAt:         configEntity.UpdatedAt,
	}

	ccu.ConfigUpdated.SetPayload(dto)
	ccu.EventDispatcher.Dispatch(ccu.ConfigUpdated, "services", fmt.Sprintf("config.%s.%s.%s", dto.Context, dto.Service, dto.Source))

	return dto, nil
}
