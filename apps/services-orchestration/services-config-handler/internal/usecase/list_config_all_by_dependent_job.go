package usecase

import (
	"apps/services-orchestration/services-config-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-config-handler/output"
)

type ListAllConfigsByDependentJobUseCase struct {
	ConfigRepository entity.ConfigInterface
}

func NewListAllConfigsByDependentJobUseCase(
	repository entity.ConfigInterface,
) *ListAllConfigsByDependentJobUseCase {
	return &ListAllConfigsByDependentJobUseCase{
		ConfigRepository: repository,
	}
}

func (la *ListAllConfigsByDependentJobUseCase) Execute(service string, source string) ([]outputDTO.ConfigDTO, error) {
	items, err := la.ConfigRepository.FindAllByDependentJod(service, source)
	if err != nil {
		return []outputDTO.ConfigDTO{}, err
	}
	var result []outputDTO.ConfigDTO
	for _, item := range items {
		dto := outputDTO.ConfigDTO{
			ID:                string(item.ID),
			Name:              item.Name,
			Active:            item.Active,
			Frequency:         item.Frequency,
			Service:           item.Service,
			Source:            item.Source,
			Context:           item.Context,
            InputMethod:       item.InputMethod,
            OutputMethod:      item.OutputMethod,
			DependsOn:         ConvertEntityToUseCaseDependencies(item.DependsOn),
			ConfigID:          item.ConfigID,
			ServiceParameters: item.ServiceParameters,
			JobParameters:     item.JobParameters,
			CreatedAt:         item.CreatedAt,
			UpdatedAt:         item.UpdatedAt,
		}
		result = append(result, dto)
	}
	return result, nil
}
