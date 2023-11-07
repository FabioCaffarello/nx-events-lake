package usecase

import (
	"apps/services-orchestration/services-schema-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-schema-handler/output"
)

type UpdateSchemaVersionUseCase struct {
	SchemaVersionRepository entity.SchemaVersionInterface
}

func NewUpdateSchemaVersionUseCase(
	repository entity.SchemaVersionInterface,
) *UpdateSchemaVersionUseCase {
	return &UpdateSchemaVersionUseCase{
		SchemaVersionRepository: repository,
	}
}

func (uc *UpdateSchemaVersionUseCase) Execute(schema outputDTO.SchemaDTO) (outputDTO.SchemaDTO, error) {
	schemaEntity := ConvertSchemaDTOToEntity(schema)
	err := uc.SchemaVersionRepository.Update(schemaEntity)
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
		CreatedAt:  schemaEntity.CreatedAt,
		UpdatedAt:  schemaEntity.UpdatedAt,
	}

	return dto, nil
}
