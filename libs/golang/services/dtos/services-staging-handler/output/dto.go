package output

import (
	sharedDTO "libs/golang/services/dtos/services-staging-handler/shared"
)

type ProcessingJobDependenciesDTO struct {
	ID                    string                                `json:"id"`
	Service               string                                `json:"service"`
	Source                string                                `json:"source"`
	Context               string                                `json:"context"`
	ParentJobProcessingId string                                `json:"parent_job_processing_id"`
	JobDependencies       []sharedDTO.ProcessingJobDependencies `json:"job_dependencies"`
}
