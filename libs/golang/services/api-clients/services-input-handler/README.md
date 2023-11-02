# services-input-handler

The `services-input-handler` API Client is a Go library that allows you to interact with a remote HTTP API for managing input data. This client provides a set of methods to perform various operations on inputs, such as creating new input, updating an input, listing all inputs, and retrieving specific input by different criteria.

## Getting Started

### Creating a Client

To create a new instance of the `services-input-handler` API Client, use the `NewClient` function. It initializes the client with default configuration:

```go
client := inputClient.NewClient()
```

You can also customize the client's configuration by setting a custom base URL or providing a context.

### Creating an Input

You can create a new configuration using the `CreateInput` method. It takes a `InputDTO` as input and sends a POST request to the API to create a new configuration.

```go
inputData := inputDTO.InputDTO{
    // Set your configuration details here
}

input, err := client.CreateConfig(inputData)
if err != nil {
    // Handle the error
}
```

### Listing Inputs
You can retrieve a list of inputs based on various criteria. Here are some examples:

#### List All Inputs by Service and Source
```go
inputs, err := client.ListAllInputsByServiceAndSource("your-service-name", "your-source-name")
if err != nil {
    // Handle the error
}
```

#### List All Inputs by Service
```go
inputs, err := client.ListAllInputsByService("your-service-name")
if err is not nil {
    // Handle the error
}
```

#### List One Input by ID, Service, and Source
```go
input, err := client.ListOneInputByIdAndService("your-input-id", "your-service-name", "your-source-name")
if err != nil {
    // Handle the error
}
```

#### List All Inputs by Service, Source, and Status
```go
inputs, err := client.ListAllInputsByServiceAndSourceAndStatus("your-service-name", "your-source-name", your-status)
if err != nil {
    // Handle the error
}
```

### Updating Input Status
You can update the status of an input using the UpdateInputStatus method. It takes the new status, context environment, service name, source name, and input ID as input and sends a POST request to the API to update the input's status.

```go
newStatus := sharedDTO.Status{
    // Set the new status details here
}

updatedInput, err := client.UpdateInputStatus(newStatus, "your-context-environment", "your-service-name", "your-source-name", "your-input-id")
if err != nil {
    // Handle the error
}
```

## Error Handling

In case of any errors during API requests, the methods return an error value that you can handle according to your application's needs.

## Contributing

Contributions to this library are welcome. If you find issues or have suggestions for improvements, feel free to create an issue or submit a pull request.

Feel free to customize this README to include information specific to your library and its usage.
