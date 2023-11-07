package usecase

import (
	apiClient "libs/golang/services/api-clients/services-input-handler/client"
	inputHandlerOutputDTO "libs/golang/services/dtos/services-input-handler/output"
	inputHandlerSharedDTO "libs/golang/services/dtos/services-input-handler/shared"
)

type UpdateStatusInputUseCase struct {
	InputHandlerAPIClient apiClient.Client
}

func NewUpdateStatusInputUseCase() *UpdateStatusInputUseCase {
	return &UpdateStatusInputUseCase{
		InputHandlerAPIClient: *apiClient.NewClient(),
	}
}

func (uiu *UpdateStatusInputUseCase) Execute(status inputHandlerSharedDTO.Status, contextEnv string, service string, source string, id string) (inputHandlerOutputDTO.InputDTO, error) {
	input, err := uiu.InputHandlerAPIClient.UpdateInputStatus(status, contextEnv, service, source, id)
	if err != nil {
		return inputHandlerOutputDTO.InputDTO{}, err
	}
	return input, nil
}
