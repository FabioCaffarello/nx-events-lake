package client

import (
	"context"
	"fmt"
	inputDTO "libs/golang/services/dtos/services-staging-handler/input"
	outputDTO "libs/golang/services/dtos/services-staging-handler/output"
	sharedDTO "libs/golang/services/dtos/services-staging-handler/shared"
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
		baseURL: "http://staging-handler:8000",
	}
}

func (c *Client) CreateProcessingJobDependencies(jobInput inputDTO.ProcessingJobDependenciesDTO) (outputDTO.ProcessingJobDependenciesDTO, error) {
	url := fmt.Sprintf("%s/jobs-dependencies", c.baseURL)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodPost, url, jobInput)
	if err != nil {
		return outputDTO.ProcessingJobDependenciesDTO{}, err
	}

	var dependenciesOutput outputDTO.ProcessingJobDependenciesDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &dependenciesOutput)
	if err != nil {
		return outputDTO.ProcessingJobDependenciesDTO{}, err
	}

	return dependenciesOutput, nil
}

func (c *Client) ListOneProcessingJobDependenciesById(id string) (outputDTO.ProcessingJobDependenciesDTO, error) {
	url := fmt.Sprintf("%s/jobs-dependencies/%s", c.baseURL, id)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return outputDTO.ProcessingJobDependenciesDTO{}, err
	}

	var dependenciesOutput outputDTO.ProcessingJobDependenciesDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &dependenciesOutput)
	if err != nil {
		return outputDTO.ProcessingJobDependenciesDTO{}, err
	}

	return dependenciesOutput, nil
}

func (c *Client) RemoveProcessingJobDependencies(id string) (outputDTO.ProcessingJobDependenciesDTO, error) {
	url := fmt.Sprintf("%s/jobs-dependencies/%s", c.baseURL, id)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodDelete, url, nil)
	if err != nil {
		return outputDTO.ProcessingJobDependenciesDTO{}, err
	}

	var dependenciesOutput outputDTO.ProcessingJobDependenciesDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &dependenciesOutput)
	if err != nil {
		return outputDTO.ProcessingJobDependenciesDTO{}, err
	}

	return dependenciesOutput, nil
}

func (c *Client) UpdateProcessingJobDependencies(id string, jobDep sharedDTO.ProcessingJobDependencies) (outputDTO.ProcessingJobDependenciesDTO, error) {
	url := fmt.Sprintf("%s/jobs-dependencies/%s", c.baseURL, id)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodPost, url, jobDep)
	if err != nil {
		return outputDTO.ProcessingJobDependenciesDTO{}, err
	}

	var dependenciesOutput outputDTO.ProcessingJobDependenciesDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &dependenciesOutput)
	if err != nil {
		return outputDTO.ProcessingJobDependenciesDTO{}, err
	}

	return dependenciesOutput, nil
}
