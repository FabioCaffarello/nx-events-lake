package usecase

import (
	"apps/services-orchestration/services-staging-handler/internal/entity"
    outputDTO "libs/golang/services/dtos/services-staging-handler/output"
)

type UpdateProcessingGraphTaskStatusUseCase struct {
	ProcessingGraphRepository entity.ProcessingGraphInterface
}

func NewUpdateTaskProcessingStatusGraphUseCase(
	repository entity.ProcessingGraphInterface,
) *UpdateProcessingGraphTaskStatusUseCase {
	return &UpdateProcessingGraphTaskStatusUseCase{
		ProcessingGraphRepository: repository,
	}
}

func (la *UpdateProcessingGraphTaskStatusUseCase) Execute(source string, processingId string, statusCode int, processingTimestamp string) (outputDTO.ProcessingGraphDTO, error) {
	item, err := la.ProcessingGraphRepository.UpdateTaskStatus(source, processingId, statusCode)
	if err != nil {
		return outputDTO.ProcessingGraphDTO{}, err
	}
    dto := outputDTO.ProcessingGraphDTO{
        ID:                string(item.ID),
        Source:            item.Source,
        Context:           item.Context,
        StartProcessingId: item.StartProcessingId,
        Tasks:             ConvertEntityToUsecaseTasksWithProcessingTimestamp(item.Tasks, processingTimestamp),
        CreatedAt:         item.CreatedAt,
        UpdatedAt:         item.UpdatedAt,
    }
    return dto, nil
}
