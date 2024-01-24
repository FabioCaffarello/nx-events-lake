package usecase

import (
	apiClient "libs/golang/services/api-clients/services-config-handler/client"
	configHandlerInputDTO "libs/golang/services/dtos/services-config-handler/output"
)

type ListAllConfigsByDependentJobUseCase struct {
	ConfigHandlerAPIClient apiClient.Client
}

func NewListAllConfigsByDependentJobUseCase() *ListAllConfigsByDependentJobUseCase {
	return &ListAllConfigsByDependentJobUseCase{
		ConfigHandlerAPIClient: *apiClient.NewClient(),
	}
}

func (la *ListAllConfigsByDependentJobUseCase) Execute(service string, source string) ([]configHandlerInputDTO.ConfigDTO, error) {
	configs, err := la.ConfigHandlerAPIClient.ListAllConfigsByDependentJob(service, source)
	if err != nil {
		return []configHandlerInputDTO.ConfigDTO{}, err
	}
	return configs, nil
}
