package channels

import (
	"encoding/json"
	"errors"
	"fmt"

	"apps/services-orchestration/services-events-handler/internal/usecase"
	eventsHandlerInputDTO "libs/golang/services/dtos/services-events-handler/input"
	inputHandlerInputDTO "libs/golang/services/dtos/services-input-handler/input"
	inputHandlerSharedDTO "libs/golang/services/dtos/services-input-handler/shared"
	outputHandlerInputDTO "libs/golang/services/dtos/services-output-handler/input"
	outputHandlerISharedDTO "libs/golang/services/dtos/services-output-handler/shared"
	stagingHandlerSharedDTO "libs/golang/services/dtos/services-staging-handler/shared"
	configID "libs/golang/shared/go-id/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	ErrorInvalidServiceFeedbackDTO = errors.New("invalid service feedback message")
	ErrorInvalidStatus             = errors.New("invalid status code")
)

type ServiceFeedbackListener struct {
}

func NewServiceFeedbackListener() *ServiceFeedbackListener {
	return &ServiceFeedbackListener{}
}

func extractStatusCodeRange(statusCode int) string {
	if statusCode >= 200 && statusCode < 300 {
		return "2XX"
	} else if statusCode >= 400 && statusCode < 500 {
		return "4XX"
	} else if statusCode >= 500 && statusCode < 600 {
		return "5XX"
	}
	return "invalid"
}

func (l *ServiceFeedbackListener) Handle(msg amqp.Delivery) error {
	// fmt.Println(string(msg.Body))
	var serviceFeedbackDTO eventsHandlerInputDTO.ServiceFeedbackDTO
	err := json.Unmarshal(msg.Body, &serviceFeedbackDTO)
	if err != nil {
		return ErrorInvalidServiceFeedbackDTO
	}
	statusCodeRange := extractStatusCodeRange(serviceFeedbackDTO.Status.Code)

	statusDTO := getStatusInputDTO(serviceFeedbackDTO)
	service := serviceFeedbackDTO.Metadata.Service
	source := serviceFeedbackDTO.Metadata.Source
	contextEnv := serviceFeedbackDTO.Metadata.Context
	id := serviceFeedbackDTO.Metadata.Input.ID

	updateInputUseCase := usecase.NewUpdateStatusInputUseCase()

	_, err = updateInputUseCase.Execute(statusDTO, contextEnv, service, source, id)
	if err != nil {
		fmt.Println(err)
	}

	switch statusCodeRange {
	case "2XX":
		l.HandleFeedback200(serviceFeedbackDTO, service, source)
	case "4XX":
		l.HandleFeedback400(serviceFeedbackDTO, service, source)
	case "5XX":
		l.HandleFeedback500(serviceFeedbackDTO, service, source)
	default:
		return ErrorInvalidStatus
	}
	return nil
}

func (l *ServiceFeedbackListener) HandleFeedback200(msg eventsHandlerInputDTO.ServiceFeedbackDTO, service string, source string) {
	outputData := getServiceOutputDTO(msg)
	createOutputUseCase := usecase.NewCreateOutputUseCase()
	_, err := createOutputUseCase.Execute(service, outputData)
	if err != nil {
		fmt.Println(err)
	}

	findAllDependentJobUseCase := usecase.NewListAllConfigsByDependentJobUseCase()
	createInputUseCase := usecase.NewCreateInputUseCase()

	updateProcessingJobDependenciesUseCase := usecase.NewUpdateProcessingJobDependenciesUseCase()
	checkAllJobDependenciesStatus200UseCase := usecase.NewCheckAllJobDependenciesStatus200UseCase()
	removeProcessingJobDependenciesUseCase := usecase.NewRemoveProcessingJobDependenciesUseCase()

	dependentJobs, err := findAllDependentJobUseCase.Execute(service, source)
	if err != nil {
		fmt.Println(err)
	}

	inputDTO := inputHandlerInputDTO.InputDTO{
		Data: map[string]interface{}{
			"documentUri": msg.Data["documentUri"],
			"partition":   msg.Data["partition"],
		},
	}

	jobDep := getProcessingJobDependencies(msg)

	for _, dependentJob := range dependentJobs {
		processingJobDepId := configID.NewID(dependentJob.Service, dependentJob.Source)

		updateProcessingJobDependenciesUseCase.Execute(processingJobDepId, jobDep)
		shouldRun, err := checkAllJobDependenciesStatus200UseCase.Execute(processingJobDepId)
		if err != nil {
			fmt.Println(err)
		}
		if shouldRun {
			_, err := createInputUseCase.Execute(dependentJob.Service, dependentJob.Source, dependentJob.Context, inputDTO)
			if err != nil {
				fmt.Println(err)
			}
			removeProcessingJobDependenciesUseCase.Execute(processingJobDepId)
		}

	}

}

func (l *ServiceFeedbackListener) HandleFeedback400(msg eventsHandlerInputDTO.ServiceFeedbackDTO, service string, source string) {

}

func (l *ServiceFeedbackListener) HandleFeedback500(msg eventsHandlerInputDTO.ServiceFeedbackDTO, service string, source string) {

}

func getServiceOutputDTO(msg eventsHandlerInputDTO.ServiceFeedbackDTO) outputHandlerInputDTO.ServiceOutputDTO {
	return outputHandlerInputDTO.ServiceOutputDTO{
		Data:    msg.Data,
		Context: msg.Metadata.Context,
		Metadata: outputHandlerISharedDTO.Metadata{
			InputId: msg.Metadata.Input.ID,
			Input: outputHandlerISharedDTO.MetadataInput{
				ID:                  msg.Metadata.Input.ID,
				Data:                msg.Metadata.Input.Data,
				ProcessingId:        msg.Metadata.Input.ProcessingId,
				ProcessingTimestamp: msg.Metadata.Input.ProcessingTimestamp,
			},
			Service: msg.Metadata.Service,
			Source:  msg.Metadata.Source,
		},
	}
}

func getStatusInputDTO(msg eventsHandlerInputDTO.ServiceFeedbackDTO) inputHandlerSharedDTO.Status {
	return inputHandlerSharedDTO.Status{
		Code:   msg.Status.Code,
		Detail: msg.Status.Detail,
	}
}

func getProcessingJobDependencies(msg eventsHandlerInputDTO.ServiceFeedbackDTO) stagingHandlerSharedDTO.ProcessingJobDependencies {
	return stagingHandlerSharedDTO.ProcessingJobDependencies{
		Service:             msg.Metadata.Service,
		Source:              msg.Metadata.Source,
		ProcessingId:        msg.Metadata.Input.ProcessingId,
		ProcessingTimestamp: msg.Metadata.ProcessingTimestamp,
		StatusCode:          msg.Status.Code,
	}
}
