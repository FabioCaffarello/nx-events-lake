package usecase

import (
	"apps/services-orchestration/services-schema-handler/internal/entity"
	"fmt"
	inputDTO "libs/golang/services/dtos/services-schema-handler/input"
	outputDTO "libs/golang/services/dtos/services-schema-handler/output"
	"libs/golang/shared/go-events/events"
)

type UpdateSchemaUseCase struct {
	SchemaRepository entity.SchemaInterface
	SchemaUpdated    events.EventInterface
	EventDispatcher  events.EventDispatcherInterface
}

func NewUpdateSchemaUseCase(
	repository entity.SchemaInterface,
	SchemaUpdated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *UpdateSchemaUseCase {
	return &UpdateSchemaUseCase{
		SchemaRepository: repository,
		SchemaUpdated:    SchemaUpdated,
		EventDispatcher:  EventDispatcher,
	}
}

func (ccu *UpdateSchemaUseCase) Execute(schema inputDTO.SchemaDTO) (outputDTO.SchemaDTO, error) {
	item, err := ccu.SchemaRepository.FindOneByServiceAndSourceAndContextAndSchemaType(schema.Service, schema.Source, schema.Context, schema.SchemaType)
	if err != nil {
		return outputDTO.SchemaDTO{}, err
	}

    schemaEntity, err := entity.NewSchema(
        schema.SchemaType,
        schema.Context,
        schema.Service,
        schema.Source,
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
		SchemaID:   string(schemaEntity.SchemaID),
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  schemaEntity.UpdatedAt,
	}

	ccu.SchemaUpdated.SetPayload(dto)
	ccu.EventDispatcher.Dispatch(ccu.SchemaUpdated, "services", fmt.Sprintf("schema.%s.%s.%s", dto.Context, dto.Service, dto.Source))

	return dto, nil
}
