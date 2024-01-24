package usecase

import (
	apiClient "libs/golang/services/api-clients/services-staging-handler/client"
	stagingHandlerInputDTO "libs/golang/services/dtos/services-staging-handler/input"
	stagingHandlerOutputDTO "libs/golang/services/dtos/services-staging-handler/output"
)

type CreateProcessingJobDependenciesUseCase struct {
	StagingHandlerAPIClient apiClient.Client
}

func NewCreateProcessingJobDependenciesUseCase() *CreateProcessingJobDependenciesUseCase {
	return &CreateProcessingJobDependenciesUseCase{
		StagingHandlerAPIClient: *apiClient.NewClient(),
	}
}

func (la *CreateProcessingJobDependenciesUseCase) Execute(jobDependencies stagingHandlerInputDTO.ProcessingJobDependenciesDTO) (stagingHandlerOutputDTO.ProcessingJobDependenciesDTO, error) {
	jobDependenciesCreated, err := la.StagingHandlerAPIClient.CreateProcessingJobDependencies(jobDependencies)
	if err != nil {
		return stagingHandlerOutputDTO.ProcessingJobDependenciesDTO{}, err
	}
	return jobDependenciesCreated, nil
}
