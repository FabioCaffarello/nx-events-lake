package client

import (
	"context"
	"fmt"
	outputDTO "libs/golang/services/dtos/services-schema-handler/output"
	gorequest "libs/golang/shared/go-request/request"
	"net/http"
)

type Client struct {
	ctx     context.Context
	baseURL string
}

func NewClient() *Client {
	return &Client{
		ctx:     context.Background(),
		baseURL: "http://schema-handler:8000",
	}
}

func (c *Client) ListOneSchemaByServiceAndSourceAndContextAndSchemaType(service string, source string, context string, schemaType string) (outputDTO.SchemaDTO, error) {
	url := fmt.Sprintf("%s/schemas/service/%s/source/%s/context/%s/schema-type/%s", c.baseURL, service, source, context, schemaType)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return outputDTO.SchemaDTO{}, err
	}

	var schemaOutput outputDTO.SchemaDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &schemaOutput)
	if err != nil {
		return outputDTO.SchemaDTO{}, err
	}

	return schemaOutput, nil
}
