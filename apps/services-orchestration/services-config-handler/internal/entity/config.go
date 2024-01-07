package entity

import (
	"errors"
	"libs/golang/shared/go-id/config"
	configuuid "libs/golang/shared/go-id/config/uuid"
	"time"
)

var (
	ErrConfigNameEmpty         = errors.New("config name is empty")
	ErrConfigFrequencyEmpty    = errors.New("config frequency is empty")
	ErrConfigServiceEmpty      = errors.New("config service is empty")
	ErrConfigSourceEmpty       = errors.New("config source is empty")
	ErrConfigContextEmpty      = errors.New("config context is empty")
	ErrConfigOutputMethodEmpty = errors.New("config output method is empty")
	ErrConfigDependsOnEmpty    = errors.New("config dependsOn is empty")
	ErrConfigEmpty             = errors.New("config is empty")
)

type JobDependencies struct {
	Service string `json:"service"`
	Source  string `json:"source"`
}

type Config struct {
	ID                config.ID              `bson:"id"`
	Name              string                 `bson:"name"`
	Active            bool                   `bson:"active"`
	Frequency         string                 `bson:"frequency"`
	Service           string                 `bson:"service"`
	Source            string                 `bson:"source"`
	Context           string                 `bson:"context"`
	OutputMethod      string                 `json:"output_method"`
	DependsOn         []JobDependencies      `bson:"depends_on"`
	ServiceParameters map[string]interface{} `bson:"service_parameters"`
	JobParameters     map[string]interface{} `bson:"job_parameters"`
	ConfigID          configuuid.ID          `bson:"config_id"`
	CreatedAt         string                 `bson:"created_at"`
	UpdatedAt         string                 `bson:"updated_at"`
}

func NewConfig(
	name string,
	active bool,
	frequency string,
	service string,
	source string,
	context string,
	outputMethod string,
	dependsOn []JobDependencies,
	jobParameters map[string]interface{},
	serviceParameters map[string]interface{},
) (*Config, error) {
	config := &Config{
		ID:                config.NewID(service, source),
		Name:              name,
		Active:            active,
		Frequency:         frequency,
		Service:           service,
		Source:            source,
		Context:           context,
		OutputMethod:      outputMethod,
		DependsOn:         dependsOn,
		ServiceParameters: serviceParameters,
		JobParameters:     jobParameters,
		CreatedAt:         time.Now().Format(time.RFC3339),
		UpdatedAt:         time.Now().Format(time.RFC3339),
	}
	err := config.IsConfigValid()
	if err != nil {
		return nil, err
	}
	err = config.SetConfigID()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func convertDependsOn(dependsOn []JobDependencies) []map[string]interface{} {
	mapDependsOn := make([]map[string]interface{}, len(dependsOn))
	for i, dep := range dependsOn {
		mapDependsOn[i] = map[string]interface{}{
			"service": dep.Service,
			"source":  dep.Source,
		}
	}
	return mapDependsOn
}

func getConfigParamsToGenerateVersionID(config *Config) map[string]interface{} {
	return map[string]interface{}{
		"active":             config.Active,
		"service":            config.Service,
		"source":             config.Source,
		"frequency":          config.Frequency,
		"context":            config.Context,
		"output_method":      config.OutputMethod,
		"depends_on":         convertDependsOn(config.DependsOn),
		"job_parameters":     config.JobParameters,
		"service_parameters": config.ServiceParameters,
	}
}

func (config *Config) SetConfigID() error {
	if config == nil {
		return ErrConfigEmpty
	}
	configDataForVersionID := getConfigParamsToGenerateVersionID(config)
	configID, err := configuuid.GenerateConfigID(configDataForVersionID)
	if err != nil {
		return ErrGenerateConfigID
	}
	config.ConfigID = configID
	return nil
}

func (config *Config) IsConfigValid() error {
	if config.Name == "" {
		return ErrConfigNameEmpty
	}
	if config.Frequency == "" {
		return ErrConfigFrequencyEmpty
	}
	if config.Service == "" {
		return ErrConfigServiceEmpty
	}
	if config.Source == "" {
		return ErrConfigSourceEmpty
	}
	if config.Context == "" {
		return ErrConfigContextEmpty
	}
	if config.OutputMethod == "" {
		return ErrConfigOutputMethodEmpty
	}
	return nil
}
