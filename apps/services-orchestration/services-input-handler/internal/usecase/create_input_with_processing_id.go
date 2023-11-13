package usecase

import (
	"apps/services-orchestration/services-input-handler/internal/entity"
	"fmt"
	inputDTO "libs/golang/services/dtos/services-input-handler/input"
	outputDTO "libs/golang/services/dtos/services-input-handler/output"
	sharedDTO "libs/golang/services/dtos/services-input-handler/shared"
	"libs/golang/shared/go-events/events"
	"libs/golang/shared/go-id/uuid"
)

type CreateInputWithProcessingIDUseCase struct {
	InputRepository entity.InputInterface
	InputCreated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewCreateInputWithProcessingIDUseCase(
	repository entity.InputInterface,
	InputCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *CreateInputWithProcessingIDUseCase {
	return &CreateInputWithProcessingIDUseCase{
		InputRepository: repository,
		InputCreated:    InputCreated,
		EventDispatcher: EventDispatcher,
	}
}

func (ciu *CreateInputWithProcessingIDUseCase) Execute(input inputDTO.InputDTO, service string, source string, contextEnv string, processingId string) (outputDTO.InputDTO, error) {
	inputEntity, err := entity.NewInput(input.Data, source, service, contextEnv)
	if err != nil {
		return outputDTO.InputDTO{}, err
	}

    uuidProcessingId, err := uuid.ParseID(processingId)
    if err != nil {
        return outputDTO.InputDTO{}, err
    }

    inputEntity.Metadata.ProcessingId = uuidProcessingId

	err = ciu.InputRepository.Save(inputEntity, service)
	if err != nil {
		return outputDTO.InputDTO{}, err
	}

	dto := outputDTO.InputDTO{
		ID:   string(inputEntity.ID),
		Data: inputEntity.Data,
		Metadata: sharedDTO.Metadata{
			ProcessingId:        inputEntity.Metadata.ProcessingId.String(),
			ProcessingTimestamp: inputEntity.Metadata.ProcessingTimestamp,
			Source:              inputEntity.Metadata.Source,
			Service:             inputEntity.Metadata.Service,
			Context:             inputEntity.Metadata.Context,
		},
		Status: sharedDTO.Status{
			Code:   inputEntity.Status.Code,
			Detail: inputEntity.Status.Detail,
		},
	}
	ciu.InputCreated.SetPayload(dto)
	ciu.EventDispatcher.Dispatch(ciu.InputCreated, "services", fmt.Sprintf("%s.%s.inputs.%s", dto.Metadata.Context, dto.Metadata.Service, dto.Metadata.Source))

	return dto, nil
}
