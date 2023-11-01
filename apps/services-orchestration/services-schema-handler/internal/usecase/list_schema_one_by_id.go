package usecase

import (
	"apps/services-orchestration/services-schema-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-schema-handler/output"
)

type ListOneSchemaByIdUseCase struct {
	SchemaRepository entity.SchemaInterface
}

func NewListOneSchemaByIdUseCase(
	repository entity.SchemaInterface,
) *ListOneSchemaByIdUseCase {
	return &ListOneSchemaByIdUseCase{
		SchemaRepository: repository,
	}
}

func (la *ListOneSchemaByIdUseCase) Execute(id string) (outputDTO.SchemaDTO, error) {
	item, err := la.SchemaRepository.FindOneById(id)
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
