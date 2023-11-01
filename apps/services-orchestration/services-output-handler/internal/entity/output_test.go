package entity

import (
	"libs/golang/shared/go-id/md5"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceOutputEntitySuite struct {
	suite.Suite
}

func TestServiceOutputEntitySuite(t *testing.T) {
	suite.Run(t, new(ServiceOutputEntitySuite))
}

func (suite *ServiceOutputEntitySuite) TestGivenAnEmptyService_WhenCreateANewServiceOutput_ThenShouldReceiveAnError() {
	serviceOutput, err := NewServiceOutput(map[string]interface{}{"test": "test"}, "test", map[string]interface{}{"test": "test"}, "test", "test", "", "test", "test")
	suite.Nil(serviceOutput)
	suite.Error(err, "invalid service output")
}

func (suite *ServiceOutputEntitySuite) TestGivenAnEmptySource_WhenCreateANewServiceOutput_ThenShouldReceiveAnError() {
	serviceOutput, err := NewServiceOutput(map[string]interface{}{"test": "test"}, "test", map[string]interface{}{"test": "test"}, "test", "test", "file-downloader", "", "test")
	suite.Nil(serviceOutput)
	suite.Error(err, "invalid service output")
}

func (suite *ServiceOutputEntitySuite) TestGivenAnEmptyData_WhenCreateANewServiceOutput_ThenShouldReceiveAnError() {
	serviceOutput, err := NewServiceOutput(nil, "test", map[string]interface{}{"test": "test"}, "test", "test", "file-downloader", "test", "test")
	suite.Nil(serviceOutput)
	suite.Error(err, "invalid service output")
}

func (suite *ServiceOutputEntitySuite) TestGivenNoSourceAndService_WhenCreateANewServiceOutput_ThenShouldReceiveAnError() {
	serviceOutput, err := NewServiceOutput(map[string]interface{}{"test": "test"}, "", map[string]interface{}{"test": "test"}, "test", "test", "", "", "test")
	suite.Nil(serviceOutput)
	suite.Error(err, "invalid service output")
}

func (suite *ServiceOutputEntitySuite) TestGivenAnEmptyInputId_WhenCreateANewServiceOutput_ThenShouldReceiveAnError() {
	serviceOutput, err := NewServiceOutput(map[string]interface{}{"test": "test"}, "", map[string]interface{}{"test": "test"}, "test", "test", "file-downloader", "test", "test")
	suite.Nil(serviceOutput)
	suite.Error(err, "invalid service output")
}

func (suite *ServiceOutputEntitySuite) TestGivenAnEmptyInputData_WhenCreateANewServiceOutput_ThenShouldReceiveAnError() {
	serviceOutput, err := NewServiceOutput(map[string]interface{}{"test": "test"}, "test", nil, "test", "test", "file-downloader", "test", "test")
	suite.Nil(serviceOutput)
	suite.Error(err, "invalid service output")
}

func (suite *ServiceOutputEntitySuite) TestGivenAnEmptyInputProcessingId_WhenCreateANewServiceOutput_ThenShouldReceiveAnError() {
	serviceOutput, err := NewServiceOutput(map[string]interface{}{"test": "test"}, "test", map[string]interface{}{"test": "test"}, "", "test", "file-downloader", "test", "test")
	suite.Nil(serviceOutput)
	suite.Error(err, "invalid service output")
}

func (suite *ServiceOutputEntitySuite) TestGivenAnEmptyInputProcessingTimestamp_WhenCreateANewServiceOutput_ThenShouldReceiveAnError() {
	serviceOutput, err := NewServiceOutput(map[string]interface{}{"test": "test"}, "test", map[string]interface{}{"test": "test"}, "test", "", "file-downloader", "test", "test")
	suite.Nil(serviceOutput)
	suite.Error(err, "invalid service output")
}

func (suite *ServiceOutputEntitySuite) TestGivenAnEmptyContext_WhenCreateANewServiceOutput_ThenShouldReceiveAnError() {
	serviceOutput, err := NewServiceOutput(map[string]interface{}{"test": "test"}, "test", map[string]interface{}{"test": "test"}, "test", "test", "file-downloader", "test", "")
	suite.Nil(serviceOutput)
	suite.Error(err, "invalid service output")
}

func (suite *ServiceOutputEntitySuite) TestGivenAValidParams_WhenCallNewServiceOutput_ThenShouldReceiveCreateServiceOutputWithAllParams() {
	serviceOutput, err := NewServiceOutput(map[string]interface{}{"test": "test"}, "test", map[string]interface{}{"test": "test"}, "test", "test", "file-downloader", "test", "test")
	suite.NoError(err)
	suite.Equal(map[string]interface{}{"test": "test"}, serviceOutput.Data)
	suite.Equal(md5.ID("763981cb596da0f9e847db8e1319f9f2"), serviceOutput.ID)
	suite.Equal("test", serviceOutput.Metadata.Source)
	suite.Equal("file-downloader", serviceOutput.Metadata.Service)
	suite.Equal("test", serviceOutput.Metadata.InputID)
	suite.Equal(map[string]interface{}{"test": "test"}, serviceOutput.Metadata.Input.Data)
	suite.Equal("test", serviceOutput.Metadata.Input.ProcessingId)
	suite.Equal("test", serviceOutput.Metadata.Input.ProcessingTimestamp)
	suite.Equal("test", serviceOutput.Context)
	suite.Equal("file-downloader", serviceOutput.Service)
	suite.Equal("test", serviceOutput.Source)
	suite.NotEmpty(serviceOutput.CreatedAt)
	suite.NotEmpty(serviceOutput.UpdatedAt)
}
