package input

import (
	outputDTO "libs/golang/services/dtos/services-config-handler/output"
	sharedDTO "libs/golang/services/dtos/services-config-handler/shared"
)

type ConfigDTO struct {
	Name              string                      `json:"name"`
	Active            bool                        `json:"active"`
	Frequency         string                      `json:"frequency"`
	Service           string                      `json:"service"`
	Source            string                      `json:"source"`
	Context           string                      `json:"context"`
	OutputMethod      string                      `json:"output_method"`
	DependsOn         []sharedDTO.JobDependencies `json:"depends_on"`
	ServiceParameters map[string]interface{}      `json:"service_parameters"`
	JobParameters     map[string]interface{}      `json:"job_parameters"`
}

type ConfigVersionData struct {
	ConfigID string               `json:"config_id"`
	Config   *outputDTO.ConfigDTO `json:"config"`
}

type ConfigVersionDTO struct {
	ID       string              `json:"id"`
	Versions []ConfigVersionData `json:"versions"`
}
