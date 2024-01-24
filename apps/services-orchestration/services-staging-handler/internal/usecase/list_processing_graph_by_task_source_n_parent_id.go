package usecase

import (
    "apps/services-orchestration/services-staging-handler/internal/entity"
    outputDTO "libs/golang/services/dtos/services-staging-handler/output"
)

type ListOneProcessingGraphByTaskSourceAndParentProcessingIdUseCase struct {
    ProcessingGraphRepository entity.ProcessingGraphInterface
}

func NewListOneProcessingGraphByTaskSourceAndParentProcessingIdUseCase(
    repository entity.ProcessingGraphInterface,
) *ListOneProcessingGraphByTaskSourceAndParentProcessingIdUseCase {
    return &ListOneProcessingGraphByTaskSourceAndParentProcessingIdUseCase{
        ProcessingGraphRepository: repository,
    }
}

func (la *ListOneProcessingGraphByTaskSourceAndParentProcessingIdUseCase) Execute(source string, parentProcessingId string) (outputDTO.ProcessingGraphDTO, error) {
    item, err := la.ProcessingGraphRepository.FindOneByTaskSourceAndProcessingId(source, parentProcessingId)
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
