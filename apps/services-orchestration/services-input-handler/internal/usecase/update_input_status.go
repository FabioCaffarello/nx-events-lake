package usecase

import (
	"apps/services-orchestration/services-input-handler/internal/entity"
	"fmt"
	outputDTO "libs/golang/services/dtos/services-input-handler/output"
	sharedDTO "libs/golang/services/dtos/services-input-handler/shared"
	"libs/golang/shared/go-events/events"
	"time"
)

type UpdateStatusInputUseCase struct {
	InputRepository    entity.InputInterface
	InputStatusUpdated events.EventInterface
	EventDispatcher    events.EventDispatcherInterface
}

func NewUpdateStatusInputUseCase(
	repository entity.InputInterface,
	InputStatusUpdated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *UpdateStatusInputUseCase {
	return &UpdateStatusInputUseCase{
		InputRepository:    repository,
		InputStatusUpdated: InputStatusUpdated,
		EventDispatcher:    EventDispatcher,
	}
}

func (uiu *UpdateStatusInputUseCase) Execute(service string, source string, contextEnv string, id string, status sharedDTO.Status) (outputDTO.InputDTO, error) {
	input, err := uiu.InputRepository.FindOneByIdAndService(id, service)
	if err != nil {
		return outputDTO.InputDTO{}, err
	}

	input.Status = entity.Status{
		Code:   status.Code,
		Detail: status.Detail,
	}
	input.Metadata.ProcessingTimestamp = time.Now().Format(time.RFC3339)

	err = uiu.InputRepository.Save(input, service)
	if err != nil {
		return outputDTO.InputDTO{}, err
	}

	dto := outputDTO.InputDTO{
		ID:   string(input.ID),
		Data: input.Data,
		Metadata: sharedDTO.Metadata{
			ProcessingId:        input.Metadata.ProcessingId.String(),
			ProcessingTimestamp: input.Metadata.ProcessingTimestamp,
			Source:              input.Metadata.Source,
			Service:             input.Metadata.Service,
			Context:             contextEnv,
		},
		Status: sharedDTO.Status{
			Code:   input.Status.Code,
			Detail: input.Status.Detail,
		},
	}
	uiu.InputStatusUpdated.SetPayload(dto)
	uiu.EventDispatcher.Dispatch(uiu.InputStatusUpdated, "services", fmt.Sprintf("%s.%s.status-updated.%s", dto.Metadata.Context, dto.Metadata.Service, dto.Metadata.Source))

	return dto, nil
}
