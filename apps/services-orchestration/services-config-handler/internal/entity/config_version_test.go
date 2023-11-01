package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigVersionSuite struct {
	suite.Suite
}

func TestConfigVersionSuite(t *testing.T) {
	suite.Run(t, new(ConfigVersionSuite))
}

func (suite *ConfigVersionSuite) TestNewConfigVersionWhenIsANewConfigVersion() {
	jobParams := map[string]interface{}{
		"test": "test",
	}
	serviceParams := map[string]interface{}{
		"test": "test",
	}
	dependsOn := []JobDependencies{
		{
			Service: "test",
			Source:  "test",
		},
	}
	config, err := NewConfig("name-test", true, "daily", "service-test", "source-test", "context-test", dependsOn, jobParams, serviceParams)
	suite.NoError(err)
	suite.NotNil(config)

	configVersion, err := NewConfigVersion(config.Name, config.Active, config.Frequency, config.Service, config.Source, config.Context, config.DependsOn, config.JobParameters, config.ServiceParameters, config.ConfigID, config.CreatedAt, config.UpdatedAt)
	suite.NoError(err)
	suite.NotNil(configVersion)

	assert.Equal(suite.T(), "service-test-source-test", configVersion.ID)
	assert.Equal(suite.T(), "5ced1d38-d35c-5bea-9f7a-46f9680d681a", configVersion.Versions[0].ConfigID)
	assert.Equal(suite.T(), "name-test", configVersion.Versions[0].Config.Name)
}
