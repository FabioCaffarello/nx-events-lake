package entity

import (
	"errors"
	"fmt"
	"libs/golang/shared/go-id/md5"
	"time"
)

var (
	ErrContextEmpty           = errors.New("pipeline context is empty")
	ErrSourceEmpty            = errors.New("pipeline source is empty")
	ErrStartProcessingIdEmpty = errors.New("start processing id is empty")
)

type Task struct {
	Source                string `bson:"source"`
	Service               string `bson:"service"`
	ProcessingId          string `bson:"processing_id"`
	ParentProcessingId    string `bson:"parent_processing_id"`
	ConfigVersionId       string `bson:"config_version_id"`
	InputSchemaVersionId  string `bson:"input_schema_version_id"`
	OutputSchemaVersionId string `bson:"output_schema_version_id"`
	InputId               string `bson:"input_id"`
	OutputId              string `bson:"output_id"`
	StatusCode            int    `bson:"status_code"`
	ProcessingTimestamp   string `bson:"processing_timestamp"`
}

type ProcessingGraph struct {
	ID                md5.ID `bson:"id"`
	Context           string `bson:"context"`
	Source            string `bson:"source"`
	StartProcessingId string `bson:"start_processing_id"`
	Tasks             []Task `bson:"tasks"`
	CreatedAt         string `bson:"created_at"`
	UpdatedAt         string `bson:"updated_at"`
}

func NewProcessingGraph(
	context string,
	source string,
	startProcessingId string,
) (*ProcessingGraph, error) {
	pipeline := &ProcessingGraph{
		ID:                md5.NewMd5Hash(fmt.Sprintf("%s-%s-%s", context, source, startProcessingId)),
		Context:           context,
		Source:            source,
		StartProcessingId: startProcessingId,
		CreatedAt:         time.Now().Format(time.RFC3339),
		UpdatedAt:         time.Now().Format(time.RFC3339),
	}
	err := pipeline.IsProcessingGraphValid()
	if err != nil {
		return nil, err
	}
	return pipeline, nil
}

func (p *ProcessingGraph) IsProcessingGraphValid() error {
	if p.Context == "" {
		return ErrContextEmpty
	}
	if p.Source == "" {
		return ErrSourceEmpty
	}
	if p.StartProcessingId == "" {
		return ErrStartProcessingIdEmpty
	}
	return nil
}
