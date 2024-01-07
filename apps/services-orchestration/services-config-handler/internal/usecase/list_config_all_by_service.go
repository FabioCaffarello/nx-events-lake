package usecase

import (
	"apps/services-orchestration/services-config-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-config-handler/output"
)

type ListAllConfigsByServiceUseCase struct {
	ConfigRepository entity.ConfigInterface
}

func NewListAllConfigsByServiceUseCase(
	repository entity.ConfigInterface,
) *ListAllConfigsByServiceUseCase {
	return &ListAllConfigsByServiceUseCase{
		ConfigRepository: repository,
	}
}

func (la *ListAllConfigsByServiceUseCase) Execute(service string) ([]outputDTO.ConfigDTO, error) {
	items, err := la.ConfigRepository.FindAllByService(service)
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
