package output

import (
    sharedDTO "libs/golang/services/dtos/services-input-handler/shared"
)


type InputDTO struct {
	ID       string                 `json:"id"`
	Data     map[string]interface{} `json:"data"`
	Metadata sharedDTO.Metadata               `json:"metadata"`
	Status   sharedDTO.Status                 `json:"status"`
}
