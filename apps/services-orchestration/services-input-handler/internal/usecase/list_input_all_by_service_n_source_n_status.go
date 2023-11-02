package usecase

import (
	"apps/services-orchestration/services-input-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-input-handler/output"
	sharedDTO "libs/golang/services/dtos/services-input-handler/shared"
)

type ListAllByServiceAndSourceAndStatusUseCase struct {
	InputRepository entity.InputInterface
}

func NewListAllByServiceAndSourceAndStatusUseCase(
	repository entity.InputInterface,
) *ListAllByServiceAndSourceAndStatusUseCase {
	return &ListAllByServiceAndSourceAndStatusUseCase{
		InputRepository: repository,
	}
}

func (la *ListAllByServiceAndSourceAndStatusUseCase) Execute(service string, source string, status int) ([]outputDTO.InputDTO, error) {
	items, err := la.InputRepository.FindAllByServiceAndSourceAndStatus(service, source, status)
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
