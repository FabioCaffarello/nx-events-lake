# services-output-handler

The `services-output-handler` API Client is a Go library that allows you to interact with a remote HTTP API for managing output data. This client provides a set of methods to perform various operations on outputs, such as creating new output, listing all outputs, and retrieving specific output by different criteria.

## Getting Started

### Creating a Client

To create a new instance of the `services-output-handler` API Client, use the `NewClient` function. It initializes the client with default configuration:

```go
client := outputClient.NewClient()
```

You can also customize the client's configuration by setting a custom base URL or providing a context.

### Creating an Input

You can create a new output using the `CreateOutput` method. It takes a `ServiceOutputDTO` as input and sends a POST request to the API to create a new output.

```go
outputData := inputDTO.ServiceOutputDTO{
    // Set your output details here
}

output, err := client.CreateOutput(outputData)
if err != nil {
    // Handle the error
}
```

### Listing Outputs
You can retrieve a list of outputs based on various criteria. Here are some examples:

#### List All Outputs by Service and Source
```go
outputs, err := client.ListAllOutputsByServiceAndSource("your-service-name", "your-source-name")
if err != nil {
    // Handle the error
}
```

#### List All Outputs by Service
```go
outputs, err := client.ListAllOutputsByService("your-service-name")
if err is not nil {
    // Handle the error
}
```

#### List One Output by ID, Service, and Source
```go
outputs, err := client.ListOneOutputsByServiceAndId("your-output-id", "your-service-name", "your-source-name")
if err != nil {
    // Handle the error
}
```


## Error Handling

In case of any errors during API requests, the methods return an error value that you can handle according to your application's needs.

## Contributing

Contributions to this library are welcome. If you find issues or have suggestions for improvements, feel free to create an issue or submit a pull request.

Feel free to customize this README to include information specific to your library and its usage.
