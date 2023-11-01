package usecase

import (
	"apps/services-orchestration/services-config-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-config-handler/output"
)

type ListOneConfigVersionByIdAndVersionIdUseCase struct {
	ConfigVersionRepository entity.ConfigVersionInterface
}

func NewListOneConfigVersionByIdAndVersionIdUseCase(
	repository entity.ConfigVersionInterface,
) *ListOneConfigVersionByIdAndVersionIdUseCase {
	return &ListOneConfigVersionByIdAndVersionIdUseCase{
		ConfigVersionRepository: repository,
	}
}

func (lcv *ListOneConfigVersionByIdAndVersionIdUseCase) Execute(id string, versionId string) (outputDTO.ConfigDTO, error) {
	item, err := lcv.ConfigVersionRepository.FindOneByIdAndVersionId(id, versionId)
	if err != nil {
		return outputDTO.ConfigDTO{}, err
	}

	dto := outputDTO.ConfigDTO{
		ID:                string(item.ID),
		Name:              item.Name,
		Active:            item.Active,
		Frequency:         item.Frequency,
		Service:           item.Service,
		Source:            item.Source,
		Context:           item.Context,
		DependsOn:         ConvertEntityToUseCaseDependencies(item.DependsOn),
		ConfigID:          item.ConfigID,
		ServiceParameters: item.ServiceParameters,
		JobParameters:     item.JobParameters,
		CreatedAt:         item.CreatedAt,
		UpdatedAt:         item.UpdatedAt,
	}

	return dto, nil
}
