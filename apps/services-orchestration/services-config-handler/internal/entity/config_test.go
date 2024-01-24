package entity

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigSuite struct {
	suite.Suite
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}

func (suite *ConfigSuite) TestNewConfigWhenIsANewConfig() {
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
	config, err := NewConfig("test", true, "daily", "test", "test", "test", "GenerateInputUsingBucketUriAndPartition", "event", dependsOn, jobParams, serviceParams)
	suite.NoError(err)
	suite.NotNil(config)
}

func (suite *ConfigSuite) TestNewConfigWhenNameIsEmpty() {
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
	config, err := NewConfig("", true, "daily", "test", "test", "test", "GenerateInputUsingBucketUriAndPartition", "event", dependsOn, jobParams, serviceParams)
	suite.Error(err)
	suite.Nil(config)
}

func (suite *ConfigSuite) TestNewConfigWhenServiceIsEmpty() {
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
	config, err := NewConfig("test", true, "daily", "", "test", "test", "GenerateInputUsingBucketUriAndPartition", "event", dependsOn, jobParams, serviceParams)
	suite.Error(err)
	suite.Nil(config)
}

func (suite *ConfigSuite) TestNewConfigWhenSourceIsEmpty() {
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
	config, err := NewConfig("test", true, "daily", "test", "", "test", "GenerateInputUsingBucketUriAndPartition", "event", dependsOn, jobParams, serviceParams)
	suite.Error(err)
	suite.Nil(config)
}

func (suite *ConfigSuite) TestNewConfigWhenContextIsEmpty() {
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
	config, err := NewConfig("test", true, "daily", "test", "test", "", "GenerateInputUsingBucketUriAndPartition", "event", dependsOn, jobParams, serviceParams)
	suite.Error(err)
	suite.Nil(config)
}

func (suite *ConfigSuite) TestIsConfigValid() {
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
	config, err := NewConfig("test", true, "daily", "test", "test", "test", "GenerateInputUsingBucketUriAndPartition", "event", dependsOn, jobParams, serviceParams)
	suite.NoError(err)
	err = config.IsConfigValid()
	suite.NoError(err)

  assert.Equal(suite.T(), "test-test", config.ID)
  assert.Equal(suite.T(), "478a21cc-8100-520c-8107-7a07bb4b52bd", config.ConfigID)
}
