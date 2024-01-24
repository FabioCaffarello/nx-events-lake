package database

import (
	"apps/services-orchestration/services-jobs-handler/internal/entity"
	"context"
	"errors"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrorExistingHttpGatewayParamsVersion = errors.New("http gateway params version already exists")
)

type HttpGatewayParamsVersionRepository struct {
	log        *log.Logger
	Client     *mongo.Client
	Database   string
	Collection *mongo.Collection
}

func NewHttpGatewayParamsVersionRepository(client *mongo.Client, database string) *HttpGatewayParamsVersionRepository {
	return &HttpGatewayParamsVersionRepository{
		log:        log.New(os.Stdout, "[HTTP-GATEWAY-PARAMS-VERSION-REPOSITORY] ", log.LstdFlags),
		Client:     client,
		Database:   database,
		Collection: client.Database(database).Collection("http-gateway-params-versions"),
	}
}

func (hgpvr *HttpGatewayParamsVersionRepository) getOneById(id string) (*entity.HttpGatewayParamsVersion, error) {
	filter := bson.M{"id": id}
	existingDoc := hgpvr.Collection.FindOne(context.Background(), filter)
	// Check if the document does not exist
	if existingDoc.Err() != nil {
		return nil, existingDoc.Err()
	}

	var result entity.HttpGatewayParamsVersion
	if err := existingDoc.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (hgpvr *HttpGatewayParamsVersionRepository) Save(httpGatewayParamsVersion *entity.HttpGatewayParamsVersion) error {
	// Check if the document already exists based on the ID
	_, err := hgpvr.getOneById(string(httpGatewayParamsVersion.ID))
	if err != nil {
		// Insert new document
		_, err := hgpvr.Collection.InsertOne(context.Background(), bson.M{
			"id":       httpGatewayParamsVersion.ID,
			"versions": httpGatewayParamsVersion.Versions,
		})
		return err
	}

	return ErrorExistingHttpGatewayParamsVersion
}

func (hgpvr *HttpGatewayParamsVersionRepository) Update(httpGatewayParams *entity.HttpGatewayParams) error {
    // Check if the document already exists based on the ID
    existingDoc, err := hgpvr.getOneById(string(httpGatewayParams.ID))
    if err != nil {
        return err
    }

    // Update existing document
    existingDoc.Versions = append(existingDoc.Versions, entity.HttpGatewayParamsData{
        JobParamsID: httpGatewayParams.JobParamsID,
        Params:      httpGatewayParams,
    })
    _, err = hgpvr.Collection.UpdateOne(context.Background(), bson.M{"id": existingDoc.ID}, bson.M{"$set": bson.M{"versions": existingDoc.Versions}})
    return err
}
