package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SchemaSuite struct {
	suite.Suite
}

func TestSchemaSuite(t *testing.T) {
	suite.Run(t, new(SchemaSuite))
}

func (suite *SchemaSuite) TestNewSchema() {
	jsonSchema := map[string]interface{}{
		"field1": map[string]interface{}{
			"type": "string",
		},
		"field2": map[string]interface{}{
			"type": "string",
		},
	}

	schemaType := "example"
	service := "test"
	source := "source"
	context := "context"

	schemaInstance, err := NewSchema(schemaType, context, service, source, jsonSchema)

	suite.NoError(err)
	suite.NotNil(schemaInstance)
	suite.NotNil(schemaInstance.ID)
	suite.Equal(schemaType, schemaInstance.SchemaType)
	suite.Equal(jsonSchema, schemaInstance.JsonSchema)
	suite.NotNil(schemaInstance.SchemaID)
}

func (suite *SchemaSuite) TestIsSchemaValid_ValidSchema() {
	jsonSchema := map[string]interface{}{
		"field1": map[string]interface{}{
			"type": "string",
		},
		"field2": map[string]interface{}{
			"type": "string",
		},
	}

	schemaInstance := &Schema{
		SchemaType: "example",
		Service:    "test",
		Source:     "source",
		Context:    "context",
		JsonSchema: jsonSchema,
	}

	err := schemaInstance.IsSchemaValid()

	suite.NoError(err)
}

func (suite *SchemaSuite) TestIsSchemaValid_EmptyService() {
	schemaInstance := &Schema{
		SchemaType: "example",
		Source:     "source",
		Context:    "context",
		JsonSchema: map[string]interface{}{
			"field1": map[string]interface{}{
				"type": "string",
			},
		},
	}

	err := schemaInstance.IsSchemaValid()

	suite.Error(err)
	suite.EqualError(err, "service is empty")
}

func (suite *SchemaSuite) TestIsSchemaValid_EmptySource() {
	schemaInstance := &Schema{
		SchemaType: "example",
		Service:    "test",
		Context:    "context",
		JsonSchema: map[string]interface{}{
			"field1": map[string]interface{}{
				"type": "string",
			},
		},
	}

	err := schemaInstance.IsSchemaValid()

	suite.Error(err)
	suite.EqualError(err, "source is empty")
}

func (suite *SchemaSuite) TestIsSchemaValid_EmptySchemaType() {
	schemaInstance := &Schema{
		Service:    "test",
		Source:     "source",
		Context:    "context",
		JsonSchema: map[string]interface{}{},
	}

	err := schemaInstance.IsSchemaValid()

	suite.Error(err)
	suite.EqualError(err, "schema type is empty")
}

func (suite *SchemaSuite) TestIsSchemaValid_EmptyContext() {
	schemaInstance := &Schema{
		SchemaType: "example",
		Service:    "test",
		Source:     "source",
		JsonSchema: map[string]interface{}{},
	}

	err := schemaInstance.IsSchemaValid()

	suite.Error(err)
	suite.EqualError(err, "context is empty")
}

func (suite *SchemaSuite) TestIsSchemaValid_EmptyJsonSchema() {
	schemaInstance := &Schema{
		SchemaType: "example",
		Service:    "test",
		Source:     "source",
		Context:    "context",
		JsonSchema: nil,
	}

	err := schemaInstance.IsSchemaValid()

	suite.Error(err)
	suite.EqualError(err, "json schema is empty")
}

func (suite *SchemaSuite) TestValidateJSONSchema_ValidSchema() {
	jsonSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type": "string",
			},
		},
	}

	err := ValidateJSONSchema(jsonSchema)

	assert.NoError(suite.T(), err)
}

func (suite *SchemaSuite) TestValidateJSONSchema_InvalidSchema() {
	invalidJsonSchema := map[string]interface{}{
		"type": "invalid_type",
	}

	err := ValidateJSONSchema(invalidJsonSchema)

	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "jsonSchema validation failed")
}
