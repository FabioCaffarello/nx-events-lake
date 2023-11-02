package shared

type MetadataInput struct {
	ID                  string                 `json:"id"`
	Data                map[string]interface{} `json:"data"`
	ProcessingId        string                 `json:"processing_id"`
	ProcessingTimestamp string                 `json:"processing_timestamp"`
	InputSchemaId       string                 `json:"input_schema_id"`
}

type Metadata struct {
	Input               MetadataInput `json:"input"`
	Service             string        `json:"service"`
	Source              string        `json:"source"`
	Context             string        `json:"context"`
	ProcessingTimestamp string        `json:"processing_timestamp"`
	JobFrequency        string        `json:"job_frequency"`
	JobConfigId         string        `json:"job_config_id"`
}

type Status struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}
