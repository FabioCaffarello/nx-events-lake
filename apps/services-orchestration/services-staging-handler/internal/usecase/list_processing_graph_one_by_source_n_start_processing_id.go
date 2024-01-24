package usecase

import (
	"apps/services-orchestration/services-staging-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-staging-handler/output"
)

type ListOneProcessingGraphBySourceAndStartProcessingIdUseCase struct {
	ProcessingGraphRepository entity.ProcessingGraphInterface
}

func NewListOneProcessingGraphBySourceAndStartProcessingIdUseCase(
	repository entity.ProcessingGraphInterface,
) *ListOneProcessingGraphBySourceAndStartProcessingIdUseCase {
	return &ListOneProcessingGraphBySourceAndStartProcessingIdUseCase{
		ProcessingGraphRepository: repository,
	}
}

func (la *ListOneProcessingGraphBySourceAndStartProcessingIdUseCase) Execute(source string, startProcessingId string) (outputDTO.ProcessingGraphDTO, error) {
	item, err := la.ProcessingGraphRepository.FindOneBySourceAndStartProcessingId(source, startProcessingId)
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
