package usecase

import (
	apiClient "libs/golang/services/api-clients/services-staging-handler/client"
	stagingHandlerOutputDTO "libs/golang/services/dtos/services-staging-handler/output"
)

type RemoveProcessingJobDependenciesUseCase struct {
	ConfigHandlerAPIClient apiClient.Client
}

func NewRemoveProcessingJobDependenciesUseCase() *RemoveProcessingJobDependenciesUseCase {
	return &RemoveProcessingJobDependenciesUseCase{
		ConfigHandlerAPIClient: *apiClient.NewClient(),
	}
}

func (la *RemoveProcessingJobDependenciesUseCase) Execute(id string) (stagingHandlerOutputDTO.ProcessingJobDependenciesDTO, error) {
	_, err := la.ConfigHandlerAPIClient.RemoveProcessingJobDependencies(id)
	if err != nil {
		return stagingHandlerOutputDTO.ProcessingJobDependenciesDTO{}, err
	}
	return stagingHandlerOutputDTO.ProcessingJobDependenciesDTO{}, nil
}
