package entity

import (
	"libs/golang/shared/go-id/md5"
	"libs/golang/shared/go-id/uuid"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type InputEntitySuite struct {
	suite.Suite
}

func TestInputEntitySuite(t *testing.T) {
	suite.Run(t, new(InputEntitySuite))
}

func (suite *InputEntitySuite) TestGivenAnEmptySource_WhenCreateANewInput_ThenShouldReceiveAnError() {
	input := Input{
		Data: map[string]interface{}{"test": "test"},
		Metadata: Metadata{
			Source:  "",
			Service: "file-downloader",
            Context: "test",
		},
	}
	assert.Error(suite.T(), input.IsValid(), "invalid source")
}

func (suite *InputEntitySuite) TestGivenAnEmptyService_WhenCreateANewInput_ThenShouldReceiveAnError() {
	input := Input{
		Data: map[string]interface{}{"test": "test"},
		Metadata: Metadata{
			Source:  "test",
			Service: "",
            Context: "test",
		},
	}
	assert.Error(suite.T(), input.IsValid(), "invalid service")
}

func (suite *InputEntitySuite) TestGivenAnEmptyContext_WhenCreateANewInput_ThenShouldReceiveAnError() {
	input := Input{
		Data: map[string]interface{}{"test": "test"},
		Metadata: Metadata{
			Source:  "test",
			Service: "test",
            Context: "",
		},
	}
	assert.Error(suite.T(), input.IsValid(), "invalid context")
}

func (suite *InputEntitySuite) TestGivenNoSourceAndService_WhenCreateANewInput_ThenShouldReceiveAnError() {
	input := Input{
		Data:     map[string]interface{}{"test": "test"},
		Metadata: Metadata{},
	}
	assert.Error(suite.T(), input.IsValid(), "invalid source and service")
}

func (suite *InputEntitySuite) TestGivenAnEmptyData_WhenCreateANewInput_ThenShouldReceiveAnError() {
	input := Input{}
	assert.Error(suite.T(), input.IsValid(), "invalid data")
}

func (suite *InputEntitySuite) TestGivenAValidParams_WhenCallNewInput_ThenShouldReceiveCreateInputWithAllParams() {
	input := Input{
		Data: map[string]interface{}{"test": "test"},
		Metadata: Metadata{
			Source:  "test",
			Service: "file-downloader",
            Context: "test",
		},
	}
	assert.Equal(suite.T(), map[string]interface{}{"test": "test"}, input.Data)
	assert.Nil(suite.T(), input.IsValid())
}

func (suite *InputEntitySuite) TestGivenAValidParams_WhenCallNewInputFunc_ThenShouldReceiveCreateInputWithAllParams() {
	input, err := NewInput(map[string]interface{}{"test": "test"}, "test", "file-downloader", "test-context")
	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), map[string]interface{}{"test": "test"}, input.Data)
	assert.Equal(suite.T(), md5.ID("6653310993a16531c52a943792dd1767"), input.ID)
	assert.Equal(suite.T(), "test", input.Metadata.Source)
	assert.Equal(suite.T(), "file-downloader", input.Metadata.Service)
    assert.Equal(suite.T(), "test-context", input.Metadata.Context)
	assert.Equal(suite.T(), Status{Code: 0, Detail: ""}, input.Status)

	assert.IsType(suite.T(), Metadata{}, input.Metadata)
	assert.IsType(suite.T(), uuid.ID{}, input.Metadata.ProcessingId)
	assert.Equal(suite.T(), time.Now().Format(time.RFC3339), input.Metadata.ProcessingTimestamp)

	assert.Nil(suite.T(), input.IsValid())
}
