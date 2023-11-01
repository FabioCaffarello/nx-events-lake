package output

import (
	sharedDTO "libs/golang/services/dtos/services-output-handler/shared"
)

type ServiceOutputDTO struct {
	ID        string                 `json:"id"`
	Data      map[string]interface{} `json:"data"`
	Metadata  sharedDTO.Metadata     `json:"metadata"`
	Service   string                 `json:"service"`
	Source    string                 `json:"source"`
	Context   string                 `json:"context"`
	CreatedAt string                 `json:"created_at"`
	UpdatedAt string                 `json:"updated_at"`
}
