package pipelineid

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GoidPipelineIdSuite struct {
	suite.Suite
}

func TestGoidSchemaidSuite(t *testing.T) {
	suite.Run(t, new(GoidPipelineIdSuite))
}

func (suite *GoidPipelineIdSuite) TestNewID() {
	id := NewID("br", "service", "source")

	assert.Equal(suite.T(), "br-service-source", id, "NewID() returned an incorrect ID")
}
