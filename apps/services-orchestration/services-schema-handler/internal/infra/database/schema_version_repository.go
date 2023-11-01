package database

import (
	"apps/services-orchestration/services-schema-handler/internal/entity"
	"context"
	"errors"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrorExistingSchema = errors.New("schema already exists")
)

type SchemaVersionRepository struct {
	log        *log.Logger
	Client     *mongo.Client
	Database   string
	Collection *mongo.Collection
}

func NewSchemaVersionRepository(client *mongo.Client, database string) *SchemaVersionRepository {
	return &SchemaVersionRepository{
		log:        log.New(os.Stdout, "[SCHEMA-VERSION-REPOSITORY] ", log.LstdFlags),
		Client:     client,
		Database:   database,
		Collection: client.Database(database).Collection("schemas-versions"),
	}
}

func (cvr *SchemaVersionRepository) getOneById(id string) (*entity.SchemaVersion, error) {
	filter := bson.M{"id": id}
	existingDoc := cvr.Collection.FindOne(context.Background(), filter)
	// Check if the document does not exist
	if existingDoc.Err() != nil {
		return nil, existingDoc.Err()
	}

	var result entity.SchemaVersion
	if err := existingDoc.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (cvr *SchemaVersionRepository) Save(schemaVersion *entity.SchemaVersion) error {
	// Check if the document already exists based on the ID
	_, err := cvr.getOneById(string(schemaVersion.ID))
	if err != nil {
		// Insert new document
		_, err := cvr.Collection.InsertOne(context.Background(), bson.M{
			"id":       schemaVersion.ID,
			"versions": schemaVersion.Versions,
		})
		return err
	}
	return ErrorExistingSchema
}



func (cvr *SchemaVersionRepository) Update(schema *entity.Schema) error {
	existingSchema, err := cvr.getOneById(string(schema.ID))
	if err != nil {
		return err
	}
	schemaData := entity.SchemaData{
		SchemaID: schema.SchemaID,
		Schema:   schema,
	}
	updatedVersions := append(existingSchema.Versions, schemaData)
	filter := bson.M{"id": string(schema.ID)}
	update := bson.M{"$set": bson.M{"versions": updatedVersions}}
	_, err = cvr.Collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (cvr *SchemaVersionRepository) FindAll() ([]*entity.SchemaVersion, error) {
	cursor, err := cvr.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []*entity.SchemaVersion
	for cursor.Next(context.Background()) {
		var result entity.SchemaVersion
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

func (cvr *SchemaVersionRepository) FindOneById(id string) (*entity.SchemaVersion, error) {
	result, err := cvr.getOneById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (cvr *SchemaVersionRepository) FindOneByIdAndVersionId(id string, versionId string) (*entity.Schema, error) {
	schemaVersion, err := cvr.getOneById(id)
	if err != nil {
		return nil, err
	}
	for _, version := range schemaVersion.Versions {
		if string(version.SchemaID) == versionId {
			return version.Schema, nil
		}
	}
	return nil, nil
}

