package usecase

import (
	"apps/services-orchestration/services-config-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-config-handler/output"
)

type ListOneConfigByServiceAndSourceAndContextUseCase struct {
	ConfigRepository entity.ConfigInterface
}

func NewListOneConfigByServiceAndSourceAndContextUseCase(
	repository entity.ConfigInterface,
) *ListOneConfigByServiceAndSourceAndContextUseCase {
	return &ListOneConfigByServiceAndSourceAndContextUseCase{
		ConfigRepository: repository,
	}
}

func (la *ListOneConfigByServiceAndSourceAndContextUseCase) Execute(service string, source string, context string) (outputDTO.ConfigDTO, error) {
	item, err := la.ConfigRepository.FindOneByServiceAndSourceAndContext(service, source, context)
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