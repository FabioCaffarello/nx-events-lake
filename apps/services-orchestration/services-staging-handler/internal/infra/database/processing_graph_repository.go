package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

    "apps/services-orchestration/services-staging-handler/internal/entity"
)

type ProcessingGraphRepository struct {
	log        *log.Logger
	Client     *mongo.Client
	Database   string
	Collection *mongo.Collection
}

func NewProcessingGaphRepository(client *mongo.Client, database string) *ProcessingGraphRepository {
	return &ProcessingGraphRepository{
		log:        log.New(os.Stdout, "[PROCESSING-GRAPH-REPOSITORY] ", log.LstdFlags),
		Client:     client,
		Database:   database,
		Collection: client.Database(database).Collection("processing-graph"),
	}
}

func (pgr *ProcessingGraphRepository) getOneById(id string) (*entity.ProcessingGraph, error) {
	filter := bson.M{"id": id}
	existingDoc := pgr.Collection.FindOne(context.Background(), filter)
	// Check if the document does not exist
	if existingDoc.Err() != nil {
		return nil, existingDoc.Err()
	}

	var result entity.ProcessingGraph
	if err := existingDoc.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (pgr *ProcessingGraphRepository) Save(processingGraph *entity.ProcessingGraph) error {
    _, err := pgr.getOneById(string(processingGraph.ID))
    if err != nil {
        // Insert new document
        _, err := pgr.Collection.InsertOne(context.Background(), bson.M{
            "id": processingGraph.ID,
            "frequency": processingGraph.Frequency,
            "context": processingGraph.Context,
            "source": processingGraph.Source,
            "service_start": processingGraph.ServiceStart,
            "processing_id": processingGraph.ProcessingId,
            "tasks": processingGraph.Tasks,
            "created_at": processingGraph.CreatedAt,
            "updated_at": processingGraph.UpdatedAt,
        })
        if err != nil {
            return err
        }
        return nil
    }
    return nil
}

// func (pgr *ProcessingGraphRepository) UpdateProcessingGraphTask(processingGraph *entity.Task, id string) error {
//     existingProcessingGraph, err := pgr.getOneById(id)
// }

func (pgr *ProcessingGraphRepository) FindOneById(id string) (*entity.ProcessingGraph, error) {
    return pgr.getOneById(id)
}

func (pgr *ProcessingGraphRepository) Delete(id string) error {
    filter := bson.M{"id": id}
    _, err := pgr.Collection.DeleteOne(context.Background(), filter)
    if err != nil {
        return err
    }
    return nil
}
