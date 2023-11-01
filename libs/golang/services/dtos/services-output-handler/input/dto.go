package input

import (
	sharedDTO "libs/golang/services/dtos/services-output-handler/shared"
)

type ServiceOutputDTO struct {
	Data     map[string]interface{} `json:"data"`
	Metadata sharedDTO.Metadata     `json:"metadata"`
	Context  string                 `json:"context"`
}
