package usecase

import (
	"apps/services-orchestration/services-input-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-input-handler/output"
	sharedDTO "libs/golang/services/dtos/services-input-handler/shared"
)

type ListAllByServiceAndSourceUseCase struct {
	InputRepository entity.InputInterface
}

func NewListAllByServiceAndSourceUseCase(
	repository entity.InputInterface,
) *ListAllByServiceAndSourceUseCase {
	return &ListAllByServiceAndSourceUseCase{
		InputRepository: repository,
	}
}

func (la *ListAllByServiceAndSourceUseCase) Execute(service string, source string) ([]outputDTO.InputDTO, error) {
	items, err := la.InputRepository.FindAllByServiceAndSource(service, source)
	if err != nil {
		return []outputDTO.InputDTO{}, err
	}
	var result []outputDTO.InputDTO
	for _, item := range items {
		dto := outputDTO.InputDTO{
			ID:   string(item.ID),
			Data: item.Data,
			Metadata: sharedDTO.Metadata{
				ProcessingId:        item.Metadata.ProcessingId.String(),
				ProcessingTimestamp: item.Metadata.ProcessingTimestamp,
				Source:              item.Metadata.Source,
				Service:             item.Metadata.Service,
				Context:             item.Metadata.Context,
			},
			Status: sharedDTO.Status{
				Code:   item.Status.Code,
				Detail: item.Status.Detail,
			},
		}
		result = append(result, dto)
	}
	return result, nil
}
