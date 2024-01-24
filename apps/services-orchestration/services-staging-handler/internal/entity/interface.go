package entity

type ProcessingJobDependenciesInterface interface {
	Save(processingJobDependencies *ProcessingJobDependencies) error
	UpdateProcessingJobDependencies(jobDep *JobDependenciesWithProcessingData, id string) error
	Delete(id string) error
	FindOneById(id string) (*ProcessingJobDependencies, error)
}

type ProcessingGraphInterface interface {
    Save(processingGraph *ProcessingGraph) error
    FindOneById(id string) (*ProcessingGraph, error)
    Delete(id string) error
    FindOneBySourceAndStartProcessingId(source string, startProcessingId string) (*ProcessingGraph, error)
    FindOneByTaskSourceAndProcessingId(source string, parentProcessingId string) (*ProcessingGraph, error)
    CreateTask(source string, startProcessingId string, task *Task) (*ProcessingGraph, error)
    UpdateTaskStatus(source string, processingId string, statusCode int) (*ProcessingGraph, error)
    UpdateTaskOutput(source string, processingId string, outputId string) (*ProcessingGraph, error)
}
