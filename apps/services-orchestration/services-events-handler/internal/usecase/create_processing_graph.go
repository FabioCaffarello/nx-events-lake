package usecase


import (
	apiClient "libs/golang/services/api-clients/services-staging-handler/client"
	stagingHandlerInputDTO "libs/golang/services/dtos/services-staging-handler/input"
	stagingHandlerOutputDTO "libs/golang/services/dtos/services-staging-handler/output"
)

type CreateProcessingGraphUseCase struct {
    StagingHandlerAPIClient apiClient.Client
}

func NewCreateProcessingGraphUseCase() *CreateProcessingGraphUseCase {
    return &CreateProcessingGraphUseCase{
        StagingHandlerAPIClient: *apiClient.NewClient(),
    }
}

func (la *CreateProcessingGraphUseCase) Execute(graph stagingHandlerInputDTO.ProcessingGraphDTO, processingId string) (stagingHandlerOutputDTO.ProcessingGraphDTO, string, error) {
    _, err := la.StagingHandlerAPIClient.ListOneProcessingGraphBySourceAndStartProcessingId(graph.Source, processingId)
    if err == nil {
        return stagingHandlerOutputDTO.ProcessingGraphDTO{}, processingId ,nil
    }
    existingProcessGraphStart, err := la.StagingHandlerAPIClient.ListOneProcessingGraphByTaskSourceAndParentProcessingId(graph.Source, processingId)
    if err == nil {
        return stagingHandlerOutputDTO.ProcessingGraphDTO{}, existingProcessGraphStart.StartProcessingId , nil
    }
    createProcessingGraph,  err := la.StagingHandlerAPIClient.CreateProcessingGraph(graph)
    if err != nil {
        return stagingHandlerOutputDTO.ProcessingGraphDTO{}, "", err
    }
    return createProcessingGraph, createProcessingGraph.StartProcessingId, nil
}
