package usecase

import (
	"apps/services-orchestration/services-schema-handler/internal/entity"
	"fmt"
	inputDTO "libs/golang/services/dtos/services-schema-handler/input"
	outputDTO "libs/golang/services/dtos/services-schema-handler/output"
	"libs/golang/shared/go-events/events"
)

type CreateSchemaUseCase struct {
	SchemaRepository entity.SchemaInterface
	SchemaCreated    events.EventInterface
	EventDispatcher  events.EventDispatcherInterface
}

func NewCreateSchemaUseCase(
	repository entity.SchemaInterface,
	SchemaCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *CreateSchemaUseCase {
	return &CreateSchemaUseCase{
		SchemaRepository: repository,
		SchemaCreated:    SchemaCreated,
		EventDispatcher:  EventDispatcher,
	}
}

func (ccu *CreateSchemaUseCase) Execute(schema inputDTO.SchemaDTO) (outputDTO.SchemaDTO, error) {
	schemaEntity, err := entity.NewSchema(
		schema.SchemaType,
		schema.Service,
		schema.Source,
		schema.Context,
		schema.JsonSchema,
	)
	if err != nil {
		return outputDTO.SchemaDTO{}, err
	}

	err = ccu.SchemaRepository.Save(schemaEntity)
	if err != nil {
		return outputDTO.SchemaDTO{}, err
	}

	dto := outputDTO.SchemaDTO{
		ID:         string(schemaEntity.ID),
		SchemaType: schemaEntity.SchemaType,
		Service:    schemaEntity.Service,
		Source:     schemaEntity.Source,
		Context:    schemaEntity.Context,
		JsonSchema: schemaEntity.JsonSchema,
		SchemaID:   schemaEntity.SchemaID,
		CreatedAt:  schemaEntity.CreatedAt,
		UpdatedAt:  schemaEntity.UpdatedAt,
	}

	ccu.SchemaCreated.SetPayload(dto)
	ccu.EventDispatcher.Dispatch(ccu.SchemaCreated, "services", fmt.Sprintf("schema.%s.%s.%s", dto.Context, dto.Service, dto.Source))

	return dto, nil
}
