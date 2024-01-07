package entity

import (
	"errors"
	"fmt"
	"libs/golang/shared/go-id/md5"
)

var (
	ErrProcessingJobDependenciesInvalid                    = errors.New("processing Job Dependencies is invalid")
	ErrProcessingJobDependenciesServiceEmpty               = errors.New("processing Job Dependencies service is empty")
	ErrProcessingJobDependenciesSourceEmpty                = errors.New("processing Job Dependencies source is empty")
	ErrProcessingJobDependenciesContextEmpty               = errors.New("processing Job Dependencies context is empty")
	ErrProcessingJobDependenciesParentJobProcessingIdEmpty = errors.New("processing Job Dependencies parent job processing id is empty")
)

type JobDependenciesWithProcessingData struct {
	Service             string `bson:"service"`
	Source              string `bson:"source"`
	ProcessingId        string `bson:"processing_id"`
	ProcessingTimestamp string `bson:"processing_timestamp"`
	StatusCode          int    `bson:"status_code"`
}

type ProcessingJobDependencies struct {
	ID                    md5.ID                              `bson:"id"`
	Service               string                              `bson:"service"`
	Source                string                              `bson:"source"`
	Context               string                              `bson:"context"`
	ParentJobProcessingId string                              `bson:"parent_job_processing_id"`
	JobDependencies       []JobDependenciesWithProcessingData `bson:"job_dependencies"`
}

func NewProcessingJobDependencies(
	service string,
	source string,
	context string,
	parentJobProcessingId string,
	jobDependencies []JobDependenciesWithProcessingData,
) (*ProcessingJobDependencies, error) {
	processingJobDependencies := &ProcessingJobDependencies{
		ID:                    md5.NewMd5Hash(fmt.Sprintf("%s-%s-%s-%s", context, service, source, parentJobProcessingId)),
		Service:               service,
		Source:                source,
		Context:               context,
		ParentJobProcessingId: parentJobProcessingId,
		JobDependencies:       jobDependencies,
	}

	err := processingJobDependencies.IsProcessingJobDependenciesValid()
	if err != nil {
		return nil, err
	}
	return processingJobDependencies, nil
}

func (p *ProcessingJobDependencies) IsProcessingJobDependenciesValid() error {
	if p.Service == "" {
		return ErrProcessingJobDependenciesServiceEmpty
	}
	if p.Source == "" {
		return ErrProcessingJobDependenciesSourceEmpty
	}
	if p.ParentJobProcessingId == "" {
		return ErrProcessingJobDependenciesParentJobProcessingIdEmpty
	}
	if p.Context == "" {
		return ErrProcessingJobDependenciesContextEmpty
	}
	if len(p.JobDependencies) == 0 {
		return ErrProcessingJobDependenciesInvalid
	}
	return nil
}
