package gooutputs

import (
	"context"
	"fmt"
	"net/http"

	inputDTO "libs/golang/services/dtos/services-output-handler/input"
	outputDTO "libs/golang/services/dtos/services-output-handler/output"
	gorequest "libs/golang/shared/go-request/request"
)

type Client struct {
	ctx     context.Context
	baseURL string
}

func NewClient() *Client {
	return &Client{
		ctx:     context.Background(),
		baseURL: "http://output-handler:8000",
	}
}

func (c *Client) CreateOutput(service string, serviceOutput inputDTO.ServiceOutputDTO) (outputDTO.ServiceOutputDTO, error) {
	url := fmt.Sprintf("%s/outputs/service/%s", c.baseURL, service)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodPost, url, serviceOutput)
	if err != nil {
		return outputDTO.ServiceOutputDTO{}, err
	}

	var output outputDTO.ServiceOutputDTO
	if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &output); err != nil {
		return outputDTO.ServiceOutputDTO{}, err
	}

	return output, nil
}

func (c *Client) ListOneOutputsByServiceAndId(service string, id string) (outputDTO.ServiceOutputDTO, error) {
	url := fmt.Sprintf("%s/outputs/service/%s/%s", c.baseURL, service, id)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return outputDTO.ServiceOutputDTO{}, err
	}

	var output outputDTO.ServiceOutputDTO
	if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &output); err != nil {
		return outputDTO.ServiceOutputDTO{}, err
	}

	return output, nil
}

func (c *Client) ListAllOutputsByService(service string) ([]outputDTO.ServiceOutputDTO, error) {
	url := fmt.Sprintf("%s/outputs/service/%s", c.baseURL, service)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var outputs []outputDTO.ServiceOutputDTO
	if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &outputs); err != nil {
		return nil, err
	}

	return outputs, nil
}

func (c *Client) ListAllOutputsByServiceAndSource(service string, source string) ([]outputDTO.ServiceOutputDTO, error) {
	url := fmt.Sprintf("%s/outputs/service/%s/source/%s", c.baseURL, service, source)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var outputs []outputDTO.ServiceOutputDTO
	if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &outputs); err != nil {
		return nil, err
	}

	return outputs, nil
}
