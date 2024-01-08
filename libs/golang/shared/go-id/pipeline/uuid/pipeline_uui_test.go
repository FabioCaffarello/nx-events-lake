package pipelineuuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PipelineUUIDSuite struct {
	suite.Suite
}

func TestPipelineIDSuite(t *testing.T) {
	suite.Run(t, new(PipelineUUIDSuite))
}

func (suite *PipelineUUIDSuite) TestGeneratePipelineID() {
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

	pipelineID, err := GeneratePipelineID(properties)

	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), pipelineID)
	assert.Equal(suite.T(), "ec65ec83-f75d-5814-bf63-751ceda7ef0b", pipelineID)
}
