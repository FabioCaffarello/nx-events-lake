package usecase

import (
	apiConfigClient "libs/golang/services/api-clients/services-config-handler/client"
	apiSchemaClient "libs/golang/services/api-clients/services-schema-handler/client"
	apiClient "libs/golang/services/api-clients/services-staging-handler/client"
	stagingHandlerInputDTO "libs/golang/services/dtos/services-staging-handler/input"
	stagingHandlerOutputDTO "libs/golang/services/dtos/services-staging-handler/output"
	configId "libs/golang/shared/go-id/config"
)

type CreateProcessingGraphTaskUseCase struct {
    ConfigHandlerAPIClient  apiConfigClient.Client
	SchemaHandlerAPIClient  apiSchemaClient.Client
	StagingHandlerAPIClient apiClient.Client
	SchemaInputType         string
	SchemaOutputType        string
}

func NewCreateProcessingGraphTaskUseCase() *CreateProcessingGraphTaskUseCase {
	return &CreateProcessingGraphTaskUseCase{
        ConfigHandlerAPIClient:  *apiConfigClient.NewClient(),
		SchemaHandlerAPIClient:  *apiSchemaClient.NewClient(),
		StagingHandlerAPIClient: *apiClient.NewClient(),
		SchemaInputType:         "service-input",
		SchemaOutputType:        "service-output",
	}
}

func (la *CreateProcessingGraphTaskUseCase) Execute(
	source string,
	service string,
	processingId string,
	contextEnv string,
	parentProcessingId string,
	inputId string,
	outputId string,
	processingTimestamp string,
	startProcessingId string,
) (stagingHandlerOutputDTO.ProcessingGraphDTO, error) {
    statusCode := 0
    _, err := la.StagingHandlerAPIClient.ListOneProcessingGraphByTaskSourceAndParentProcessingId(source, processingId)
	if err == nil {
		return stagingHandlerOutputDTO.ProcessingGraphDTO{}, nil
	}

	schemaInput, err := la.SchemaHandlerAPIClient.ListOneSchemaByServiceAndSourceAndContextAndSchemaType(service, source, contextEnv, la.SchemaInputType)
	if err != nil {
		return stagingHandlerOutputDTO.ProcessingGraphDTO{}, err
	}
	schemaOutput, err := la.SchemaHandlerAPIClient.ListOneSchemaByServiceAndSourceAndContextAndSchemaType(service, source, contextEnv, la.SchemaOutputType)
	if err != nil {
		return stagingHandlerOutputDTO.ProcessingGraphDTO{}, err
	}

    id := getConfigId(service, source)
    config, err := la.ConfigHandlerAPIClient.ListOneConfigById(id)
	if err != nil {
		return stagingHandlerOutputDTO.ProcessingGraphDTO{}, err
	}


	graphTask := setProcessingGraphTask(
		source,
		service,
		processingId,
		parentProcessingId,
		config.ConfigID,
		schemaInput.SchemaID,
		schemaOutput.SchemaID,
		inputId,
		outputId,
		statusCode,
		processingTimestamp,
	)

	ProcessingGraphTask, err := la.StagingHandlerAPIClient.CreateProcessingGraphTask(source, startProcessingId, graphTask)
	if err != nil {
		return stagingHandlerOutputDTO.ProcessingGraphDTO{}, err
	}
	return ProcessingGraphTask, nil
}

func setProcessingGraphTask(
	source string,
	service string,
	processingId string,
	parentProcessingId string,
	configVersionId string,
	inputSchemaVersionId string,
	outputSchemaVersionId string,
	inputId string,
	outputId string,
	statusCode int,
	processingTimestamp string,
) stagingHandlerInputDTO.Task {
	return stagingHandlerInputDTO.Task{
		Source:                source,
		Service:               service,
		ProcessingId:          processingId,
		ParentProcessingId:    parentProcessingId,
		ConfigVersionId:       configVersionId,
		InputSchemaVersionId:  inputSchemaVersionId,
		OutputSchemaVersionId: outputSchemaVersionId,
		InputId:               inputId,
		OutputId:              outputId,
		StatusCode:            statusCode,
		ProcessingTimestamp:   processingTimestamp,
	}
}

func getConfigId(service string, source string) string {
	return configId.NewID(service, source)
}
