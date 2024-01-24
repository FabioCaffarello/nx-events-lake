package usecase

import (
	"apps/services-orchestration/services-staging-handler/internal/entity"
	sharedDTO "libs/golang/services/dtos/services-staging-handler/shared"
	outputDTO "libs/golang/services/dtos/services-staging-handler/output"
    inputDTO "libs/golang/services/dtos/services-staging-handler/input"
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

func ConvertEntityToUsecaseTasks(entityTasks []entity.Task) []outputDTO.Task {
    usecaseTasks := make([]outputDTO.Task, len(entityTasks))
    for i, task := range entityTasks {
        usecaseTasks[i] = outputDTO.Task{
            Service:               task.Service,
            Source:                task.Source,
            ProcessingId:          task.ProcessingId,
            ParentProcessingId:    task.ParentProcessingId,
            ConfigVersionId:       task.ConfigVersionId,
            InputSchemaVersionId:  task.InputSchemaVersionId,
            OutputSchemaVersionId: task.OutputSchemaVersionId,
            InputId:               task.InputId,
            OutputId:              task.OutputId,
            StatusCode:            task.StatusCode,
            ProcessingTimestamp:   task.ProcessingTimestamp,
        }
    }
    return usecaseTasks
}

func ConvertEntityToUsecaseTasksWithProcessingTimestamp(entityTasks []entity.Task, processingTimestamp string) []outputDTO.Task {
    usecaseTasks := make([]outputDTO.Task, len(entityTasks))
    for i, task := range entityTasks {
        usecaseTasks[i] = outputDTO.Task{
            Service:               task.Service,
            Source:                task.Source,
            ProcessingId:          task.ProcessingId,
            ParentProcessingId:    task.ParentProcessingId,
            ConfigVersionId:       task.ConfigVersionId,
            InputSchemaVersionId:  task.InputSchemaVersionId,
            OutputSchemaVersionId: task.OutputSchemaVersionId,
            InputId:               task.InputId,
            OutputId:              task.OutputId,
            StatusCode:            task.StatusCode,
            ProcessingTimestamp:   processingTimestamp,
        }
    }
    return usecaseTasks
}

func ConvertUsecaseToEntityTask(usecaseTasks inputDTO.Task) *entity.Task {
    entityTask := entity.Task{
        Service:               usecaseTasks.Service,
        Source:                usecaseTasks.Source,
        ProcessingId:          usecaseTasks.ProcessingId,
        ParentProcessingId:    usecaseTasks.ParentProcessingId,
        ConfigVersionId:       usecaseTasks.ConfigVersionId,
        InputSchemaVersionId:  usecaseTasks.InputSchemaVersionId,
        OutputSchemaVersionId: usecaseTasks.OutputSchemaVersionId,
        InputId:               usecaseTasks.InputId,
        OutputId:              usecaseTasks.OutputId,
        StatusCode:            usecaseTasks.StatusCode,
        ProcessingTimestamp:   usecaseTasks.ProcessingTimestamp,
    }
    return &entityTask
}
