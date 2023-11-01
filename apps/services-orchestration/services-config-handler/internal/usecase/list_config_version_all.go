package usecase

import (
	"apps/services-orchestration/services-config-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-config-handler/output"
)

type ListAllConfigsVersionUseCase struct {
	ConfigVersionRepository entity.ConfigVersionInterface
}

func NewListAllConfigsVersionUseCase(
	repository entity.ConfigVersionInterface,
) *ListAllConfigsVersionUseCase {
	return &ListAllConfigsVersionUseCase{
		ConfigVersionRepository: repository,
	}
}

func (la *ListAllConfigsVersionUseCase) Execute() ([]outputDTO.ConfigVersionDTO, error) {
	items, err := la.ConfigVersionRepository.FindAll()
	if err != nil {
		return []outputDTO.ConfigVersionDTO{}, err
	}
	var result []outputDTO.ConfigVersionDTO
	for _, item := range items {
		dto := outputDTO.ConfigVersionDTO{
			ID:       string(item.ID),
			Versions: ConvertEntityToUseCaseConfigVersion(item.Versions),
		}
		result = append(result, dto)
	}
	return result, nil
}
