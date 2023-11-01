package entity

type ConfigInterface interface {
	Save(config *Config) error
	FindAll() ([]*Config, error)
	FindAllByService(service string) ([]*Config, error)
	FindOneById(id string) (*Config, error)
	FindAllByDependentJod(service string, source string) ([]*Config, error)
	FindAllByServiceAndContext(service string, contextEnv string) ([]*Config, error)
	FindOneByServiceAndSourceAndContext(service string, source string, contextEnv string) (*Config, error)
}

type ConfigVersionInterface interface {
	Save(configVersion *ConfigVersion) error
	Update(config *Config) error
	FindAll() ([]*ConfigVersion, error)
	FindOneById(id string) (*ConfigVersion, error)
	FindOneByIdAndVersionId(id string, versionId string) (*Config, error)
}
