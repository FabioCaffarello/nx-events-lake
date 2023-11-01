package usecase

import (
	"apps/services-orchestration/services-staging-handler/internal/entity"
	sharedDTO "libs/golang/services/dtos/services-staging-handler/shared"
	"log"
)

type UpdateProcessingJobDependenciesUseCase struct {
	ProcessingJobDependenciesRepository entity.ProcessingJobDependenciesInterface
}

func NewUpdateProcessingJobDependenciesUseCase(
	repository entity.ProcessingJobDependenciesInterface,
) *UpdateProcessingJobDependenciesUseCase {
	return &UpdateProcessingJobDependenciesUseCase{
		ProcessingJobDependenciesRepository: repository,
	}
}

func (u *UpdateProcessingJobDependenciesUseCase) Execute(jobDep sharedDTO.ProcessingJobDependencies, id string) error {
	entityJobDep := &entity.JobDependenciesWithProcessingData{
		Service:             jobDep.Service,
		Source:              jobDep.Source,
		ProcessingId:        jobDep.ProcessingId,
		ProcessingTimestamp: jobDep.ProcessingTimestamp,
		StatusCode:          jobDep.StatusCode,
	}
	log.Printf("UpdateProcessingJobDependenciesUseCase.Execute: entityJobDep=%v, id=%s", entityJobDep, id)

	err := u.ProcessingJobDependenciesRepository.UpdateProcessingJobDependencies(entityJobDep, id)
	if err != nil {
		return err
	}
	return nil
}
