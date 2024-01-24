package database

import (
	"context"
	"log"
	"os"

	"apps/services-orchestration/services-jobs-handler/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type HttpGatewayParamsRepository struct {
	log        *log.Logger
	Client     *mongo.Client
	Database   string
	Collection *mongo.Collection
}

func NewHttpGatewayParamsRepository(client *mongo.Client, database string) *HttpGatewayParamsRepository {
	return &HttpGatewayParamsRepository{
		log:        log.New(os.Stdout, "[HTTP-GATEWAY-PARAMS-REPOSITORY] ", log.LstdFlags),
		Client:     client,
		Database:   database,
		Collection: client.Database(database).Collection("http-gateway-params"),
	}
}

func (cr *HttpGatewayParamsRepository) getOneById(id string) (*entity.HttpGatewayParams, error) {
	filter := bson.M{"id": id}
	existingDoc := cr.Collection.FindOne(context.Background(), filter)
	// Check if the document does not exist
	if existingDoc.Err() != nil {
		return nil, existingDoc.Err()
	}

	var result entity.HttpGatewayParams
	if err := existingDoc.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (cr *HttpGatewayParamsRepository) Save(httpGatewayParams *entity.HttpGatewayParams) error {
	// Check if the document already exists based on the ID
	existingHttpGatewayParams, err := cr.getOneById(string(httpGatewayParams.ID))
	if err != nil {
		// Insert new document
		_, err := cr.Collection.InsertOne(context.Background(), bson.M{
			"id":              httpGatewayParams.ID,
            "service":         httpGatewayParams.Service,
            "source":          httpGatewayParams.Source,
            "context":         httpGatewayParams.Context,
			"base_url":        httpGatewayParams.BaseUrl,
			"url_domains":     httpGatewayParams.UrlDomains,
			"headers":         httpGatewayParams.Headers,
            "job_params_id":   httpGatewayParams.JobParamsID,
			"enable_proxy":    httpGatewayParams.EnableProxy,
			"proxy_loaders":   httpGatewayParams.ProxyLoaders,
			"enable_captcha":  httpGatewayParams.EnableCaptcha,
			"captcha_solvers": httpGatewayParams.CaptchaSolvers,
			"created_at":      httpGatewayParams.CreatedAt,
			"updated_at":      httpGatewayParams.UpdatedAt,
		})
		return err
	}

	// Update existing document
	_, err = cr.Collection.UpdateOne(context.Background(), bson.M{"id": existingHttpGatewayParams.ID}, bson.M{
		"$set": bson.M{
            "service":         httpGatewayParams.Service,
            "source":          httpGatewayParams.Source,
            "context":         httpGatewayParams.Context,
			"base_url":        httpGatewayParams.BaseUrl,
			"url_domains":     httpGatewayParams.UrlDomains,
			"headers":         httpGatewayParams.Headers,
			"enable_proxy":    httpGatewayParams.EnableProxy,
            "job_params_id":   httpGatewayParams.JobParamsID,
			"proxy_loaders":   httpGatewayParams.ProxyLoaders,
			"enable_captcha":  httpGatewayParams.EnableCaptcha,
			"captcha_solvers": httpGatewayParams.CaptchaSolvers,
			"updated_at":      httpGatewayParams.UpdatedAt,
		},
	})
	return err
}

func (cr *HttpGatewayParamsRepository) FindAll() ([]*entity.HttpGatewayParams, error) {
	var httpGatewayParams []*entity.HttpGatewayParams
	cursor, err := cr.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var httpGatewayParam entity.HttpGatewayParams
		if err = cursor.Decode(&httpGatewayParam); err != nil {
			return nil, err
		}
		httpGatewayParams = append(httpGatewayParams, &httpGatewayParam)
	}
	return httpGatewayParams, nil
}

func (cr *HttpGatewayParamsRepository) FindAllBySource(source string) ([]*entity.HttpGatewayParams, error) {
    var httpGatewayParams []*entity.HttpGatewayParams
    cursor, err := cr.Collection.Find(context.Background(), bson.M{"source": source})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())
    for cursor.Next(context.Background()) {
        var httpGatewayParam entity.HttpGatewayParams
        if err = cursor.Decode(&httpGatewayParam); err != nil {
            return nil, err
        }
        httpGatewayParams = append(httpGatewayParams, &httpGatewayParam)
    }
    return httpGatewayParams, nil
}

func (cr *HttpGatewayParamsRepository) FindAllByService(service string) ([]*entity.HttpGatewayParams, error) {
    var httpGatewayParams []*entity.HttpGatewayParams
    cursor, err := cr.Collection.Find(context.Background(), bson.M{"service": service})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())
    for cursor.Next(context.Background()) {
        var httpGatewayParam entity.HttpGatewayParams
        if err = cursor.Decode(&httpGatewayParam); err != nil {
            return nil, err
        }
        httpGatewayParams = append(httpGatewayParams, &httpGatewayParam)
    }
    return httpGatewayParams, nil
}

func (cr *HttpGatewayParamsRepository) FindOneByServiceAndSourceAndContext(service string, source string, contextEnv string) (*entity.HttpGatewayParams, error) {
    filter := bson.M{
        "service": service,
        "source": source,
        "context": contextEnv,
    }
    existingDoc := cr.Collection.FindOne(context.Background(), filter)
    // Check if the document does not exist
    if existingDoc.Err() != nil {
        log.Println("1 Error: ", existingDoc.Err())
        return nil, existingDoc.Err()
    }

    var result entity.HttpGatewayParams
    if err := existingDoc.Decode(&result); err != nil {
        log.Println("2 Error: ", err)
        return nil, err
    }

    return &result, nil
}


func (cr *HttpGatewayParamsRepository) FindOneById(id string) (*entity.HttpGatewayParams, error) {
	return cr.getOneById(id)
}

func (cr *HttpGatewayParamsRepository) Delete(id string) error {
	_, err := cr.Collection.DeleteOne(context.Background(), bson.M{"id": id})
	return err
}
