package input


import (
	outputDTO "libs/golang/services/dtos/services-schema-handler/output"
)


type SchemaDTO struct {
	SchemaType string                 `json:"schema_type"`
	Service    string                 `json:"service"`
	Source     string                 `json:"source"`
    Context    string                 `json:"context"`
	JsonSchema map[string]interface{} `json:"json_schema"`
}

type SchemaVersionData struct {
	SchemaID string               `json:"schema_id"`
	Schema   *outputDTO.SchemaDTO `json:"schema"`
}

type SchemaVersionDTO struct {
	ID       string              `json:"id"`
	Versions []SchemaVersionData `json:"versions"`
}
