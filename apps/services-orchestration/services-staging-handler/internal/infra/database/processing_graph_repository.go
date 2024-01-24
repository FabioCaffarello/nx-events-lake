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
            "context": processingGraph.Context,
            "source": processingGraph.Source,
            "start_processing_id": processingGraph.StartProcessingId,
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

func (pgr *ProcessingGraphRepository) CreateTask(source string, startProcessingId string, task *entity.Task) (*entity.ProcessingGraph, error) {
    processingGraph, err := pgr.FindOneBySourceAndStartProcessingId(source, startProcessingId)
    if err != nil {
        return nil, err
    }
    processingGraph.Tasks = append(processingGraph.Tasks, *task)
    _, err = pgr.Collection.UpdateOne(context.Background(), bson.M{"id": processingGraph.ID}, bson.M{"$set": bson.M{"tasks": processingGraph.Tasks}})
    if err != nil {
        return nil, err
    }
    return processingGraph, nil
}

func (pgr *ProcessingGraphRepository) UpdateTaskStatus(source string, processingId string, statusCode int) (*entity.ProcessingGraph, error) {
    processingGraph, err := pgr.FindOneByTaskSourceAndProcessingId(source, processingId)
    if err != nil {
        return nil, err
    }
    for i, t := range processingGraph.Tasks {
        if t.Source == source && t.ProcessingId == processingId {
            processingGraph.Tasks[i].StatusCode = statusCode
            break
        }
    }
    _, err = pgr.Collection.UpdateOne(context.Background(), bson.M{"id": processingGraph.ID}, bson.M{"$set": bson.M{"tasks": processingGraph.Tasks}})
    if err != nil {
        return nil, err
    }
    return processingGraph, nil
}

func (pgr *ProcessingGraphRepository) UpdateTaskOutput(source string, processingId string, outputId string) (*entity.ProcessingGraph, error) {
    processingGraph, err := pgr.FindOneByTaskSourceAndProcessingId(source, processingId)
    if err != nil {
        return nil, err
    }
    for i, t := range processingGraph.Tasks {
        if t.Source == source && t.ProcessingId == processingId {
            processingGraph.Tasks[i].OutputId = outputId
            break
        }
    }
    _, err = pgr.Collection.UpdateOne(context.Background(), bson.M{"id": processingGraph.ID}, bson.M{"$set": bson.M{"tasks": processingGraph.Tasks}})
    if err != nil {
        return nil, err
    }
    return processingGraph, nil
}

func (pgr *ProcessingGraphRepository) FindOneBySourceAndStartProcessingId(source string, startProcessingId string) (*entity.ProcessingGraph, error) {
    filter := bson.M{"source": source, "start_processing_id": startProcessingId}
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

func (pgr *ProcessingGraphRepository) FindOneByTaskSourceAndProcessingId(source string, processingId string) (*entity.ProcessingGraph, error) {
    filter := bson.M{"tasks.source": source, "tasks.processing_id": processingId}
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
