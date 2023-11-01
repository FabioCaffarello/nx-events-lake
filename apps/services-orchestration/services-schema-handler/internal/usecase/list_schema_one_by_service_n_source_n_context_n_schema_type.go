package usecase

import (
	"apps/services-orchestration/services-schema-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-schema-handler/output"
)

type ListOneSchemaByServiceSourceAndContextAndSchemaTypeUseCase struct {
	SchemaRepository entity.SchemaInterface
}

func NewListOneSchemaByServiceSourceAndContextAndSchemaTypeUseCase(
	repository entity.SchemaInterface,
) *ListOneSchemaByServiceSourceAndContextAndSchemaTypeUseCase {
	return &ListOneSchemaByServiceSourceAndContextAndSchemaTypeUseCase{
		SchemaRepository: repository,
	}
}

func (la *ListOneSchemaByServiceSourceAndContextAndSchemaTypeUseCase) Execute(service string, source string, context string, schemaType string) (outputDTO.SchemaDTO, error) {
	item, err := la.SchemaRepository.FindOneByServiceAndSourceAndContextAndSchemaType(service, source, schemaType)
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
