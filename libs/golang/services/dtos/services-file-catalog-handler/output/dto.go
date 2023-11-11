package output

type FileCatalogDTO struct {
	ID         string                 `json:"id"`
	Service    string                 `json:"service"`
	Source     string                 `json:"source"`
	Context    string                 `json:"context"`
	LakeLayer  string                 `json:"lake_layer"`
	SchemaType string                 `json:"schema_type"`
	CatalogID  string                 `json:"catalog_id"`
	Catalog    map[string]interface{} `json:"catalog"`
	CreatedAt  string                 `json:"created_at"`
	UpdatedAt  string                 `json:"updated_at"`
}
