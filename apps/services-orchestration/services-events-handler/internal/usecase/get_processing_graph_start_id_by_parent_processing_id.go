package usecase

import (
	apiClient "libs/golang/services/api-clients/services-staging-handler/client"
)

type GetProcessingGraphStartIdByParentProcessingIdUseCase struct {
	StagingHandlerAPIClient apiClient.Client
}

func NewGetProcessingGraphStartIdByParentProcessingIdUseCase() *GetProcessingGraphStartIdByParentProcessingIdUseCase {
	return &GetProcessingGraphStartIdByParentProcessingIdUseCase{
		StagingHandlerAPIClient: *apiClient.NewClient(),
	}
}

func (la *GetProcessingGraphStartIdByParentProcessingIdUseCase) Execute(source string, parentProcessingId string) (string, error) {
	item, err := la.StagingHandlerAPIClient.ListOneProcessingGraphByTaskSourceAndParentProcessingId(source, parentProcessingId)
	if err != nil {
		return "", err
	}
	return item.StartProcessingId, nil
}
