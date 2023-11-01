package entity

import (
	"errors"
	"libs/golang/shared/go-id/config"
	configuuid "libs/golang/shared/go-id/config/uuid"
)

var (
	ErrConfigIDEmpty      = errors.New("config id is empty")
	ErrConfigVersionEmpty = errors.New("config version is empty")
	ErrGenerateConfigID   = errors.New("error generating config id")
)

type ConfigData struct {
	ConfigID configuuid.ID `json:"config_id"`
	Config   *Config       `json:"config"`
}

type ConfigVersion struct {
	ID       config.ID    `bson:"id"`
	Versions []ConfigData `json:"versions"`
}

func NewConfigVersion(
	name string,
	active bool,
	frequency string,
	service string,
	source string,
	context string,
	dependsOn []JobDependencies,
	jobParameters map[string]interface{},
	serviceParameters map[string]interface{},
    configID configuuid.ID,
    createdAt string,
    updatedAt string,
) (*ConfigVersion, error) {
	configData := ConfigData{
		ConfigID: configID,
		Config:   &Config{
            ID:                config.NewID(service, source),
            Name:              name,
            Active:            active,
            Frequency:         frequency,
            Service:           service,
            Source:            source,
            Context:           context,
            DependsOn:         dependsOn,
            JobParameters:     jobParameters,
            ServiceParameters: serviceParameters,
            ConfigID:          configID,
            CreatedAt:         createdAt,
            UpdatedAt:         updatedAt,
        },
	}
	configVersion := &ConfigVersion{
		ID:       configData.Config.ID,
		Versions: []ConfigData{configData},
	}
	err := configVersion.IsConfigVersionValid()
	if err != nil {
		return nil, err
	}
	return configVersion, nil
}

func (configVersion *ConfigVersion) IsConfigVersionValid() error {
	if configVersion.ID == "" {
		return ErrConfigIDEmpty
	}
	if len(configVersion.Versions) == 0 {
		return ErrConfigVersionEmpty
	}
	return nil
}
