package usecase

import (
	"apps/services-orchestration/services-output-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-output-handler/output"
	sharedDTO "libs/golang/services/dtos/services-output-handler/shared"
)

type ListOneServiceOutputByServiceAndIdUseCase struct {
	ServiceOutputRepository entity.ServiceOutputInterface
}

func NewListOneServiceOutputByServiceAndIdUseCase(
	repository entity.ServiceOutputInterface,
) *ListOneServiceOutputByServiceAndIdUseCase {
	return &ListOneServiceOutputByServiceAndIdUseCase{
		ServiceOutputRepository: repository,
	}
}

func (la *ListOneServiceOutputByServiceAndIdUseCase) Execute(service string, id string) (outputDTO.ServiceOutputDTO, error) {
	item, err := la.ServiceOutputRepository.FindOneByIdAndService(id, service)
	if err != nil {
		return outputDTO.ServiceOutputDTO{}, err
	}
	dto := outputDTO.ServiceOutputDTO{
		ID:   string(item.ID),
		Data: item.Data,
		Metadata: sharedDTO.Metadata{
			InputId: item.Metadata.InputID,
			Input: sharedDTO.MetadataInput{
				ID:                  item.Metadata.Input.ID,
				Data:                item.Metadata.Input.Data,
				ProcessingId:        item.Metadata.Input.ProcessingId,
				ProcessingTimestamp: item.Metadata.Input.ProcessingTimestamp,
			},
			Service: item.Metadata.Service,
			Source:  item.Metadata.Source,
		},
		Service:   item.Service,
		Source:    item.Source,
		Context:   item.Context,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
	return dto, nil
}
