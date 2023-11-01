package entity

import (
	"errors"
	schemaId "libs/golang/shared/go-id/schema/id"
	schemauuid "libs/golang/shared/go-id/schema/uuid"
)

var (
    ErrSchemaIDEmpty      = errors.New("schema id is empty")
    ErrSchemaVersionEmpty = errors.New("schema version is empty")
)

type SchemaData struct {
	SchemaID schemauuid.ID `json:"schema_id"`
	Schema   *Schema       `json:"schema"`
}

type SchemaVersion struct {
	ID       schemaId.ID    `bson:"id"`
	Versions []SchemaData `json:"versions"`
}


func NewSchemaVersion(
    schemaType string,
    context string,
    service string,
    source string,
    jsonSchema map[string]interface{},
    schemaID schemauuid.ID,
    createdAt string,
    updatedAt string,
) (*SchemaVersion, error) {
    schemaData := SchemaData{
        SchemaID: schemaID,
        Schema:   &Schema{
            ID:         schemaId.NewID(schemaType, service, source),
            SchemaType: schemaType,
            JsonSchema: jsonSchema,
            Service:    service,
            Source:     source,
            Context:    context,
            SchemaID:   schemaID,
            CreatedAt:  createdAt,
            UpdatedAt:  updatedAt,
        },
    }

    schemaVersion := &SchemaVersion{
        ID:       schemaData.Schema.ID,
        Versions: []SchemaData{schemaData},
    }
    err := schemaVersion.IsSchemaVersionValid()
    if err != nil {
        return nil, err
    }

    return schemaVersion, nil

}

func (schemaVersion *SchemaVersion) IsSchemaVersionValid() error {
    if schemaVersion.ID == "" {
        return ErrSchemaIDEmpty
    }
    if len(schemaVersion.Versions) == 0 {
        return ErrSchemaVersionEmpty
    }
    return nil
}
