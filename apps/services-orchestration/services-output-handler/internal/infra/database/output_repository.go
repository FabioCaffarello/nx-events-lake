package database

import (
     "apps/services-orchestration/services-output-handler/internal/entity"
     "context"
     "log"
     "os"

     "go.mongodb.org/mongo-driver/bson"
     "go.mongodb.org/mongo-driver/mongo"
)

type ServiceOutputRepository struct {
     log      *log.Logger
     Client   *mongo.Client
     Database string
}

func NewServiceOutputRepository(client *mongo.Client, database string) *ServiceOutputRepository {
     return &ServiceOutputRepository{
          log:      log.New(os.Stdout, "[SERVICE-OUTPUT-REPOSITORY] ", log.LstdFlags),
          Client:   client,
          Database: database,
     }
}

func (sor *ServiceOutputRepository) getOneById(id string, service string) (*entity.ServiceOutput, error) {
     collection := sor.Client.Database(sor.Database).Collection(service)

     filter := bson.M{"id": id}
     existingDoc := collection.FindOne(context.Background(), filter)
     // Check if the document does not exist
     if existingDoc.Err() != nil {
          return nil, existingDoc.Err()
     }

     var result entity.ServiceOutput
     if err := existingDoc.Decode(&result); err != nil {
          return nil, err
     }

     return &result, nil
}

func (sor *ServiceOutputRepository) FindOneByIdAndService(id string, service string) (*entity.ServiceOutput, error) {
     result, err := sor.getOneById(id, service)
     if err != nil {
          return nil, err
     }
     return result, nil
}

func (sor *ServiceOutputRepository) Save(serviceOutput *entity.ServiceOutput, service string) error {
     collection := sor.Client.Database(sor.Database).Collection(service)
     // Check if the document already exists based on the ID
     _, err := sor.getOneById(string(serviceOutput.ID), service)
     if err != nil {
          // Insert new document
          _, err := collection.InsertOne(context.Background(), bson.M{
               "id":       serviceOutput.ID,
               "data":     serviceOutput.Data,
               "service":  serviceOutput.Service,
               "source":   serviceOutput.Source,
               "context":  serviceOutput.Context,
               "metadata": serviceOutput.Metadata,
               "created_at": serviceOutput.CreatedAt,
               "updated_at": serviceOutput.UpdatedAt,
          })
          if err != nil {
               return err
          }
          return nil
     }
     // Update existing document
     _, err = collection.UpdateOne(context.Background(), bson.M{"id": serviceOutput.ID}, bson.M{"$set": bson.M{
          "data":     serviceOutput.Data,
          "service":  serviceOutput.Service,
          "source":   serviceOutput.Source,
          "context":  serviceOutput.Context,
          "metadata": serviceOutput.Metadata,
          "updated_at": serviceOutput.UpdatedAt,
     }})
     if err != nil {
          return err
     }
     return nil
}

func (sor *ServiceOutputRepository) FindAllByService(service string) ([]*entity.ServiceOutput, error) {
     collection := sor.Client.Database(sor.Database).Collection(service)

     cursor, err := collection.Find(context.Background(), bson.M{})
     if err != nil {
          panic(err)
     }
     defer cursor.Close(context.Background())

     var results []*entity.ServiceOutput
     for cursor.Next(context.Background()) {
          var result *entity.ServiceOutput
          if err := cursor.Decode(&result); err != nil {
               return nil, err
          }
          results = append(results, result)
     }
     if err := cursor.Err(); err != nil {
          return nil, err
     }
     return results, nil
}

func (sor *ServiceOutputRepository) FindAllByServiceAndSource(service string, source string) ([]*entity.ServiceOutput, error) {
     collection := sor.Client.Database(sor.Database).Collection(service)

     filter := bson.M{"source": source}
     cursor, err := collection.Find(context.Background(), filter)
     if err != nil {
          panic(err)
     }
     defer cursor.Close(context.Background())

     var results []*entity.ServiceOutput
     for cursor.Next(context.Background()) {
          var result *entity.ServiceOutput
          if err := cursor.Decode(&result); err != nil {
               return nil, err
          }
          results = append(results, result)
     }
     if err := cursor.Err(); err != nil {
          return nil, err
     }
     return results, nil
}

func (sor *ServiceOutputRepository) FindAllByServiceAndSourceAndContext(service string, source string, contextEnv string) ([]*entity.ServiceOutput, error) {
        collection := sor.Client.Database(sor.Database).Collection(service)
        filter := bson.M{"source": source, "context": contextEnv}
        cursor, err := collection.Find(context.Background(), filter)
        if err != nil {
            panic(err)
        }
        defer cursor.Close(context.Background())

        var results []*entity.ServiceOutput
        for cursor.Next(context.Background()) {
            var result *entity.ServiceOutput
            if err := cursor.Decode(&result); err != nil {
                return nil, err
            }
            results = append(results, result)
        }
        if err := cursor.Err(); err != nil {
            return nil, err
        }
        return results, nil
}

