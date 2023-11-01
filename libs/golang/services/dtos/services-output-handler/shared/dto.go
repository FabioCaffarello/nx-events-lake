package shared

type MetadataInput struct {
	ID                  string                 `json:"id"`
	Data                map[string]interface{} `json:"data"`
	ProcessingId        string                 `json:"processing_id"`
	ProcessingTimestamp string                 `json:"processing_timestamp"`
}

type Metadata struct {
	InputId string        `json:"input_id"`
	Input   MetadataInput `json:"input"`
	Service string        `json:"service"`
	Source  string        `json:"source"`
}
