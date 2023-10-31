package configuuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigUUIDSuite struct {
	suite.Suite
}

func TestConfigIDSuite(t *testing.T) {
	suite.Run(t, new(ConfigUUIDSuite))
}

func (suite *ConfigUUIDSuite) TestGenerateConfigID() {
	properties := map[string]interface{}{
		"active":    true,
		"service":   "service",
		"source":    "source",
		"frequency": "frequency",
		"context":   "context",
		"depends_on": []map[string]interface{}{
			{"service": "service", "source": "source"},
		},
		"job_parameters": map[string]interface{}{
			"job_parameter": "job_parameter",
		},
		"service_parameters": map[string]interface{}{
			"service_parameter": "service_parameter",
		},
	}

	configID, err := GenerateConfigID(properties)

	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), configID)
	assert.Equal(suite.T(), "ec65ec83-f75d-5814-bf63-751ceda7ef0b", configID)
}
