package usecase

import (
	"apps/services-orchestration/services-config-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-config-handler/output"
)

type UpdateConfigVersionUseCase struct {
	ConfigVersionRepository entity.ConfigVersionInterface
}

func NewUpdateConfigVersionUseCase(
	repository entity.ConfigVersionInterface,
) *UpdateConfigVersionUseCase {
	return &UpdateConfigVersionUseCase{
		ConfigVersionRepository: repository,
	}
}

func (uc *UpdateConfigVersionUseCase) Execute(config outputDTO.ConfigDTO) (outputDTO.ConfigDTO, error) {
	configEntity := ConvertConfigDTOToEntity(config)
	err := uc.ConfigVersionRepository.Update(configEntity)
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

	return dto, nil
}
