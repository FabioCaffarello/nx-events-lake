package entity

import (
	"errors"
	"libs/golang/shared/go-id/md5"
	"time"
)

var (
	ErrServiceOutputDataEmpty                     = errors.New("service output data is empty")
	ErrServiceOutputInputIDEmpty                  = errors.New("service output input id is empty")
	ErrServiceOutputInputDataEmpty                = errors.New("service output input data is empty")
	ErrServiceOutputInputProcessingIdEmpty        = errors.New("service output input processing id is empty")
	ErrServiceOutputInputProcessingTimestampEmpty = errors.New("service output input processing timestamp is empty")
	ErrServiceOutputMetadataServiceEmpty          = errors.New("service output metadata service is empty")
	ErrServiceOutputMetadataSourceEmpty           = errors.New("service output metadata source is empty")
	ErrServiceOutputContextEmpty                  = errors.New("service output context is empty")
)

type MetadataInput struct {
	ID                  string                 `bson:"id"`
	Data                map[string]interface{} `bson:"data"`
	ProcessingId        string                 `bson:"processing_id"`
	ProcessingTimestamp string                 `bson:"processing_timestamp"`
}

type Metadata struct {
	InputID string        `bson:"input_id"`
	Input   MetadataInput `bson:"input"`
	Service string        `bson:"service"`
	Source  string        `bson:"source"`
}

type ServiceOutput struct {
	ID        md5.ID                 `bson:"id"`
	Data      map[string]interface{} `bson:"data"`
	Service   string                 `bson:"service"`
	Source    string                 `bson:"source"`
	Context   string                 `bson:"context"`
	Metadata  Metadata               `bson:"metadata"`
	CreatedAt string                 `bson:"created_at"`
	UpdatedAt string                 `bson:"updated_at"`
}

func NewServiceOutput(
	data map[string]interface{},
	inputId string,
	inputData map[string]interface{},
	processingId string,
	processingTimestamp string,
	service string,
	source string,
	context string,
) (*ServiceOutput, error) {
	serviceOutput := &ServiceOutput{
		ID:   md5.NewWithSourceAndServiceID(data, source, service),
		Data: data,
		Metadata: Metadata{
			InputID: inputId,
			Input: MetadataInput{
				ID:                  inputId,
				Data:                inputData,
				ProcessingId:        processingId,
				ProcessingTimestamp: processingTimestamp,
			},
			Service: service,
			Source:  source,
		},
		Service:   service,
		Source:    source,
		Context:   context,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
	err := serviceOutput.IsServiceOutputValid()
	if err != nil {
		return nil, err
	}
	return serviceOutput, nil
}

func (s *ServiceOutput) IsServiceOutputValid() error {
	if s.Data == nil {
		return ErrServiceOutputDataEmpty
	}
	if s.Metadata.InputID == "" {
		return ErrServiceOutputInputIDEmpty
	}
	if s.Metadata.Input.Data == nil {
		return ErrServiceOutputInputDataEmpty
	}
	if s.Metadata.Input.ProcessingId == "" {
		return ErrServiceOutputInputProcessingIdEmpty
	}
	if s.Metadata.Input.ProcessingTimestamp == "" {
		return ErrServiceOutputInputProcessingTimestampEmpty
	}
	if s.Metadata.Service == "" {
		return ErrServiceOutputMetadataServiceEmpty
	}
	if s.Metadata.Source == "" {
		return ErrServiceOutputMetadataSourceEmpty
	}
	if s.Context == "" {
		return ErrServiceOutputContextEmpty
	}
	return nil
}
