package usecase

import (
	apiClient "libs/golang/services/api-clients/services-staging-handler/client"
	stagingHandlerOutputDTO "libs/golang/services/dtos/services-staging-handler/output"
	stagingHandlerSharedDTO "libs/golang/services/dtos/services-staging-handler/shared"
	"log"
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
    log.Println("jobDependenciesUpdated", jobDependenciesUpdated)
    log.Println("err", err)
    log.Println("len(jobDependenciesUpdated.JobDependencies)", len(jobDependenciesUpdated.JobDependencies))
	if err != nil {
		return stagingHandlerOutputDTO.ProcessingJobDependenciesDTO{}, err
	}
	return jobDependenciesUpdated, nil
}
