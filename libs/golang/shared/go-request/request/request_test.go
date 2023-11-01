package gorequest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RequestTestSuite struct {
	suite.Suite
}

func TestRequestSuite(t *testing.T) {
	suite.Run(t, new(RequestTestSuite))
}

func (suite *RequestTestSuite) TestCreateRequest() {
	ctx := context.Background()
	method := "POST"
	url := "https://example.com"
	body := map[string]string{"key": "value"}

	req, err := CreateRequest(ctx, method, url, body)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), method, req.Method)
	assert.Equal(suite.T(), url, req.URL.String())

	assert.Equal(suite.T(), "application/json", req.Header.Get("Content-Type"))
	assert.Equal(suite.T(), "value", body["key"])

}

func (suite *RequestTestSuite) TestSendRequest() {
	// Create a mock server to handle requests
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// Respond with a JSON payload
		w.Write([]byte(`{"response": "success"}`))
	}))
	defer server.Close()

	ctx := context.Background()
	method := "GET"
	url := server.URL
	var result map[string]string

	req, err := CreateRequest(ctx, method, url, nil)
	assert.NoError(suite.T(), err)

	client := DefaultHTTPClient

	err = SendRequest(req, client, &result)
	assert.NoError(suite.T(), err)

	expectedResponse := map[string]string{"response": "success"}
	assert.Equal(suite.T(), expectedResponse, result)

	assert.Equal(suite.T(), method, req.Method)
	assert.Equal(suite.T(), url, req.URL.String())
}
