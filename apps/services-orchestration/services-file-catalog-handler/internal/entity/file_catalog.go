package entity

import (
	"errors"
	"libs/golang/shared/go-id/config"
	uuid "libs/golang/shared/go-id/schema/uuid"
	"time"
)

var (
	ErrFileCatalogServiceEmpty    = errors.New("file catalog service is empty")
	ErrFileCatalogSourceEmpty     = errors.New("file catalog source is empty")
	ErrFileCatalogContextEmpty    = errors.New("file catalog context is empty")
	ErrFileCatalogEmpty           = errors.New("file catalog is empty")
	ErrFileCatalogLakeLayerEmpty  = errors.New("file catalog lake layer is empty")
	ErrFileCatalogSchemaTypeEmpty = errors.New("file catalog schema type is empty")
)

type FileCatalog struct {
	ID         config.ID              `bson:"id"`
	Service    string                 `bson:"service"`
	Source     string                 `bson:"source"`
	Context    string                 `bson:"context"`
	LakeLayer  string                 `bson:"lake_layer"`
	SchemaType string                 `bson:"schema_type"`
	CatalogID  uuid.ID                `bson:"catalog_id"`
	Catalog    map[string]interface{} `bson:"catalog"`
	CreatedAt  string                 `bson:"created_at"`
	UpdatedAt  string                 `bson:"updated_at"`
}

func NewFileCatalog(
	service string,
	source string,
	context string,
	lakeLayer string,
	schemaType string,
	catalog map[string]interface{},
) (*FileCatalog, error) {
	catalogId, err := uuid.GenerateSchemaID(schemaType, catalog)
	if err != nil {
		return nil, err
	}

	fileCatalog := &FileCatalog{
		ID:         config.NewID(service, source),
		Service:    service,
		Source:     source,
		Context:    context,
		LakeLayer:  lakeLayer,
		SchemaType: schemaType,
		CatalogID:  catalogId,
		Catalog:    catalog,
		CreatedAt:  time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}
	err = fileCatalog.IsFileCatalogValid()
	if err != nil {
		return nil, err
	}
	return fileCatalog, nil
}

func (fileCatalog *FileCatalog) IsFileCatalogValid() error {
	if fileCatalog.Service == "" {
		return ErrFileCatalogServiceEmpty
	}
	if fileCatalog.Source == "" {
		return ErrFileCatalogSourceEmpty
	}
	if fileCatalog.Context == "" {
		return ErrFileCatalogContextEmpty
	}
	if len(fileCatalog.Catalog) == 0 {
		return ErrFileCatalogEmpty
	}
	if fileCatalog.LakeLayer == "" {
		return ErrFileCatalogLakeLayerEmpty
	}
	if fileCatalog.SchemaType == "" {
		return ErrFileCatalogSchemaTypeEmpty
	}
	return nil
}
