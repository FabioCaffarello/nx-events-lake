package usecase

import (
	apiClient "libs/golang/services/api-clients/services-staging-handler/client"
	stagingHandlerInputDTO "libs/golang/services/dtos/services-staging-handler/input"
	stagingHandlerOutputDTO "libs/golang/services/dtos/services-staging-handler/output"
	"log"
)

type CreateProcessingJobDependenciesUseCase struct {
	ConfigHandlerAPIClient apiClient.Client
}

func NewCreateProcessingJobDependenciesUseCase() *CreateProcessingJobDependenciesUseCase {
	return &CreateProcessingJobDependenciesUseCase{
		ConfigHandlerAPIClient: *apiClient.NewClient(),
	}
}

func (la *CreateProcessingJobDependenciesUseCase) Execute(jobDependencies stagingHandlerInputDTO.ProcessingJobDependenciesDTO) (stagingHandlerOutputDTO.ProcessingJobDependenciesDTO, error) {
	jobDependenciesCreated, err := la.ConfigHandlerAPIClient.CreateProcessingJobDependencies(jobDependencies)
	log.Println("jobDependenciesCreated", jobDependenciesCreated)
	log.Println("err jobDependenciesCreated", err)
	if err != nil {
		return stagingHandlerOutputDTO.ProcessingJobDependenciesDTO{}, err
	}
	return jobDependenciesCreated, nil
}
