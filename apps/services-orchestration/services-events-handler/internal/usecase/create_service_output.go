package usecase

import (
	apiClient "libs/golang/services/api-clients/services-output-handler/client"
	outputHandlerInputDTO "libs/golang/services/dtos/services-output-handler/input"
	outputHandlerOutputDTO "libs/golang/services/dtos/services-output-handler/output"
)

type CreateOutputUseCase struct {
	OutputsAPIClient apiClient.Client
}

func NewCreateOutputUseCase() *CreateOutputUseCase {
	return &CreateOutputUseCase{
		OutputsAPIClient: *apiClient.NewClient(),
	}
}

func (cou *CreateOutputUseCase) Execute(service string, serviceOutput outputHandlerInputDTO.ServiceOutputDTO) (outputHandlerOutputDTO.ServiceOutputDTO, error) {
	serviceOutputCreated, err := cou.OutputsAPIClient.CreateOutput(service, serviceOutput)
	if err != nil {
		return outputHandlerOutputDTO.ServiceOutputDTO{}, err
	}
	return serviceOutputCreated, nil
}
