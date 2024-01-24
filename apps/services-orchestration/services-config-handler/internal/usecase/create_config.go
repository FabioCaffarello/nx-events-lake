package usecase

import (
	"apps/services-orchestration/services-config-handler/internal/entity"
	"fmt"
	inputDTO "libs/golang/services/dtos/services-config-handler/input"
	outputDTO "libs/golang/services/dtos/services-config-handler/output"
	"libs/golang/shared/go-events/events"
)

type CreateConfigUseCase struct {
	ConfigRepository entity.ConfigInterface
	ConfigCreated    events.EventInterface
	EventDispatcher  events.EventDispatcherInterface
}

func NewCreateConfigUseCase(
	repository entity.ConfigInterface,
	ConfigCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *CreateConfigUseCase {
	return &CreateConfigUseCase{
		ConfigRepository: repository,
		ConfigCreated:    ConfigCreated,
		EventDispatcher:  EventDispatcher,
	}
}

func (ccu *CreateConfigUseCase) Execute(config inputDTO.ConfigDTO) (outputDTO.ConfigDTO, error) {
	configEntity, err := entity.NewConfig(
		config.Name,
		config.Active,
		config.Frequency,
		config.Service,
		config.Source,
		config.Context,
        config.InputMethod,
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
        InputMethod:       configEntity.InputMethod,
        OutputMethod:      configEntity.OutputMethod,
		DependsOn:         usecaseDeps,
		ConfigID:          configEntity.ConfigID,
		ServiceParameters: configEntity.ServiceParameters,
		JobParameters:     configEntity.JobParameters,
		CreatedAt:         configEntity.CreatedAt,
		UpdatedAt:         configEntity.UpdatedAt,
	}

	ccu.ConfigCreated.SetPayload(dto)
	ccu.EventDispatcher.Dispatch(ccu.ConfigCreated, "services", fmt.Sprintf("config.%s.%s.%s", dto.Context, dto.Service, dto.Source))

	return dto, nil
}
