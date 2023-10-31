# go-request

The `go-request` library provides utilities for making HTTP requests and handling responses in your Go applications. It simplifies the process of creating HTTP requests, sending them, and handling the responses. It is designed to make interacting with HTTP services easier and more convenient.

## Usage

### Creating Requests
The CreateRequest function allows you to create HTTP requests with ease. It takes a context, HTTP method, URL, and an optional request body. Here's how to use it:

```golang
import gorequest "libs/golang/shared/go-request/request"

ctx := context.Background()
method := "POST"
url := "https://example.com"
body := map[string]string{"key": "value"}

req, err := gorequest.CreateRequest(ctx, method, url, body)
if err != nil {
    log.Fatalf("CreateRequest() failed: %v", err)
}
```

### Sending Requests

The `SendRequest` function simplifies sending HTTP requests and handling responses. It takes an HTTP request, an HTTP client (you can use the provided `DefaultHTTPClient` or create your own), and a result structure for the response. Here's how to use it:

```golang
import gorequest "libs/golang/shared/go-request/request"

ctx := context.Background()
method := "GET"
url := "https://example.com"
var result map[string]interface{}

req, err := gorequest.CreateRequest(ctx, method, url, nil)
if err != nil {
    log.Fatalf("CreateRequest() failed: %v", err)
}

client := gorequest.DefaultHTTPClient

err = gorequest.SendRequest(req, client, &result)
if err != nil {
    log.Fatalf("SendRequest() failed: %v", err)
}
```
