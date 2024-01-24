package input

import (
	sharedDTO "libs/golang/services/dtos/services-staging-handler/shared"
)

type ProcessingJobDependenciesDTO struct {
	Service               string                                `json:"service"`
	Source                string                                `json:"source"`
	Context               string                                `json:"context"`
	ParentJobProcessingId string                                `json:"parent_job_processing_id"`
	JobDependencies       []sharedDTO.ProcessingJobDependencies `json:"job_dependencies"`
}

type ProcessingGraphDTO struct {
    Context string `json:"context"`
    Source string `json:"source"`
    StartProcessingId string `json:"start_processing_id"`
}

type Task struct {
	Source                string `json:"source"`
	Service               string `json:"service"`
	ProcessingId          string `json:"processing_id"`
	ParentProcessingId    string `json:"parent_processing_id"`
	ConfigVersionId       string `json:"config_version_id"`
	InputSchemaVersionId  string `json:"input_schema_version_id"`
	OutputSchemaVersionId string `json:"output_schema_version_id"`
	InputId               string `json:"input_id"`
	OutputId              string `json:"output_id"`
	StatusCode            int    `json:"status_code"`
	ProcessingTimestamp   string `json:"processing_timestamp"`
}
