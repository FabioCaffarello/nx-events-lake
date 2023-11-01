package usecase

import (
	"apps/services-orchestration/services-staging-handler/internal/entity"
	sharedDTO "libs/golang/services/dtos/services-staging-handler/shared"
)

func ConvertEntityToUsecaseJobDependenciesWithProcessingData(entityDeps []entity.JobDependenciesWithProcessingData) []sharedDTO.ProcessingJobDependencies {
	usecaseDeps := make([]sharedDTO.ProcessingJobDependencies, len(entityDeps))
	for i, dep := range entityDeps {
		usecaseDeps[i] = sharedDTO.ProcessingJobDependencies{
			Service:             dep.Service,
			Source:              dep.Source,
			ProcessingId:        dep.ProcessingId,
			ProcessingTimestamp: dep.ProcessingTimestamp,
			StatusCode:          dep.StatusCode,
		}
	}
	return usecaseDeps
}
