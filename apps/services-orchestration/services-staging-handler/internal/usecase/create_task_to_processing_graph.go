package usecase

import (
    "apps/services-orchestration/services-staging-handler/internal/entity"
    inputDTO "libs/golang/services/dtos/services-staging-handler/input"
    outputDTO "libs/golang/services/dtos/services-staging-handler/output"
)

type CreateTaskToProcessingGraphUseCase struct {
    ProcessingGraphRepository entity.ProcessingGraphInterface
}

func NewCreateTaskToProcessingGraphUseCase(
    repository entity.ProcessingGraphInterface,
) *CreateTaskToProcessingGraphUseCase {
    return &CreateTaskToProcessingGraphUseCase{
        ProcessingGraphRepository: repository,
    }
}

func (la *CreateTaskToProcessingGraphUseCase) Execute(source string, startProcessingId string, task inputDTO.Task) (outputDTO.ProcessingGraphDTO, error) {
    entityTask := ConvertUsecaseToEntityTask(task)
    item, err := la.ProcessingGraphRepository.CreateTask(source, startProcessingId, entityTask)
    if err != nil {
        return  outputDTO.ProcessingGraphDTO{}, err
    }
    dto := outputDTO.ProcessingGraphDTO{
        ID:                string(item.ID),
        Source:            item.Source,
        Context:           item.Context,
        StartProcessingId: item.StartProcessingId,
        Tasks:             ConvertEntityToUsecaseTasks(item.Tasks),
        CreatedAt:         item.CreatedAt,
        UpdatedAt:         item.UpdatedAt,
    }
    return dto, nil
}
