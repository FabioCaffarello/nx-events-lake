package channels

import (
	"encoding/json"
	"errors"
	"fmt"

	"apps/services-orchestration/services-events-handler/internal/usecase"
	inputHandlerOutputDTO "libs/golang/services/dtos/services-input-handler/output"
	inputHandlerSharedDTO "libs/golang/services/dtos/services-input-handler/shared"
	inputHandlersharedDTO "libs/golang/services/dtos/services-input-handler/shared"
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
	statusInputDTO := setStatusFlagToProcessing()

	updateInputUseCase := usecase.NewUpdateStatusInputUseCase()
	findAllDependentJobUseCase := usecase.NewListAllConfigsByDependentJobUseCase()
	createProcessingJobDependenciesUseCase := usecase.NewCreateProcessingJobDependenciesUseCase()

	_, err = updateInputUseCase.Execute(statusInputDTO, contextEnv, service, source, id)
	if err != nil {
		return err
	}

	dependentJobs, err := findAllDependentJobUseCase.Execute(service, source)
	if err != nil {
		fmt.Println(err)
	}

	for _, dependentJob := range dependentJobs {
		jobDeps := make([]statingHandlerSharedDTO.ProcessingJobDependencies, len(dependentJob.DependsOn))
		for i, dep := range dependentJob.DependsOn {
			jobDeps[i] = statingHandlerSharedDTO.ProcessingJobDependencies{
				Service: dep.Service,
				Source:  dep.Source,
			}
		}

		processingJobDependency := statingHandlerInputDTO.ProcessingJobDependenciesDTO{
			Service:         dependentJob.Service,
			Source:          dependentJob.Source,
			JobDependencies: jobDeps,
		}

		_, err = createProcessingJobDependenciesUseCase.Execute(processingJobDependency)
		if err != nil {
			fmt.Println(err)
		}

	}

	return nil
}

func setStatusFlagToProcessing() inputHandlerSharedDTO.Status {
	return inputHandlersharedDTO.Status{
		Code:   1,
		Detail: "Processing",
	}
}
