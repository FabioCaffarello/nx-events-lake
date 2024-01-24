package usecase

import (
	apiClient "libs/golang/services/api-clients/services-staging-handler/client"
	stagingHandlerOutputDTO "libs/golang/services/dtos/services-staging-handler/output"
)

type UpdateProcessingGraphTaskStatusUseCase struct {
    StagingHandlerAPIClient apiClient.Client
}

func NewUpdateProcessingGraphTaskStatusUseCase() *UpdateProcessingGraphTaskStatusUseCase {
    return &UpdateProcessingGraphTaskStatusUseCase{
        StagingHandlerAPIClient: *apiClient.NewClient(),
    }
}

func (la *UpdateProcessingGraphTaskStatusUseCase) Execute(source string, processingId string, statusCode int, processingTimestamp string) (stagingHandlerOutputDTO.ProcessingGraphDTO, error) {
    item, err := la.StagingHandlerAPIClient.UpdateProcessingGraphTaskStatus(source, processingId, statusCode, processingTimestamp)
    if err != nil {
        return stagingHandlerOutputDTO.ProcessingGraphDTO{}, err
    }
    return item, nil
}

