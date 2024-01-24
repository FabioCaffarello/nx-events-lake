package usecase

import (
	"apps/services-orchestration/services-staging-handler/internal/entity"
	inputDTO "libs/golang/services/dtos/services-staging-handler/input"
	outputDTO "libs/golang/services/dtos/services-staging-handler/output"
)

type CreateProcessingGraphUseCase struct {
	ProcessingGraphRepository entity.ProcessingGraphInterface
}

func NewCreateProcessingGraphUseCase(
	ProcessingGraphRepository entity.ProcessingGraphInterface,
) *CreateProcessingGraphUseCase {
	return &CreateProcessingGraphUseCase{
		ProcessingGraphRepository: ProcessingGraphRepository,
	}
}

func (c *CreateProcessingGraphUseCase) Execute(input inputDTO.ProcessingGraphDTO) (outputDTO.ProcessingGraphDTO, error) {
	processingGraphEntity, err := entity.NewProcessingGraph(
		input.Context,
		input.Source,
		input.StartProcessingId,
	)
	if err != nil {
		return outputDTO.ProcessingGraphDTO{}, err
	}

	err = c.ProcessingGraphRepository.Save(processingGraphEntity)
	if err != nil {
		return outputDTO.ProcessingGraphDTO{}, err
	}

	dto := outputDTO.ProcessingGraphDTO{
		ID:                string(processingGraphEntity.ID),
		Source:            processingGraphEntity.Source,
		Context:           processingGraphEntity.Context,
		StartProcessingId: processingGraphEntity.StartProcessingId,
        CreatedAt:         processingGraphEntity.CreatedAt,
        UpdatedAt:         processingGraphEntity.UpdatedAt,
	}

	return dto, nil
}
