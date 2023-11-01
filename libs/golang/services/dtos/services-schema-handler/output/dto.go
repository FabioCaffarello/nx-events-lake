package output

type SchemaDTO struct {
	ID         string                 `json:"id"`
	SchemaType string                 `json:"schema_type"`
	Service    string                 `json:"service"`
	Source     string                 `json:"source"`
    Context    string                 `json:"context"`
	JsonSchema map[string]interface{} `json:"json_schema"`
	SchemaID   string                 `json:"schema_id"`
	CreatedAt  string                 `json:"created_at"`
	UpdatedAt  string                 `json:"updated_at"`
}

type SchemaVersionData struct {
	SchemaID string     `json:"schema_id"`
	Schema   *SchemaDTO `json:"schema"`
}

type SchemaVersionDTO struct {
	ID       string              `json:"id"`
	Versions []SchemaVersionData `json:"versions"`
}
