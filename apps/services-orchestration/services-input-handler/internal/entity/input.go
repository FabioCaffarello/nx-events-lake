package entity

import (
	"errors"
	"libs/golang/shared/go-id/md5"
	"libs/golang/shared/go-id/uuid"
	"time"
)

var (
    ErrDataFieldIsRequired = errors.New("data field is required")
    ErrSourceFieldIsRequired = errors.New("source field is required")
    ErrServiceFieldIsRequired = errors.New("service field is required")
    ErrContextFieldIsRequired = errors.New("context field is required")
)

type Metadata struct {
	ProcessingId        uuid.ID
	ProcessingTimestamp string
	Source              string
	Service             string
	Context             string
}

type Status struct {
	Code   int    `bson:"code"`
	Detail string `bson:"detail"`
}

type Input struct {
	ID       md5.ID                 `bson:"id"`
	Data     map[string]interface{} `bson:"data"`
	Metadata Metadata               `bson:"metadata"`
	Status   Status                 `bson:"status"`
}


func NewInput(data map[string]interface{}, source string, service string, contextEnv string) (*Input, error) {
	input := &Input{
		ID:   md5.NewWithSourceID(data, source),
		Data: data,
		Metadata: Metadata{
			ProcessingId:        uuid.NewID(),
			ProcessingTimestamp: time.Now().Format(time.RFC3339),
			Source:              source,
			Service:             service,
			Context:             contextEnv,
		},
		Status: Status{
			Code:   0,
			Detail: "",
		},
	}
	err := input.IsValid()
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (i *Input) IsValid() error {
	if i.Data == nil {
		return ErrDataFieldIsRequired
	}
	if i.Metadata.Source == "" {
		return ErrSourceFieldIsRequired
	}
	if i.Metadata.Service == "" {
		return ErrServiceFieldIsRequired
	}
    if i.Metadata.Context == "" {
        return ErrContextFieldIsRequired
    }
	return nil
}
