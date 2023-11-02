package input

import (
	sharedDTO "libs/golang/services/dtos/services-events-handler/shared"
)

type ServiceFeedbackDTO struct {
	Data     map[string]interface{} `json:"data"`
	Metadata sharedDTO.Metadata     `json:"metadata"`
	Status   sharedDTO.Status       `json:"status"`
}
