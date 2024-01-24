package usecase

import (
	apiClient "libs/golang/services/api-clients/services-staging-handler/client"
	stagingHandlerOutputDTO "libs/golang/services/dtos/services-staging-handler/output"
)

type UpdateProcessingGraphTaskOutputIdUseCase struct {
    StagingHandlerAPIClient apiClient.Client
}

func NewUpdateProcessingGraphTaskOutputIdUseCase() *UpdateProcessingGraphTaskOutputIdUseCase {
    return &UpdateProcessingGraphTaskOutputIdUseCase{
        StagingHandlerAPIClient: *apiClient.NewClient(),
    }
}

func (la *UpdateProcessingGraphTaskOutputIdUseCase) Execute(source string, processingId string, outputId string) (stagingHandlerOutputDTO.ProcessingGraphDTO, error) {
    item, err := la.StagingHandlerAPIClient.UpdateProcessingGraphTaskOutput(source, processingId, outputId)
    if err != nil {
        return stagingHandlerOutputDTO.ProcessingGraphDTO{}, err
    }
    return item, nil
}
