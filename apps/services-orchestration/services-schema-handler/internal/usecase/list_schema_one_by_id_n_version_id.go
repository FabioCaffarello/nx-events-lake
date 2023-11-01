package usecase

import (
	"apps/services-orchestration/services-schema-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-schema-handler/output"
)

type ListOneSchemaVersionByIdAndVersionIdUseCase struct {
	SchemaVersionRepository entity.SchemaVersionInterface
}

func NewListOneSchemaVersionByIdAndVersionIdUseCase(
	repository entity.SchemaVersionInterface,
) *ListOneSchemaVersionByIdAndVersionIdUseCase {
	return &ListOneSchemaVersionByIdAndVersionIdUseCase{
		SchemaVersionRepository: repository,
	}
}

func (lcv *ListOneSchemaVersionByIdAndVersionIdUseCase) Execute(id string, versionId string) (outputDTO.SchemaDTO, error) {
	item, err := lcv.SchemaVersionRepository.FindOneByIdAndVersionId(id, versionId)
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
