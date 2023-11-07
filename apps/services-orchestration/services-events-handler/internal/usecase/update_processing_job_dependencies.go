package usecase

import (
	apiClient "libs/golang/services/api-clients/services-staging-handler/client"
	stagingHandlerOutputDTO "libs/golang/services/dtos/services-staging-handler/output"
	stagingHandlerSharedDTO "libs/golang/services/dtos/services-staging-handler/shared"
)

type UpdateProcessingJobDependenciesUseCase struct {
	ConfigHandlerAPIClient apiClient.Client
}

func NewUpdateProcessingJobDependenciesUseCase() *UpdateProcessingJobDependenciesUseCase {
	return &UpdateProcessingJobDependenciesUseCase{
		ConfigHandlerAPIClient: *apiClient.NewClient(),
	}
}

func (la *UpdateProcessingJobDependenciesUseCase) Execute(id string, jobDependencies stagingHandlerSharedDTO.ProcessingJobDependencies) (stagingHandlerOutputDTO.ProcessingJobDependenciesDTO, error) {
	jobDependenciesUpdated, err := la.ConfigHandlerAPIClient.UpdateProcessingJobDependencies(id, jobDependencies)
	if err != nil {
		return stagingHandlerOutputDTO.ProcessingJobDependenciesDTO{}, err
	}
	return jobDependenciesUpdated, nil
}
