package usecase

import (
	"apps/services-orchestration/services-schema-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-schema-handler/output"
)

type CreateSchemaVersionUseCase struct {
	SchemaVersionRepository entity.SchemaVersionInterface
}

func NewCreateSchemaVersionUseCase(
	repository entity.SchemaVersionInterface,
) *CreateSchemaVersionUseCase {
	return &CreateSchemaVersionUseCase{
		SchemaVersionRepository: repository,
	}
}

func (ccu *CreateSchemaVersionUseCase) Execute(schema outputDTO.SchemaDTO) (outputDTO.SchemaVersionDTO, error) {
	schemaVersionEntity, err := entity.NewSchemaVersion(
		schema.SchemaType,
		schema.Context,
		schema.Service,
		schema.Source,
		schema.JsonSchema,
		schema.SchemaID,
		schema.CreatedAt,
		schema.UpdatedAt,
	)

	if err != nil {
		return outputDTO.SchemaVersionDTO{}, err
	}

	err = ccu.SchemaVersionRepository.Save(schemaVersionEntity)
	if err != nil {
		return outputDTO.SchemaVersionDTO{}, err
	}

	dto := outputDTO.SchemaVersionDTO{
		ID:       string(schemaVersionEntity.ID),
		Versions: ConvertEntityToUseCaseSchemaVersion(schemaVersionEntity.Versions),
	}

	return dto, nil

}
