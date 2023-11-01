package usecase

import (
	"apps/services-orchestration/services-config-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-config-handler/output"
)

type CreateConfigVersionUseCase struct {
	ConfigVersionRepository entity.ConfigVersionInterface
}

func NewCreateConfigVersionUseCase(
	repository entity.ConfigVersionInterface,
) *CreateConfigVersionUseCase {
	return &CreateConfigVersionUseCase{
		ConfigVersionRepository: repository,
	}
}

func (ccu *CreateConfigVersionUseCase) Execute(config outputDTO.ConfigDTO) (outputDTO.ConfigVersionDTO, error) {
	configVersionEntity, err := entity.NewConfigVersion(
		config.Name,
		config.Active,
		config.Frequency,
		config.Service,
		config.Source,
		config.Context,
		ConvertDependsOnDTOToEntity(config.DependsOn),
		config.JobParameters,
		config.ServiceParameters,
		config.ConfigID,
		config.CreatedAt,
		config.UpdatedAt,
	)

	if err != nil {
		return outputDTO.ConfigVersionDTO{}, err
	}

	err = ccu.ConfigVersionRepository.Save(configVersionEntity)
	if err != nil {
		return outputDTO.ConfigVersionDTO{}, err
	}

	dto := outputDTO.ConfigVersionDTO{
		ID:       string(configVersionEntity.ID),
		Versions: ConvertEntityToUseCaseConfigVersion(configVersionEntity.Versions),
	}

	return dto, nil

}
