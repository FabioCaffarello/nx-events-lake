package entity

import (
	"errors"
	pipeid "libs/golang/shared/go-id/pipeline/md5"
	"time"
)

var (
	ErrPipelineFrequencyEmpty    = errors.New("pipeline frequency is empty")
	ErrPipelineContextEmpty      = errors.New("pipeline context is empty")
	ErrPipelineSourceEmpty       = errors.New("pipeline source is empty")
	ErrPipelineServiceStartEmpty = errors.New("pipeline service start is empty")
	ErrPipelineProcessingIdEmpty = errors.New("pipeline processing id is empty")
	ErrPipelineTasksEmpty        = errors.New("pipeline tasks is empty")
)

type ChildJob struct {
	Source  string `bson:"source"`
	Service string `bson:"service"`
	Kind    string `bson:"kind"`
}

type Task struct {
	Source                 string     `bson:"source"`
	Service                string     `bson:"service"`
	ProcessingId           string     `bson:"processing_id"`
	InputSchemaVersionId   string     `bson:"input_schema_version_id"`
	OutputSchemaVersionId  string     `bson:"output_schema_version_id"`
	JobDefinitionVersionId string     `bson:"job_definition_version_id"`
	InputId                string     `bson:"input_id"`
	OutputId               string     `bson:"output_id"`
	StatusCode             int        `bson:"status_code"`
	ChildJobs              []ChildJob `bson:"child_jobs"`
}

type ProcessingGraph struct {
	ID           pipeid.ID `bson:"id"`
	Frequency    string    `bson:"frequency"`
	Context      string    `bson:"context"`
	Source       string    `bson:"source"`
	ServiceStart string    `bson:"service_start"`
	ProcessingId string    `bson:"processing_id"`
	Tasks        []Task    `bson:"tasks"`
	CreatedAt    string    `bson:"created_at"`
	UpdatedAt    string    `bson:"updated_at"`
}

func NewPipeline(
	frequency string,
	context string,
	source string,
	serviceStart string,
	ProcessingId string,
	tasks []Task,
) (*ProcessingGraph, error) {
	pipeline := &ProcessingGraph{
		ID:           pipeid.NewPipelineGraphID(context, serviceStart, source, ProcessingId),
		Frequency:    frequency,
		Context:      context,
		Source:       source,
		ServiceStart: serviceStart,
		Tasks:        tasks,
		CreatedAt:    time.Now().Format(time.RFC3339),
		UpdatedAt:    time.Now().Format(time.RFC3339),
	}
	err := pipeline.IsPipelineValid()
	if err != nil {
		return nil, err
	}
	return pipeline, nil
}

func (p *ProcessingGraph) IsPipelineValid() error {
	if p.Frequency == "" {
		return ErrPipelineFrequencyEmpty
	}
	if p.Context == "" {
		return ErrPipelineContextEmpty
	}
	if p.Source == "" {
		return ErrPipelineSourceEmpty
	}
	if p.ServiceStart == "" {
		return ErrPipelineServiceStartEmpty
	}
	if p.ProcessingId == "" {
		return ErrPipelineProcessingIdEmpty
	}
	if len(p.Tasks) == 0 {
		return ErrPipelineTasksEmpty
	}
	return nil
}
