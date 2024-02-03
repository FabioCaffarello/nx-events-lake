package channels

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"apps/services-orchestration/services-events-handler/internal/usecase"
	eventsHandlerInputDTO "libs/golang/services/dtos/services-events-handler/input"
	inputHandlerSharedDTO "libs/golang/services/dtos/services-input-handler/shared"
	outputHandlerInputDTO "libs/golang/services/dtos/services-output-handler/input"
	outputHandlerISharedDTO "libs/golang/services/dtos/services-output-handler/shared"
	stagingHandlerSharedDTO "libs/golang/services/dtos/services-staging-handler/shared"
	generateInputs "libs/golang/services/modules/generate-inputs/core"
	configId "libs/golang/shared/go-id/config"
	"libs/golang/shared/go-id/md5"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	ErrorInvalidServiceFeedbackDTO = errors.New("invalid service feedback message")
	ErrorInvalidStatus             = errors.New("invalid status code")
)

type ServiceFeedbackListener struct {
	GenerateInputs *generateInputs.DomainFactory
}

func NewServiceFeedbackListener() *ServiceFeedbackListener {
	return &ServiceFeedbackListener{
		GenerateInputs: generateInputs.NewDomainFactory(),
	}
}

func (l *ServiceFeedbackListener) Handle(msg amqp.Delivery) error {
	var serviceFeedbackDTO eventsHandlerInputDTO.ServiceFeedbackDTO
	err := json.Unmarshal(msg.Body, &serviceFeedbackDTO)
	if err != nil {
		return ErrorInvalidServiceFeedbackDTO
	}

	statusDTO := getStatusInputDTO(serviceFeedbackDTO)
	service := serviceFeedbackDTO.Metadata.Service
	source := serviceFeedbackDTO.Metadata.Source
	contextEnv := serviceFeedbackDTO.Metadata.Context
	processingId := serviceFeedbackDTO.Metadata.Input.ProcessingId
	id := serviceFeedbackDTO.Metadata.Input.ID

	updateInputUseCase := usecase.NewUpdateStatusInputUseCase()

	_, err = updateInputUseCase.Execute(statusDTO, contextEnv, service, source, id)
	if err != nil {
		log.Println(err)
	}

	updateProcessingGraphTaskStatus := usecase.NewUpdateProcessingGraphTaskStatusUseCase()
	_, err = updateProcessingGraphTaskStatus.Execute(source, processingId, statusDTO.Code, serviceFeedbackDTO.Metadata.ProcessingTimestamp)
	if err != nil {
		log.Println(err)
	}

	outputData := getServiceOutputDTO(serviceFeedbackDTO)
	createOutputUseCase := usecase.NewCreateOutputUseCase()
	output, err := createOutputUseCase.Execute(service, outputData)
	if err != nil {
		log.Println(err)
	}

	updateProcessingGraphTaskOutput := usecase.NewUpdateProcessingGraphTaskOutputIdUseCase()
	_, err = updateProcessingGraphTaskOutput.Execute(source, processingId, output.ID)
	if err != nil {
		log.Println(err)
	}

	findAllDependentJobUseCase := usecase.NewListAllConfigsByDependentJobUseCase()
	createInputUseCase := usecase.NewCreateInputUseCase()

	updateProcessingJobDependenciesUseCase := usecase.NewUpdateProcessingJobDependenciesUseCase()
	checkAllJobDependenciesStatus200UseCase := usecase.NewCheckAllJobDependenciesStatus200UseCase()
	removeProcessingJobDependenciesUseCase := usecase.NewRemoveProcessingJobDependenciesUseCase()
	findOneJobConfigById := usecase.NewListOneConfigByIdUseCase()

	dependentJobs, err := findAllDependentJobUseCase.Execute(service, source)
	if err != nil {
		log.Println(err)
	}

	jobDep := getProcessingJobDependencies(serviceFeedbackDTO)

	createProcessingGraphTaskUseCase := usecase.NewCreateProcessingGraphTaskUseCase()
	processingGraphStartProcessingIdUseCase := usecase.NewGetProcessingGraphStartIdByParentProcessingIdUseCase()
	processingGraphStartProcessingId, err := processingGraphStartProcessingIdUseCase.Execute(source, processingId)
	if err != nil {
		log.Println(err)
	}

	if len(dependentJobs) > 0 {
		for _, dependentJob := range dependentJobs {
			processingJobParentId := string(md5.NewMd5Hash(fmt.Sprintf("%s-%s-%s-%s", dependentJob.Context, dependentJob.Service, dependentJob.Source, processingId)))

			updateProcessingJobDependenciesUseCase.Execute(processingJobParentId, jobDep)
			shouldRun, err := checkAllJobDependenciesStatus200UseCase.Execute(processingJobParentId)
			log.Println("\nshouldRun: ", shouldRun)
			if err != nil {
				log.Println(err)
			}
            // TODO: Consider batch control
            // // multiple files from the same source
			if shouldRun {
				depJobConfigId := getConfigId(dependentJob.Service, dependentJob.Source)
				depJobConfig, err := findOneJobConfigById.Execute(depJobConfigId)
				if err != nil {
					log.Println(err)
				}
				inputDTOs, err := l.GenerateInputs.GenerateInputs(depJobConfig.InputMethod, outputData.Data)
				if err != nil {
					log.Println(err)
				}

				for _, inputDTO := range inputDTOs {
					inputDep, err := createInputUseCase.Execute(dependentJob.Service, dependentJob.Source, dependentJob.Context, inputDTO)
					if err != nil {
						log.Println(err)
					}
					_, err = createProcessingGraphTaskUseCase.Execute(
						dependentJob.Source,
						dependentJob.Service,
						inputDep.Metadata.ProcessingId,
						contextEnv,
						processingId,
						inputDep.ID,
						output.ID,
						inputDep.Metadata.ProcessingTimestamp,
						processingGraphStartProcessingId,
					)
					if err != nil {
						log.Println(err)
					}
				}
				removeProcessingJobDependenciesUseCase.Execute(processingJobParentId)
			}
		}
	}
	return nil
}

func InterfaceArrayToStringArray(interfaceArray []interface{}) []string {
	stringArray := make([]string, len(interfaceArray))
	for i, v := range interfaceArray {
		stringArray[i] = v.(string)
	}
	return stringArray
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

func getConfigId(service string, source string) string {
	return configId.NewID(service, source)
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
