package usecase

import (
	apiClient "libs/golang/services/api-clients/services-input-handler/client"
	inputHandlerInputDTO "libs/golang/services/dtos/services-input-handler/input"
	inputHandlerOutputDTO "libs/golang/services/dtos/services-input-handler/output"
)

type CreateInputUseCase struct {
	InputHandlerAPIClient apiClient.Client
}

func NewCreateInputUseCase() *CreateInputUseCase {
	return &CreateInputUseCase{
		InputHandlerAPIClient: *apiClient.NewClient(),
	}
}

func (ciu *CreateInputUseCase) Execute(service string, source string, contextEnv string, input inputHandlerInputDTO.InputDTO) (inputHandlerOutputDTO.InputDTO, error) {
	inputCreated, err := ciu.InputHandlerAPIClient.CreateInput(service, source, contextEnv, input)
	if err != nil {
		return inputHandlerOutputDTO.InputDTO{}, err
	}
	return inputCreated, nil
}
