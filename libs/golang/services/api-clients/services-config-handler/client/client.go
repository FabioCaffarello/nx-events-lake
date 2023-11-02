package client

import (
	"context"
	"fmt"
	inputDTO "libs/golang/services/dtos/services-config-handler/input"
	outputDTO "libs/golang/services/dtos/services-config-handler/output"
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
		baseURL: "http://config-handler:8000",
	}
}

func (c *Client) CreateConfig(configInput inputDTO.ConfigDTO) (outputDTO.ConfigDTO, error) {
	url := fmt.Sprintf("%s/configs", c.baseURL)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodPost, url, configInput)
	if err != nil {
		return outputDTO.ConfigDTO{}, err
	}

	var configOutput outputDTO.ConfigDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &configOutput)
	if err != nil {
		return outputDTO.ConfigDTO{}, err
	}

	return configOutput, nil
}

func (c *Client) ListAllConfigs() ([]outputDTO.ConfigDTO, error) {
	url := fmt.Sprintf("%s/configs", c.baseURL)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var configList []outputDTO.ConfigDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &configList)
	if err != nil {
		return nil, err
	}

	return configList, nil
}

func (c *Client) ListOneConfigById(id string) (outputDTO.ConfigDTO, error) {
	url := fmt.Sprintf("%s/configs/%s", c.baseURL, id)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return outputDTO.ConfigDTO{}, err
	}

	var configOutput outputDTO.ConfigDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &configOutput)
	if err != nil {
		return outputDTO.ConfigDTO{}, err
	}

	return configOutput, nil
}

func (c *Client) ListAllConfigsByService(service string) ([]outputDTO.ConfigDTO, error) {
	url := fmt.Sprintf("%s/configs/service/%s", c.baseURL, service)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var configList []outputDTO.ConfigDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &configList)
	if err != nil {
		return nil, err
	}

	return configList, nil
}

func (c *Client) ListAllConfigsByServiceAndContext(service string, contextEnv string) ([]outputDTO.ConfigDTO, error) {
	url := fmt.Sprintf("%s/configs/service/%s/context/%s", c.baseURL, service, contextEnv)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var configList []outputDTO.ConfigDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &configList)
	if err != nil {
		return nil, err
	}

	return configList, nil
}

func (c *Client) ListAllConfigsByDependentJob(service string, source string) ([]outputDTO.ConfigDTO, error) {
	url := fmt.Sprintf("%s/configs/service/%s/source/%s", c.baseURL, service, source)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var configList []outputDTO.ConfigDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &configList)
	if err != nil {
		return nil, err
	}

	return configList, nil
}
