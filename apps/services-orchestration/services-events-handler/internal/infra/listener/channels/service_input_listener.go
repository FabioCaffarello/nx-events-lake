package channels

import (
	"encoding/json"
	"errors"
	"log"

	"apps/services-orchestration/services-events-handler/internal/usecase"
	inputHandlerOutputDTO "libs/golang/services/dtos/services-input-handler/output"
	inputHandlerSharedDTO "libs/golang/services/dtos/services-input-handler/shared"
	statingHandlerInputDTO "libs/golang/services/dtos/services-staging-handler/input"
	statingHandlerSharedDTO "libs/golang/services/dtos/services-staging-handler/shared"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	ErrorInvalidServiceInputDTO = errors.New("invalid service input message")
)

type ServiceInputListener struct {
}

func NewServiceInputListener() *ServiceInputListener {
	return &ServiceInputListener{}
}

func (l *ServiceInputListener) Handle(msg amqp.Delivery) error {
	var serviceInputDTO inputHandlerOutputDTO.InputDTO
	err := json.Unmarshal(msg.Body, &serviceInputDTO)
	if err != nil {
		return ErrorInvalidServiceInputDTO
	}
	service := serviceInputDTO.Metadata.Service
	source := serviceInputDTO.Metadata.Source
	contextEnv := serviceInputDTO.Metadata.Context
	id := serviceInputDTO.ID
	processingTimestamp := serviceInputDTO.Metadata.ProcessingTimestamp
	statusInputDTO := setStatusFlagToProcessing()

	updateInputUseCase := usecase.NewUpdateStatusInputUseCase()
	findAllDependentJobUseCase := usecase.NewListAllConfigsByDependentJobUseCase()
	createProcessingJobDependenciesUseCase := usecase.NewCreateProcessingJobDependenciesUseCase()

	log.Println("input Id: ", id)
	log.Println("inputStatus: ", statusInputDTO)

	inputUpdated, err := updateInputUseCase.Execute(statusInputDTO, contextEnv, service, source, id)
	if err != nil {
		return err
	}

	processingId := inputUpdated.Metadata.ProcessingId

	processingGraphInput := setProcessingGraph(source, contextEnv, processingId)
	createProcessingGraphUseCase := usecase.NewCreateProcessingGraphUseCase()
	_, _, err = createProcessingGraphUseCase.Execute(processingGraphInput, processingId)
	if err != nil {
		log.Println(err)
	}
	createProcessingGraphTaskUseCase := usecase.NewCreateProcessingGraphTaskUseCase()
	_, err = createProcessingGraphTaskUseCase.Execute(
		source,
		service,
		processingId,
		contextEnv,
		processingId,
		id,
		"",
		processingTimestamp,
		processingId,
	)
	if err != nil {
		log.Println(err)
	}

	updateProcessingGraphTaskStatus := usecase.NewUpdateProcessingGraphTaskStatusUseCase()
	_, err = updateProcessingGraphTaskStatus.Execute(source, processingId, 1, processingTimestamp)
	if err != nil {
		log.Println(err)
	}

	dependentJobs, err := findAllDependentJobUseCase.Execute(service, source)
	if err != nil {
		log.Println(err)
	}

	for _, dependentJob := range dependentJobs {
		log.Println("dependentJob", dependentJob)
		jobDeps := make([]statingHandlerSharedDTO.ProcessingJobDependencies, len(dependentJob.DependsOn))
		for i, dep := range dependentJob.DependsOn {
			jobDeps[i] = statingHandlerSharedDTO.ProcessingJobDependencies{
				Service: dep.Service,
				Source:  dep.Source,
			}
		}

		processingJobDependency := statingHandlerInputDTO.ProcessingJobDependenciesDTO{
			Service:               dependentJob.Service,
			Source:                dependentJob.Source,
			Context:               contextEnv,
			ParentJobProcessingId: processingId,
			JobDependencies:       jobDeps,
		}

		_, err = createProcessingJobDependenciesUseCase.Execute(processingJobDependency)
		log.Println("processingJobDependency", processingJobDependency)
		if err != nil {
			log.Println(err)
		}

	}

	return nil
}

func setStatusFlagToProcessing() inputHandlerSharedDTO.Status {
	return inputHandlerSharedDTO.Status{
		Code:   1,
		Detail: "Processing",
	}
}

func setProcessingGraph(source string, context string, startProcessingId string) statingHandlerInputDTO.ProcessingGraphDTO {
	return statingHandlerInputDTO.ProcessingGraphDTO{
		Context:           context,
		Source:            source,
		StartProcessingId: startProcessingId,
	}
}
