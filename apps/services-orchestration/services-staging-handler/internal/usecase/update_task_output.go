package usecase

import (
    "apps/services-orchestration/services-staging-handler/internal/entity"
    outputDTO "libs/golang/services/dtos/services-staging-handler/output"
)

type UpdateProcessingGraphTaskOutputUseCase struct {
    ProcessingGraphRepository entity.ProcessingGraphInterface
}

func NewUpdateProcessingGraphTaskOutputUseCase(
    repository entity.ProcessingGraphInterface,
) *UpdateProcessingGraphTaskOutputUseCase {
    return &UpdateProcessingGraphTaskOutputUseCase{
        ProcessingGraphRepository: repository,
    }
}

func (la *UpdateProcessingGraphTaskOutputUseCase) Execute(source string, processingId string, outputId string) (outputDTO.ProcessingGraphDTO, error) {
    item, err := la.ProcessingGraphRepository.UpdateTaskOutput(source, processingId, outputId)
    if err != nil {
        return outputDTO.ProcessingGraphDTO{}, err
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


