package usecase

import (
	"apps/services-orchestration/services-schema-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-schema-handler/output"
)

type ListAllSchemasVersionUseCase struct {
	SchemaVersionRepository entity.SchemaVersionInterface
}

func NewListAllSchemasVersionUseCase(
	repository entity.SchemaVersionInterface,
) *ListAllSchemasVersionUseCase {
	return &ListAllSchemasVersionUseCase{
		SchemaVersionRepository: repository,
	}
}

func (la *ListAllSchemasVersionUseCase) Execute() ([]outputDTO.SchemaVersionDTO, error) {
	// FIXME:
	items, err := la.SchemaVersionRepository.FindAll()
	if err != nil {
		return []outputDTO.SchemaVersionDTO{}, err
	}
	var result []outputDTO.SchemaVersionDTO
	for _, item := range items {
		dto := outputDTO.SchemaVersionDTO{
			ID:       string(item.ID),
			Versions: ConvertEntityToUseCaseSchemaVersion(item.Versions),
		}
		result = append(result, dto)
	}
	return result, nil
}
