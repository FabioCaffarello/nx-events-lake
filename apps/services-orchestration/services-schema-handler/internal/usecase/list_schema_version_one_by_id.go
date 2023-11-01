package usecase

import (
	"apps/services-orchestration/services-schema-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-schema-handler/output"
)

type ListOneSchemaVersionByIdUseCase struct {
	SchemaVersionRepository entity.SchemaVersionInterface
}

func NewListOneSchemaVersionByIdUseCase(
	repository entity.SchemaVersionInterface,
) *ListOneSchemaVersionByIdUseCase {
	return &ListOneSchemaVersionByIdUseCase{
		SchemaVersionRepository: repository,
	}
}

func (lcv *ListOneSchemaVersionByIdUseCase) Execute(id string) (outputDTO.SchemaVersionDTO, error) {
	item, err := lcv.SchemaVersionRepository.FindOneById(id)
	if err != nil {
		return outputDTO.SchemaVersionDTO{}, err
	}

	dto := outputDTO.SchemaVersionDTO{
		ID:       string(item.ID),
		Versions: ConvertEntityToUseCaseSchemaVersion(item.Versions),
	}

	return dto, nil
}
