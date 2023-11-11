package database

import (
	"apps/services-orchestration/services-file-catalog-handler/internal/entity"
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FileCatalogRepository struct {
	log        *log.Logger
	Client     *mongo.Client
	Database   string
	Collection *mongo.Collection
}

func NewFileCatalogRepository(client *mongo.Client, database string) *FileCatalogRepository {
	return &FileCatalogRepository{
		log:        log.New(os.Stdout, "[SCHEMA-CATALOG-REPOSITORY] ", log.LstdFlags),
		Client:     client,
		Database:   database,
		Collection: client.Database(database).Collection("catalog"),
	}
}

func (scr *FileCatalogRepository) getOneById(id string) (*entity.FileCatalog, error) {
	filter := bson.M{"id": id}
	existingDoc := scr.Collection.FindOne(context.Background(), filter)
	// Check if the document does not exist
	if existingDoc.Err() != nil {
		return nil, existingDoc.Err()
	}

	var result entity.FileCatalog
	if err := existingDoc.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (scr *FileCatalogRepository) Save(fileCatalog *entity.FileCatalog) error {
	// Check if the document already exists based on the ID
	existingFileCatalog, err := scr.getOneById(string(fileCatalog.ID))
	if err != nil {
		// Insert new document
		fileCatalog.CreatedAt = time.Now().Format(time.RFC3339)
		_, err := scr.Collection.InsertOne(context.Background(), bson.M{
			"id":          fileCatalog.ID,
			"service":     fileCatalog.Service,
			"source":      fileCatalog.Source,
			"context":     fileCatalog.Context,
			"lake_layer":  fileCatalog.LakeLayer,
			"schema_type": fileCatalog.SchemaType,
			"catalog_id":  fileCatalog.CatalogID,
			"catalog":     fileCatalog.Catalog,
			"created_at":  fileCatalog.CreatedAt,
			"updated_at":  fileCatalog.UpdatedAt,
		})
		if err != nil {
			return err
		}
	} else {
		// Update existing document
		fileCatalog.UpdatedAt = time.Now().Format(time.RFC3339)
		_, err := scr.Collection.UpdateOne(context.Background(), bson.M{
			"id": existingFileCatalog.ID,
		}, bson.M{
			"$set": bson.M{
				"service":     fileCatalog.Service,
				"source":      fileCatalog.Source,
				"context":     fileCatalog.Context,
				"lake_layer":  fileCatalog.LakeLayer,
				"schema_type": fileCatalog.SchemaType,
				"catalog_id":  fileCatalog.CatalogID,
				"catalog":     fileCatalog.Catalog,
				"updated_at":  fileCatalog.UpdatedAt,
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (scr *FileCatalogRepository) FindOneById(id string) (*entity.FileCatalog, error) {
	result, err := scr.getOneById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (scr *FileCatalogRepository) FindAll() ([]*entity.FileCatalog, error) {
	cursor, err := scr.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []*entity.FileCatalog
	for cursor.Next(context.Background()) {
		var result entity.FileCatalog
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

func (scr *FileCatalogRepository) FindAllByService(service string) ([]*entity.FileCatalog, error) {
	filter := bson.M{"service": service}
	cursor, err := scr.Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []*entity.FileCatalog
	for cursor.Next(context.Background()) {
		var result entity.FileCatalog
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}

	return results, nil
}

func (scr *FileCatalogRepository) FindAllByServiceAndSource(service string, source string) ([]*entity.FileCatalog, error) {
	filter := bson.M{"service": service, "source": source}
	cursor, err := scr.Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []*entity.FileCatalog
	for cursor.Next(context.Background()) {
		var result entity.FileCatalog
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}

	return results, nil
}

func (scr *FileCatalogRepository) DeleteOneById(id string) error {
	filter := bson.M{"id": id}
	_, err := scr.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (scr *FileCatalogRepository) FindOneByServiceAndSource(service string, source string) (*entity.FileCatalog, error) {
	filter := bson.M{"service": service, "source": source}
	existingDoc := scr.Collection.FindOne(context.Background(), filter)
	// Check if the document does not exist
	if existingDoc.Err() != nil {
		return nil, existingDoc.Err()
	}

	var result entity.FileCatalog
	if err := existingDoc.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
