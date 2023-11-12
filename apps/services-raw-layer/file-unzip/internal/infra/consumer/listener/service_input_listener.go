package listener

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"apps/services-raw-layer/file-unzip/configs"
	"apps/services-raw-layer/file-unzip/internal/usecase"
	"libs/golang/resources/go-rabbitmq/queue"
	eventsHandlerInputDTO "libs/golang/services/dtos/services-events-handler/input"
	eventsHandlerSharedDTO "libs/golang/services/dtos/services-events-handler/shared"
	inputHandlerOutputDTO "libs/golang/services/dtos/services-input-handler/output"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	ErrorInvalidServiceInputDTO        = errors.New("invalid service input message")
	ErrorCouldNotNotifyServiceFeedback = errors.New("could not notify service feedback")
	statusCodeOK                       = 200
	StatusCodeError                    = 500
	statusDetailOK                     = "OK"
	StatusDetailError                  = "Could not unzip file"
)

type ServiceInputListener struct {
	ContextEnv     string
	ServiceName    string
	MinioEndpoint  string
	MinioAccessKey string
	MinioSecretKey string
}

func NewServiceInputListener(config configs.Config) *ServiceInputListener {
	return &ServiceInputListener{
		ContextEnv:     config.ContextEnv,
		ServiceName:    config.ServiceName,
		MinioEndpoint:  config.MinioEndpoint,
		MinioAccessKey: config.MinioAccessKey,
		MinioSecretKey: config.MinioSecretKey,
	}
}

func (l *ServiceInputListener) Handle(rabbitMQ *queue.RabbitMQ, exchange string, msg amqp.Delivery) error {
	var serviceInputDTO inputHandlerOutputDTO.InputDTO
	log.Printf("ServiceInputListener.Handle: msg.Body=%s", msg.Body)
	err := json.Unmarshal(msg.Body, &serviceInputDTO)
	if err != nil {
		return ErrorInvalidServiceInputDTO
	}

    err = DispatchInput(rabbitMQ, exchange, msg.Body)
    if err != nil {
        return ErrorCouldNotNotifyServiceFeedback
    }

    time.Sleep(5 * time.Second)

	source := serviceInputDTO.Metadata.Source
	uri := serviceInputDTO.Data["documentUri"].(string)
	partition := serviceInputDTO.Data["partition"].(string)
	log.Printf("ServiceInputListener.Handle: uri=%s, partition=%s, source=%s", uri, partition, source)

	unzipFileUseCase := usecase.NewUnzipFileUseCase(l.ContextEnv, l.MinioEndpoint, l.MinioAccessKey, l.MinioSecretKey)

	uris, err := unzipFileUseCase.Execute(uri, partition, source)
	var serviceFeedbackDTO eventsHandlerInputDTO.ServiceFeedbackDTO
	if err != nil {
		log.Printf("ServiceInputListener.Handle: err=%s", err)
		uriResult := ""
		serviceFeedbackDTO = l.getServiceOutputDTO(uri, uriResult, partition, serviceInputDTO, StatusCodeError, StatusDetailError)
		jsonOutput, _ := json.Marshal(serviceFeedbackDTO)
		err = DispatchOutput(rabbitMQ, exchange, jsonOutput)
		if err != nil {
			return ErrorCouldNotNotifyServiceFeedback
		}

	} else {
		log.Printf("ServiceInputListener.Handle: uris=%s", uris)
		for _, uriResult := range uris {
			log.Printf("ServiceInputListener.Handle: uriResult=%s", uriResult)
			serviceFeedbackDTO = l.getServiceOutputDTO(uri, uriResult, partition, serviceInputDTO, statusCodeOK, statusDetailOK)
			jsonOutput, _ := json.Marshal(serviceFeedbackDTO)
			log.Printf("ServiceInputListener.Handle: jsonOutput=%s", jsonOutput)
			err = DispatchOutput(rabbitMQ, exchange, jsonOutput)
			if err != nil {
				return ErrorCouldNotNotifyServiceFeedback
			}
		}
	}

	return nil
}

func DispatchInput(rabbitMQ *queue.RabbitMQ, exchange string, jsonOutput []byte) error {
	err := rabbitMQ.Notify(
		jsonOutput,
		"application/json",
		exchange,
		"input-processing",
	)
	if err != nil {
		return err
	}
	return nil
}

func DispatchOutput(rabbitMQ *queue.RabbitMQ, exchange string, jsonOutput []byte) error {
	err := rabbitMQ.Notify(
		jsonOutput,
		"application/json",
		exchange,
		"feedback",
	)
	if err != nil {
		return err
	}
	return nil
}

func (l *ServiceInputListener) getServiceOutputDTO(uriOrigin string, uriResult string, partition string, serviceInputDTO inputHandlerOutputDTO.InputDTO, statusCode int, StatusDetail string) eventsHandlerInputDTO.ServiceFeedbackDTO {
	return eventsHandlerInputDTO.ServiceFeedbackDTO{
		Data: map[string]interface{}{
			"documentUri": uriResult,
			"partition":   partition,
		},
		Metadata: l.getJobMetadataDTO(uriOrigin, serviceInputDTO),
		Status: eventsHandlerSharedDTO.Status{
			Code:   statusCode,
			Detail: StatusDetail,
		},
	}
}

func (l *ServiceInputListener) getJobMetadataDTO(uriOrigin string, serviceInputDTO inputHandlerOutputDTO.InputDTO) eventsHandlerSharedDTO.Metadata {
	return eventsHandlerSharedDTO.Metadata{
		Input: eventsHandlerSharedDTO.MetadataInput{
			ID:                  serviceInputDTO.ID,
			Data:                serviceInputDTO.Data,
			ProcessingId:        serviceInputDTO.Metadata.ProcessingId,
			ProcessingTimestamp: serviceInputDTO.Metadata.ProcessingTimestamp,
		},
		Context:             l.ContextEnv,
		Service:             serviceInputDTO.Metadata.Service,
		Source:              serviceInputDTO.Metadata.Source,
		ProcessingTimestamp: time.Now().Format(time.RFC3339),
	}
}
