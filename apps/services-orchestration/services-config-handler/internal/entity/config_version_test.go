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
	config, err := NewConfig("name-test", true, "daily", "service-test", "source-test", "context-test", "event", dependsOn, jobParams, serviceParams)
	suite.NoError(err)
	suite.NotNil(config)

	configVersion, err := NewConfigVersion(config.Name, config.Active, config.Frequency, config.Service, config.Source, config.Context, config.OutputMethod, config.DependsOn, config.JobParameters, config.ServiceParameters, config.ConfigID, config.CreatedAt, config.UpdatedAt)
	suite.NoError(err)
	suite.NotNil(configVersion)

	assert.Equal(suite.T(), "service-test-source-test", configVersion.ID)
	assert.Equal(suite.T(), "8d45af0f-b38f-50ca-95db-2674d5c190f8", configVersion.Versions[0].ConfigID)
	assert.Equal(suite.T(), "name-test", configVersion.Versions[0].Config.Name)
}
