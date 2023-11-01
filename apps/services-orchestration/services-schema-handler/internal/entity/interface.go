package entity

type SchemaInterface interface {
     Save(schema *Schema) error
     FindOneById(id string) (*Schema, error)
     FindAll() ([]*Schema, error)
     FindAllByService(service string) ([]*Schema, error)
     FindOneByServiceSourceAndSchemaType(service string, source string, schemaType string) (*Schema, error)
     FindAllByServiceAndContext(service string, contextEnv string) ([]*Schema, error)
     FindOneByServiceAndSourceAndContextAndSchemaType(service string, source string, contextEnv string) (*Schema, error)
}

type SchemaVersionInterface interface {
	Save(configVersion *SchemaVersion) error
	Update(config *Schema) error
	FindAll() ([]*SchemaVersion, error)
	FindOneById(id string) (*SchemaVersion, error)
	FindOneByIdAndVersionId(id string, versionId string) (*Schema, error)
}
