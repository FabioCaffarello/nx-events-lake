package entity

type ProcessingGraph struct {
	Service             string                 `bson:"service"`
	Source              string                 `bson:"source"`
	Context             string                 `bson:"context"`
	ProcessingId        string                 `bson:"processing_id"`
	StatusCode          int                    `bson:"status_code"`
	ConfigId            string                 `bson:"config_id"`
	ConfigVersionId     string                 `bson:"config_version_id"`
	SchemaId            string                 `bson:"schema_id"`
	SchemaVersionId     string                 `bson:"schema_version_id"`
	ProcessingTimestamp string                 `bson:"processing_timestamp"`
	InputID             string                 `bson:"input_id"`
	OutputID            string                 `bson:"output_id"`
	OutputData          map[string]interface{} `bson:"output_data"`
}

type Batch struct {
	ID              string                 `bson:"id"`
	Service         string                 `bson:"service"`
	Source          string                 `bson:"source"`
	Context         string                 `bson:"context"`
	ProcessingId    string                 `bson:"processing_id"`
	ProcessingGraph []ProcessingGraph      `bson:"processing_graph"`
	InputID         string                 `bson:"input_id"`
	OutputID        string                 `bson:"output_id"`
	OutputData      map[string]interface{} `bson:"output_data"`
	CreatedAt       string                 `bson:"created_at"`
	UpdatedAt       string                 `bson:"updated_at"`
}

// If there is no dependency dont need to check the batch
// before trigger the dependent jobs the batch should check if all jobs have been processed
// when the previous job trigger returns 0 the batch should be updated
// when the previous job trigger returns 200 the batch should be updated
// when a current job input returns 0 the batch should be checked if the processing staging can be updated
