package usecase

import (
	"apps/services-orchestration/services-input-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-input-handler/output"
	sharedDTO "libs/golang/services/dtos/services-input-handler/shared"
)

type ListOneByIdAndServiceUseCase struct {
	InputRepository entity.InputInterface
}

func NewListOneByIdAndServiceUseCase(
	repository entity.InputInterface,
) *ListOneByIdAndServiceUseCase {
	return &ListOneByIdAndServiceUseCase{
		InputRepository: repository,
	}
}

func (lo *ListOneByIdAndServiceUseCase) Execute(service string, id string) (outputDTO.InputDTO, error) {
	item, err := lo.InputRepository.FindOneByIdAndService(id, service)
	if err != nil {
		return outputDTO.InputDTO{}, err
	}
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
	return dto, nil
}
