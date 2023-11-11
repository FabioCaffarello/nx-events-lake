package input

type FileCatalogDTO struct {
     Service    string                 `json:"service"`
     Source     string                 `json:"source"`
     Context    string                 `json:"context"`
     LakeLayer  string                 `json:"lake_layer"`
     SchemaType string                 `json:"schema_type"`
     Catalog    map[string]interface{} `json:"catalog"`
}
