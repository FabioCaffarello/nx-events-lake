package input

import (
	sharedDTO "libs/golang/services/dtos/services-staging-handler/shared"
)

type ProcessingJobDependenciesDTO struct {
	Service         string                                `json:"service"`
	Source          string                                `json:"source"`
    Context         string                                `json:"context"`
	JobDependencies []sharedDTO.ProcessingJobDependencies `json:"job_dependencies"`
}
