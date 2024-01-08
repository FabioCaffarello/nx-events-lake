package entity

import (
	"errors"
	pipeid "libs/golang/shared/go-id/pipeline/id"
	pipeuuid "libs/golang/shared/go-id/pipeline/uuid"
	"time"
)

var (
	ErrPipelineFrequencyEmpty         = errors.New("pipeline frequency is empty")
	ErrPipelineContextEmpty           = errors.New("pipeline context is empty")
	ErrPipelineSourceEmpty            = errors.New("pipeline source is empty")
	ErrPipelineServiceStartEmpty      = errors.New("pipeline service start is empty")
	ErrPipelinePipelineVersionIdEmpty = errors.New("pipeline pipeline version id is empty")
	ErrPipelineTasksEmpty             = errors.New("pipeline tasks is empty")
	ErrPipelineNil                    = errors.New("pipeline is nil")
	ErrGeneratePipelineID             = errors.New("error generating pipeline id")
)

type ChildJob struct {
	Source  string `bson:"source"`
	Service string `bson:"service"`
	Kind    string `bson:"kind"`
}

type Task struct {
	Source    string     `bson:"source"`
	Service   string     `bson:"service"`
	ChildJobs []ChildJob `bson:"child_jobs"`
}

type Pipeline struct {
	ID                pipeid.ID   `bson:"id"`
	Frequency         string      `bson:"frequency"`
	Context           string      `bson:"context"`
	Source            string      `bson:"source"`
	ServiceStart      string      `bson:"service_start"`
	PipelineVersionId pipeuuid.ID `bson:"pipeline_version_id"`
	Tasks             []Task      `bson:"tasks"`
	CreatedAt         string      `bson:"created_at"`
	UpdatedAt         string      `bson:"updated_at"`
}

func NewPipeline(
	frequency string,
	context string,
	source string,
	serviceStart string,
	pipelineVersionId string,
	tasks []Task,
) (*Pipeline, error) {
	pipeline := &Pipeline{
		ID:                pipeid.NewID(context, serviceStart, source),
		Frequency:         frequency,
		Context:           context,
		Source:            source,
		ServiceStart:      serviceStart,
		PipelineVersionId: pipelineVersionId,
		Tasks:             tasks,
		CreatedAt:         time.Now().Format(time.RFC3339),
		UpdatedAt:         time.Now().Format(time.RFC3339),
	}
	err := pipeline.IsPipelineValid()
	if err != nil {
		return nil, err
	}
    err = pipeline.SetPipelineID()
    if err != nil {
        return nil, err
    }

	return pipeline, nil
}

func convertChildJobs(childJobs []ChildJob) []map[string]interface{} {
	mapChildJobs := make([]map[string]interface{}, len(childJobs))
	for i, childJob := range childJobs {
		mapChildJobs[i] = map[string]interface{}{
			"source":  childJob.Source,
			"service": childJob.Service,
			"kind":    childJob.Kind,
		}
	}
	return mapChildJobs
}

func convertTasks(tasks []Task) []map[string]interface{} {
	mapTasks := make([]map[string]interface{}, len(tasks))
	for i, task := range tasks {
		mapTasks[i] = map[string]interface{}{
			"service":    task.Service,
			"source":     task.Source,
			"child_jobs": convertChildJobs(task.ChildJobs),
		}
	}
	return mapTasks
}

func getPipelineParamsToGenerateVersionID(p *Pipeline) map[string]interface{} {
	return map[string]interface{}{
		"frequency":    p.Frequency,
		"context":      p.Context,
		"source":       p.Source,
		"serviceStart": p.ServiceStart,
		"tasks":        convertTasks(p.Tasks),
	}
}

func (p *Pipeline) SetPipelineID() error {
	if p == nil {
		return ErrPipelineNil
	}
	pipelineDataForVersionID := getPipelineParamsToGenerateVersionID(p)
	pipelineID, err := pipeuuid.GeneratePipelineID(pipelineDataForVersionID)
	if err != nil {
		return ErrGeneratePipelineID
	}
	p.PipelineVersionId = pipelineID
	return nil

}

func (p *Pipeline) IsPipelineValid() error {
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
	if p.PipelineVersionId == "" {
		return ErrPipelinePipelineVersionIdEmpty
	}
	if len(p.Tasks) == 0 {
		return ErrPipelineTasksEmpty
	}
	return nil
}
