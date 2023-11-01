package usecase

import (
	"apps/services-orchestration/services-output-handler/internal/entity"
	inputDTO "libs/golang/services/dtos/services-output-handler/input"
	outputDTO "libs/golang/services/dtos/services-output-handler/output"
	sharedDTO "libs/golang/services/dtos/services-output-handler/shared"
)

type CreateServiceOutputUseCase struct {
	ServiceOutputRepository entity.ServiceOutputInterface
}

func NewCreateServiceOutputUseCase(repository entity.ServiceOutputInterface) *CreateServiceOutputUseCase {
	return &CreateServiceOutputUseCase{
		ServiceOutputRepository: repository,
	}
}

func (csouc *CreateServiceOutputUseCase) Execute(serviceOutput inputDTO.ServiceOutputDTO, service string) (outputDTO.ServiceOutputDTO, error) {
	serviceOutputEntity, err := entity.NewServiceOutput(
		serviceOutput.Data,
		serviceOutput.Metadata.InputId,
		serviceOutput.Metadata.Input.Data,
		serviceOutput.Metadata.Input.ProcessingId,
		serviceOutput.Metadata.Input.ProcessingTimestamp,
		serviceOutput.Metadata.Service,
		serviceOutput.Metadata.Source,
		serviceOutput.Context,
	)
	if err != nil {
		return outputDTO.ServiceOutputDTO{}, err
	}

	err = csouc.ServiceOutputRepository.Save(serviceOutputEntity, service)
	if err != nil {
		return outputDTO.ServiceOutputDTO{}, err
	}

	dto := outputDTO.ServiceOutputDTO{
		ID:   string(serviceOutputEntity.ID),
		Data: serviceOutputEntity.Data,
		Metadata: sharedDTO.Metadata{
			InputId: serviceOutputEntity.Metadata.InputID,
			Input: sharedDTO.MetadataInput{
				ID:                  serviceOutputEntity.Metadata.Input.ID,
				Data:                serviceOutputEntity.Metadata.Input.Data,
				ProcessingId:        serviceOutputEntity.Metadata.Input.ProcessingId,
				ProcessingTimestamp: serviceOutputEntity.Metadata.Input.ProcessingTimestamp,
			},
			Service: serviceOutputEntity.Metadata.Service,
			Source:  serviceOutputEntity.Metadata.Source,
		},
		Service:   serviceOutputEntity.Service,
		Source:    serviceOutputEntity.Source,
		Context:   serviceOutputEntity.Context,
		CreatedAt: serviceOutputEntity.CreatedAt,
		UpdatedAt: serviceOutputEntity.UpdatedAt,
	}

	return dto, nil
}
