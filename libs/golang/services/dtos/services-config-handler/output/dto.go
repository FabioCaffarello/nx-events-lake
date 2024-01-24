package output

import (
	sharedDTO "libs/golang/services/dtos/services-config-handler/shared"
)

type ConfigDTO struct {
	ID                string                      `json:"id"`
	Name              string                      `json:"name"`
	Active            bool                        `json:"active"`
	Frequency         string                      `json:"frequency"`
	Service           string                      `json:"service"`
	Source            string                      `json:"source"`
	Context           string                      `json:"context"`
    InputMethod       string                      `json:"input_method"`
	OutputMethod      string                      `json:"output_method"`
	DependsOn         []sharedDTO.JobDependencies `json:"depends_on"`
	ConfigID          string                      `json:"config_id"`
	ServiceParameters map[string]interface{}      `json:"service_parameters"`
	JobParameters     map[string]interface{}      `json:"job_parameters"`
	CreatedAt         string                      `json:"created_at"`
	UpdatedAt         string                      `json:"updated_at"`
}

type ConfigVersionData struct {
	ConfigID string     `json:"config_id"`
	Config   *ConfigDTO `json:"config"`
}

type ConfigVersionDTO struct {
	ID       string              `json:"id"`
	Versions []ConfigVersionData `json:"versions"`
}
