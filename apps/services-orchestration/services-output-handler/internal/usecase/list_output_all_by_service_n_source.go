package usecase

import (
	"apps/services-orchestration/services-output-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-output-handler/output"
	sharedDTO "libs/golang/services/dtos/services-output-handler/shared"
)

type ListAllServiceOutputByServiceAndSourceUseCase struct {
	ServiceOutputRepository entity.ServiceOutputInterface
}

func NewListAllServiceOutputByServiceAndSourceUseCase(
	repository entity.ServiceOutputInterface,
) *ListAllServiceOutputByServiceAndSourceUseCase {
	return &ListAllServiceOutputByServiceAndSourceUseCase{
		ServiceOutputRepository: repository,
	}
}

func (la *ListAllServiceOutputByServiceAndSourceUseCase) Execute(service string, source string) ([]outputDTO.ServiceOutputDTO, error) {
	items, err := la.ServiceOutputRepository.FindAllByServiceAndSource(service, source)
	if err != nil {
		return []outputDTO.ServiceOutputDTO{}, err
	}
	var result []outputDTO.ServiceOutputDTO
	for _, item := range items {
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
		result = append(result, dto)
	}
	return result, nil
}
