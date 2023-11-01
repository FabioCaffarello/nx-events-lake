package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SchemaVersionSuite struct {
	suite.Suite
}


func TestSchemaVersionSuite(t *testing.T) {
	suite.Run(t, new(SchemaVersionSuite))
}

func (suite *SchemaVersionSuite) TestNewSchemaVersionWhenIsANewSchemaVersion() {
	jsonSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type": "string",
			},
		},
	}
    schema, err := NewSchema("schema-type-test", "context-test", "service-test", "source-test", jsonSchema)
    suite.NoError(err)
    suite.NotNil(schema)

    schemaVersion, err := NewSchemaVersion(schema.SchemaType, schema.Context, schema.Service, schema.Source, schema.JsonSchema, schema.SchemaID, schema.CreatedAt, schema.UpdatedAt)
    suite.NoError(err)
    suite.NotNil(schemaVersion)

    assert.Equal(suite.T(), "schema-type-test-service-test-source-test", schemaVersion.ID)
    assert.Equal(suite.T(), "1af4c454-e8d4-5522-8e9f-7935497f5612", schemaVersion.Versions[0].SchemaID)
    assert.Equal(suite.T(), "schema-type-test", schemaVersion.Versions[0].Schema.SchemaType)
}
