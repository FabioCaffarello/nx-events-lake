package gogateway

import (
	"context"
	"fmt"
	"net/http"

	inputDTO "libs/golang/services/dtos/services-input-handler/input"
	outputDTO "libs/golang/services/dtos/services-input-handler/output"
	sharedDTO "libs/golang/services/dtos/services-input-handler/shared"
	gorequest "libs/golang/shared/go-request/request"
)

type Client struct {
	ctx     context.Context
	baseURL string
}

func NewClient() *Client {
	return &Client{
		ctx:     context.Background(),
		baseURL: "http://input-handler:8000",
	}
}

func (c *Client) CreateInput(service string, source string, contextEnv string, input inputDTO.InputDTO) (outputDTO.InputDTO, error) {
	url := fmt.Sprintf("%s/inputs/context/%s/service/%s/source/%s", c.baseURL, contextEnv, service, source)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodPost, url, input)
	if err != nil {
		return outputDTO.InputDTO{}, err
	}

	var output outputDTO.InputDTO
	if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &output); err != nil {
		return outputDTO.InputDTO{}, err
	}

	return output, nil
}

func (c *Client) CreateInputWithProcessingID(service string, source string, contextEnv string, input inputDTO.InputDTO, processingID string) (outputDTO.InputDTO, error) {
    url := fmt.Sprintf("%s/inputs/context/%s/service/%s/source/%s/%s", c.baseURL, contextEnv, service, source, processingID)
    req, err := gorequest.CreateRequest(c.ctx, http.MethodPost, url, input)
    if err != nil {
        return outputDTO.InputDTO{}, err
    }

    var output outputDTO.InputDTO
    if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &output); err != nil {
        return outputDTO.InputDTO{}, err
    }

    return output, nil
}

func (c *Client) ListAllInputsByServiceAndSource(service, source string) ([]outputDTO.InputDTO, error) {
	url := fmt.Sprintf("%s/inputs/service/%s/source/%s", c.baseURL, service, source)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var output []outputDTO.InputDTO
	if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &output); err != nil {
		return nil, err
	}

	return output, nil
}

func (c *Client) ListAllInputsByService(service string) ([]outputDTO.InputDTO, error) {
	url := fmt.Sprintf("%s/inputs/service/%s", c.baseURL, service)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var output []outputDTO.InputDTO
	if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &output); err != nil {
		return nil, err
	}

	return output, nil
}

func (c *Client) ListOneInputByIdAndService(id, service, source string) (outputDTO.InputDTO, error) {
	url := fmt.Sprintf("%s/inputs/service/%s/source/%s/%s", c.baseURL, service, source, id)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return outputDTO.InputDTO{}, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var output outputDTO.InputDTO
	if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &output); err != nil {
		return outputDTO.InputDTO{}, err
	}

	return output, nil
}

func (c *Client) ListAllInputsByServiceAndSourceAndStatus(service string, source string, status int) ([]outputDTO.InputDTO, error) {
	url := fmt.Sprintf("%s/inputs/service/%s/source/%s/status/%d", c.baseURL, service, source, status)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var output []outputDTO.InputDTO
	if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &output); err != nil {
		return nil, err
	}

	return output, nil
}

func (c *Client) UpdateInputStatus(status sharedDTO.Status, contextEnv string, service string, source string, id string) (outputDTO.InputDTO, error) {
	url := fmt.Sprintf("%s/inputs/context/%s/service/%s/source/%s/%s", c.baseURL, contextEnv, service, source, id)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodPost, url, status)
	if err != nil {
		return outputDTO.InputDTO{}, err
	}

	var output outputDTO.InputDTO
	if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &output); err != nil {
		return outputDTO.InputDTO{}, err
	}

	return output, nil
}
