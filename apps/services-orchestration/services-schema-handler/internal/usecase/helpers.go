package usecase

import (
	"apps/services-orchestration/services-schema-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-schema-handler/output"
)

func ConvertEntityToUseCaseSchemaVersion(entityDeps []entity.SchemaData) []outputDTO.SchemaVersionData {
	usecaseDeps := make([]outputDTO.SchemaVersionData, len(entityDeps))
	for i, dep := range entityDeps {
		usecaseDeps[i] = outputDTO.SchemaVersionData{
			SchemaID: dep.SchemaID,
			Schema: &outputDTO.SchemaDTO{
				ID:         string(dep.Schema.ID),
				SchemaType: dep.Schema.SchemaType,
				Service:    dep.Schema.Service,
				Source:     dep.Schema.Source,
				Context:    dep.Schema.Context,
				JsonSchema: dep.Schema.JsonSchema,
				SchemaID:   dep.Schema.SchemaID,
				CreatedAt:  dep.Schema.CreatedAt,
				UpdatedAt:  dep.Schema.UpdatedAt,
			},
		}
	}
	return usecaseDeps
}

func ConvertSchemaDTOToEntity(schema outputDTO.SchemaDTO) *entity.Schema {
	return &entity.Schema{
		ID:         string(schema.ID),
		SchemaType: schema.SchemaType,
		Service:    schema.Service,
		Source:     schema.Source,
		Context:    schema.Context,
		JsonSchema: schema.JsonSchema,
		SchemaID:   schema.SchemaID,
		CreatedAt:  schema.CreatedAt,
		UpdatedAt:  schema.UpdatedAt,
	}
}
