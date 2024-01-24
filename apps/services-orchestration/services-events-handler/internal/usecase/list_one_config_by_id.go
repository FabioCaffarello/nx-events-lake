package usecase

import (
	apiClient "libs/golang/services/api-clients/services-config-handler/client"
	configHandlerInputDTO "libs/golang/services/dtos/services-config-handler/output"
)

type ListOneConfigByIdUseCase struct {
	ConfigHandlerAPIClient apiClient.Client
}

func NewListOneConfigByIdUseCase() *ListOneConfigByIdUseCase {
	return &ListOneConfigByIdUseCase{
		ConfigHandlerAPIClient: *apiClient.NewClient(),
	}
}

func (la *ListOneConfigByIdUseCase) Execute(id string) (configHandlerInputDTO.ConfigDTO, error) {
	config, err := la.ConfigHandlerAPIClient.ListOneConfigById(id)
	if err != nil {
		return configHandlerInputDTO.ConfigDTO{}, err
	}
	return config, nil
}
