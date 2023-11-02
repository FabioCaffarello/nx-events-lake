package shared

type Metadata struct {
	ProcessingId        string `json:"processing_id"`
	ProcessingTimestamp string `json:"processing_timestamp"`
	Context             string `json:"context"`
	Source              string `json:"source"`
	Service             string `json:"service"`
}

type Status struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}
