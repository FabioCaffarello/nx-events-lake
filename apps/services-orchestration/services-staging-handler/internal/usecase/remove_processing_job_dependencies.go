package usecase

import (
	"apps/services-orchestration/services-staging-handler/internal/entity"
)

type RemoveProcessingJobDependenciesUseCase struct {
	ProcessingJobDependenciesRepository entity.ProcessingJobDependenciesInterface
}

func NewRemoveProcessingJobDependenciesUseCase(
	repository entity.ProcessingJobDependenciesInterface,
) *RemoveProcessingJobDependenciesUseCase {
	return &RemoveProcessingJobDependenciesUseCase{
		ProcessingJobDependenciesRepository: repository,
	}
}

func (c *RemoveProcessingJobDependenciesUseCase) Execute(id string) error {
	err := c.ProcessingJobDependenciesRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
