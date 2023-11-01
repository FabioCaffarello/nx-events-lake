package database

import (
	"context"
	"testing"
	"time"

	"apps/services-orchestration/services-output-handler/internal/entity"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceOutputRepositorySuite struct {
	suite.Suite
	Client     *mongo.Client
	Database   string
	Collection string
	Source     string
	repo       *ServiceOutputRepository
}

func TestServiceOutputRepositorySuite(t *testing.T) {
	suite.Run(t, new(ServiceOutputRepositorySuite))
}

func (suite *ServiceOutputRepositorySuite) SetupTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	mongoURI := "mongodb://localhost:27017"
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	suite.Client = client
	suite.Database = "test-database"
	suite.Collection = "test-service"
	suite.Source = "test-source"
	suite.repo = NewServiceOutputRepository(suite.Client, suite.Database)
}

func (suite *ServiceOutputRepositorySuite) TearDownTest() {
	suite.Client.Database(suite.Database).Drop(context.Background())
	err := suite.Client.Disconnect(context.Background())
	suite.NoError(err)
}

func (suite *ServiceOutputRepositorySuite) TestSaveWhenIsANewServiceOutput() {
	serviceOutput, err := entity.NewServiceOutput(map[string]interface{}{"test": "test"}, "test", map[string]interface{}{"test": "test"}, "test", "test", "file-downloader", "test", "test")
	suite.NoError(err)
	err = suite.repo.Save(serviceOutput, suite.Collection)
	suite.NoError(err)
}

func (suite *ServiceOutputRepositorySuite) TestFindOneByIdAndService() {
	serviceOutput, err := entity.NewServiceOutput(map[string]interface{}{"test": "test"}, "test", map[string]interface{}{"test": "test"}, "test", "test", "file-downloader", "test", "test")
	suite.NoError(err)
	err = suite.repo.Save(serviceOutput, suite.Collection)
	suite.NoError(err)
	result, err := suite.repo.FindOneByIdAndService(string(serviceOutput.ID), suite.Collection)
	suite.NoError(err)
	suite.Equal(serviceOutput.ID, result.ID)
	suite.Equal(serviceOutput.Data, result.Data)
	suite.Equal(serviceOutput.Metadata, result.Metadata)
}

func (suite *ServiceOutputRepositorySuite) TestFindAllByService() {
	serviceOutput1, err := entity.NewServiceOutput(map[string]interface{}{"test": "test"}, "test", map[string]interface{}{"test": "test"}, "test", "test", suite.Collection, "test", "test")
	suite.NoError(err)
	err = suite.repo.Save(serviceOutput1, suite.Collection)
	suite.NoError(err)

	serviceOutput2, err := entity.NewServiceOutput(map[string]interface{}{"test": "test"}, "test", map[string]interface{}{"test": "test"}, "test", "test", suite.Collection, "test2", "test")
	suite.NoError(err)
	err = suite.repo.Save(serviceOutput2, suite.Collection)
	suite.NoError(err)

	results, err := suite.repo.FindAllByService(suite.Collection)
	suite.NoError(err)
	suite.Len(results, 2)
}

func (suite *ServiceOutputRepositorySuite) TestFindAllByServiceAndSource() {
	serviceOutput1, err := entity.NewServiceOutput(map[string]interface{}{"test1": "test"}, "test", map[string]interface{}{"test": "test"}, "test", "test", suite.Collection, suite.Source, "test")
	suite.NoError(err)
	err = suite.repo.Save(serviceOutput1, suite.Collection)
	suite.NoError(err)

	serviceOutput2, err := entity.NewServiceOutput(map[string]interface{}{"test2": "test"}, "test", map[string]interface{}{"test": "test"}, "test", "test", suite.Collection, suite.Source, "test")
	suite.NoError(err)
	err = suite.repo.Save(serviceOutput2, suite.Collection)
	suite.NoError(err)

	results, err := suite.repo.FindAllByServiceAndSource(suite.Collection, suite.Source)
	suite.NoError(err)
	suite.Len(results, 2)
}
