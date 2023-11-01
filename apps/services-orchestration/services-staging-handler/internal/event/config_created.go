package event

import "time"

type ProcessingJobDependenciesCreated struct {
	Name    string
	Payload interface{}
}

func NewProcessingJobDependenciesCreated() *ProcessingJobDependenciesCreated {
	return &ProcessingJobDependenciesCreated{
		Name: "ProcessingJobDependenciesCreated",
	}
}

func (e *ProcessingJobDependenciesCreated) GetName() string {
	return e.Name
}

func (e *ProcessingJobDependenciesCreated) GetPayload() interface{} {
	return e.Payload
}

func (e *ProcessingJobDependenciesCreated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *ProcessingJobDependenciesCreated) GetDateTime() time.Time {
	return time.Now()
}
