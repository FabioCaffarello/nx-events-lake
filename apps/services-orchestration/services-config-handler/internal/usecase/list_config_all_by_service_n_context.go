package usecase

import (
	"apps/services-orchestration/services-config-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-config-handler/output"
)

type ListAllConfigsByServiceAndContextUseCase struct {
	ConfigRepository entity.ConfigInterface
}

func NewListAllConfigsByServiceAndContextUseCase(
	repository entity.ConfigInterface,
) *ListAllConfigsByServiceAndContextUseCase {
	return &ListAllConfigsByServiceAndContextUseCase{
		ConfigRepository: repository,
	}
}

func (la *ListAllConfigsByServiceAndContextUseCase) Execute(service string, contextEnv string) ([]outputDTO.ConfigDTO, error) {
	items, err := la.ConfigRepository.FindAllByServiceAndContext(service, contextEnv)
	if err != nil {
		return nil, err
	}

	var output []outputDTO.ConfigDTO
	for _, item := range items {
		output = append(output, outputDTO.ConfigDTO{
			ID:                item.ID,
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
		})
	}

	return output, nil
}
