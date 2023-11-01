package usecase

import (
	"apps/services-orchestration/services-staging-handler/internal/entity"
	"fmt"
	inputDTO "libs/golang/services/dtos/services-staging-handler/input"
	outputDTO "libs/golang/services/dtos/services-staging-handler/output"
	"libs/golang/shared/go-events/events"
)

type CreateProcessingJobDependenciesUseCase struct {
	ProcessingJobDependenciesRepository entity.ProcessingJobDependenciesInterface
	ProcessingJobDependenciesCreated    events.EventInterface
	EventDispatcher                     events.EventDispatcherInterface
}

func NewCreateProcessingJobDependenciesUseCase(
	ProcessingJobDependenciesRepository entity.ProcessingJobDependenciesInterface,
	ProcessingJobDependenciesCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *CreateProcessingJobDependenciesUseCase {
	return &CreateProcessingJobDependenciesUseCase{
		ProcessingJobDependenciesRepository: ProcessingJobDependenciesRepository,
		ProcessingJobDependenciesCreated:    ProcessingJobDependenciesCreated,
		EventDispatcher:                     EventDispatcher,
	}
}

func (c *CreateProcessingJobDependenciesUseCase) Execute(input inputDTO.ProcessingJobDependenciesDTO) (outputDTO.ProcessingJobDependenciesDTO, error) {
	entityJobDependencies := make([]entity.JobDependenciesWithProcessingData, len(input.JobDependencies))
	for i, dep := range input.JobDependencies {
		entityJobDependencies[i] = entity.JobDependenciesWithProcessingData{
			Service: dep.Service,
			Source:  dep.Source,
		}
	}

	processingJobDependenciesEntity, err := entity.NewProcessingJobDependencies(
		input.Service,
		input.Source,
		entityJobDependencies,
	)
	if err != nil {
		return outputDTO.ProcessingJobDependenciesDTO{}, err
	}

	err = c.ProcessingJobDependenciesRepository.Save(processingJobDependenciesEntity)
	if err != nil {
		return outputDTO.ProcessingJobDependenciesDTO{}, err
	}

	dto := outputDTO.ProcessingJobDependenciesDTO{
		ID:              string(processingJobDependenciesEntity.ID),
		Service:         processingJobDependenciesEntity.Service,
		Source:          processingJobDependenciesEntity.Source,
		Context:         processingJobDependenciesEntity.Context,
		JobDependencies: ConvertEntityToUsecaseJobDependenciesWithProcessingData(processingJobDependenciesEntity.JobDependencies),
	}
	c.ProcessingJobDependenciesCreated.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.ProcessingJobDependenciesCreated, "services", fmt.Sprintf("staging-dep.%s.%s.%s", dto.Context, dto.Service, dto.Source))

	return dto, nil
}
