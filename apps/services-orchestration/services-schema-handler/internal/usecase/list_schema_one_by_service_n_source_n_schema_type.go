package usecase

import (
	"apps/services-orchestration/services-schema-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-schema-handler/output"
)

type ListOneSchemaByServiceSourceAndSchemaTypeUseCase struct {
	SchemaRepository entity.SchemaInterface
}

func NewListOneSchemaByServiceSourceAndSchemaTypeUseCase(
	repository entity.SchemaInterface,
) *ListOneSchemaByServiceSourceAndSchemaTypeUseCase {
	return &ListOneSchemaByServiceSourceAndSchemaTypeUseCase{
		SchemaRepository: repository,
	}
}

func (la *ListOneSchemaByServiceSourceAndSchemaTypeUseCase) Execute(service string, source string, schemaType string) (outputDTO.SchemaDTO, error) {
	item, err := la.SchemaRepository.FindOneByServiceSourceAndSchemaType(service, source, schemaType)
	if err != nil {
		return outputDTO.SchemaDTO{}, err
	}

	dto := outputDTO.SchemaDTO{
		ID:         string(item.ID),
		SchemaType: item.SchemaType,
		Service:    item.Service,
		Source:     item.Source,
		Context:    item.Context,
		JsonSchema: item.JsonSchema,
		SchemaID:   string(item.SchemaID),
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  item.UpdatedAt,
	}

	return dto, nil
}
