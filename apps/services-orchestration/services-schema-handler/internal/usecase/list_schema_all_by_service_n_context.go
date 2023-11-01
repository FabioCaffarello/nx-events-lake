package usecase

import (
	"apps/services-orchestration/services-schema-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-schema-handler/output"
)

type ListAllSchemasByServiceAndContextUseCase struct {
	SchemaRepository entity.SchemaInterface
}

func NewListAllSchemasByServiceAndContextUseCase(
	repository entity.SchemaInterface,
) *ListAllSchemasByServiceAndContextUseCase {
	return &ListAllSchemasByServiceAndContextUseCase{
		SchemaRepository: repository,
	}
}

func (la *ListAllSchemasByServiceAndContextUseCase) Execute(service string, contextEnv string) ([]outputDTO.SchemaDTO, error) {
	items, err := la.SchemaRepository.FindAllByServiceAndContext(service, contextEnv)
	if err != nil {
		return nil, err
	}

	var output []outputDTO.SchemaDTO
	for _, item := range items {
		output = append(output, outputDTO.SchemaDTO{
			ID:         item.ID,
			SchemaType: item.SchemaType,
			Service:    item.Service,
			Source:     item.Source,
			Context:    item.Context,
			JsonSchema: item.JsonSchema,
			SchemaID:   item.SchemaID,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
		})
	}

	return output, nil
}
