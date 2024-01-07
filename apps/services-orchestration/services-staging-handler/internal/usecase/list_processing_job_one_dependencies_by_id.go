package usecase

import (
	"apps/services-orchestration/services-staging-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-staging-handler/output"
)

type ListOneProcessingJobDependenciesByIdUseCase struct {
	ProcessingJobDependenciesRepository entity.ProcessingJobDependenciesInterface
}

func NewListOneProcessingJobDependenciesByIdUseCase(
	repository entity.ProcessingJobDependenciesInterface,
) *ListOneProcessingJobDependenciesByIdUseCase {
	return &ListOneProcessingJobDependenciesByIdUseCase{
		ProcessingJobDependenciesRepository: repository,
	}
}

func (la *ListOneProcessingJobDependenciesByIdUseCase) Execute(id string) (outputDTO.ProcessingJobDependenciesDTO, error) {
	item, err := la.ProcessingJobDependenciesRepository.FindOneById(id)
	if err != nil {
		return outputDTO.ProcessingJobDependenciesDTO{}, err
	}
	dto := outputDTO.ProcessingJobDependenciesDTO{
		ID:                    string(item.ID),
		Service:               item.Service,
		Source:                item.Source,
		Context:               item.Context,
		ParentJobProcessingId: item.ParentJobProcessingId,
		JobDependencies:       ConvertEntityToUsecaseJobDependenciesWithProcessingData(item.JobDependencies),
	}
	return dto, nil
}
