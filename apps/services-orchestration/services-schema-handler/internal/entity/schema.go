package entity

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	schemaId "libs/golang/shared/go-id/schema/id"
	uuid "libs/golang/shared/go-id/schema/uuid"

	"github.com/xeipuuv/gojsonschema"
)

var (
	ErrSchemaTypeEmpty   = errors.New("schema type is empty")
	ErrServiceEmpty      = errors.New("service is empty")
	ErrSourceEmpty       = errors.New("source is empty")
	ErrContextEmpty      = errors.New("context is empty")
	ErrJsonSchemaEmpty   = errors.New("json schema is empty")
	ErrJsonSchemaInvalid = errors.New("json schema is invalid")
)

const metaschemaURL = "http://json-schema.org/draft-07/schema#"

type Schema struct {
	ID         schemaId.ID            `json:"id"`
	SchemaType string                 `bson:"schema_type"`
	JsonSchema map[string]interface{} `bson:"json_schema"`
	SchemaID   uuid.ID                `bson:"schema_id"`
	Service    string                 `bson:"service"`
	Source     string                 `bson:"source"`
	Context    string                 `bson:"context"`
	CreatedAt  string
	UpdatedAt  string
}

func NewSchema(schemaType string, context string, service string, source string, jsonSchema map[string]interface{}) (*Schema, error) {
	schemauuId, err := uuid.GenerateSchemaID(schemaType, jsonSchema)
	if err != nil {
		return nil, err
	}

	schema := &Schema{
		ID:         schemaId.NewID(schemaType, service, source),
		SchemaType: schemaType,
		JsonSchema: jsonSchema,
		Service:    service,
		Source:     source,
		Context:    context,
		SchemaID:   schemauuId,
		CreatedAt:  time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}
	err = schema.IsSchemaValid()
	if err != nil {
		return nil, err
	}
	return schema, nil
}

func ValidateJSONSchema(jsonSchema map[string]interface{}) error {
	// Convert the JSON schema map to a JSON string
	jsonSchemaBytes, err := json.Marshal(jsonSchema)
	if err != nil {
		return err
	}

	// Create a JSONLoader for the JSON schema
	schemaLoader := gojsonschema.NewStringLoader(string(jsonSchemaBytes))

	// Validate the JSON Schema structure using the gojsonschema library
	metaschemaLoader := gojsonschema.NewReferenceLoader(metaschemaURL)
	compileResult, err := gojsonschema.Validate(metaschemaLoader, schemaLoader)
	if err != nil {
		return err
	}

	if !compileResult.Valid() {
		validationErrors := compileResult.Errors()
		errorMessages := make([]string, len(validationErrors))
		for i, err := range validationErrors {
			errorMessages[i] = err.String()
		}
		return errors.New("jsonSchema validation failed: " + strings.Join(errorMessages, ", "))
	}

	return nil
}

func (schema *Schema) IsSchemaValid() error {
	if schema.SchemaType == "" {
		return ErrSchemaTypeEmpty
	}
	if schema.Service == "" {
		return ErrServiceEmpty
	}
	if schema.Source == "" {
		return ErrSourceEmpty
	}
	if schema.Context == "" {
		return ErrContextEmpty
	}
	if schema.JsonSchema == nil {
		return ErrJsonSchemaEmpty
	}
	err := ValidateJSONSchema(schema.JsonSchema)
	if err != nil {
		return ErrJsonSchemaInvalid
	}
	return nil
}
