package database

import (
	"apps/services-orchestration/services-config-handler/internal/entity"
	"context"
	"errors"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrorExistingConfig = errors.New("config already exists")
)

type ConfigVersionRepository struct {
	log        *log.Logger
	Client     *mongo.Client
	Database   string
	Collection *mongo.Collection
}

func NewConfigVersionRepository(client *mongo.Client, database string) *ConfigVersionRepository {
	return &ConfigVersionRepository{
		log:        log.New(os.Stdout, "[CONFIG-VERSION-REPOSITORY] ", log.LstdFlags),
		Client:     client,
		Database:   database,
		Collection: client.Database(database).Collection("configs-versions"),
	}
}

func (cvr *ConfigVersionRepository) getOneById(id string) (*entity.ConfigVersion, error) {
	filter := bson.M{"id": id}
	existingDoc := cvr.Collection.FindOne(context.Background(), filter)
	// Check if the document does not exist
	if existingDoc.Err() != nil {
		return nil, existingDoc.Err()
	}

	var result entity.ConfigVersion
	if err := existingDoc.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (cvr *ConfigVersionRepository) Save(configVersion *entity.ConfigVersion) error {
	// Check if the document already exists based on the ID
	_, err := cvr.getOneById(string(configVersion.ID))
	if err != nil {
		// Insert new document
		_, err := cvr.Collection.InsertOne(context.Background(), bson.M{
			"id":       configVersion.ID,
			"versions": configVersion.Versions,
		})
		return err
	}
	return ErrorExistingConfig
}

func (cvr *ConfigVersionRepository) Update(config *entity.Config) error {
	existingConfig, err := cvr.getOneById(string(config.ID))
	if err != nil {
		return err
	}
	configData := entity.ConfigData{
		ConfigID: config.ConfigID,
		Config:   config,
	}
	updatedVersions := append(existingConfig.Versions, configData)
	filter := bson.M{"id": string(config.ID)}
	update := bson.M{"$set": bson.M{"versions": updatedVersions}}
	_, err = cvr.Collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (cvr *ConfigVersionRepository) FindAll() ([]*entity.ConfigVersion, error) {
	cursor, err := cvr.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []*entity.ConfigVersion
	for cursor.Next(context.Background()) {
		var result entity.ConfigVersion
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func (cvr *ConfigVersionRepository) FindOneById(id string) (*entity.ConfigVersion, error) {
	result, err := cvr.getOneById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (cvr *ConfigVersionRepository) FindOneByIdAndVersionId(id string, versionId string) (*entity.Config, error) {
	configVersion, err := cvr.getOneById(id)
	if err != nil {
		return nil, err
	}
	for _, version := range configVersion.Versions {
		if string(version.ConfigID) == versionId {
			return version.Config, nil
		}
	}
	return nil, nil
}
