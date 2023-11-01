package entity

type ProcessingJobDependenciesInterface interface {
	Save(processingJobDependencies *ProcessingJobDependencies) error
	UpdateProcessingJobDependencies(jobDep *JobDependenciesWithProcessingData, id string) error
	Delete(id string) error
	FindOneById(id string) (*ProcessingJobDependencies, error)
}
