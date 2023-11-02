# services-config-handler

The `services-config-handler` API Client is a Go library that allows you to interact with a remote HTTP API for managing configuration data. This client provides a set of methods to perform various operations on configurations, such as creating new configurations, listing all configurations, and retrieving specific configurations by different criteria.

## Getting Started

### Creating a Client

To create a new instance of the `services-config-handler` API Client, use the `NewClient` function. It initializes the client with default configuration:

```go
client := configClient.NewClient()
```

You can also customize the client's configuration by setting a custom base URL or providing a context.

### Creating a Configuration

You can create a new configuration using the `CreateConfig` method. It takes a `ConfigDTO` as input and sends a POST request to the API to create a new configuration.

```go
configData := inputDTO.ConfigDTO{
    // Set your configuration details here
}

config, err := client.CreateConfig(configData)
if err != nil {
    // Handle the error
}
```

### Listing All Configurations

You can retrieve a list of all configurations using the `ListAllConfigs` method. It sends a GET request to the API to fetch all available configurations.

```go
configList, err := client.ListAllConfigs()
if err != nil {
    // Handle the error
}
```

### Retrieving a Configuration by ID

To retrieve a specific configuration by its ID, you can use the `ListOneConfigById` method. Provide the ID as a parameter, and it will send a GET request to the API to fetch the configuration.

```go
configID := "your-config-id"
config, err := client.ListOneConfigById(configID)
if err != nil {
    // Handle the error
}
```

### Listing Configurations by Service

You can retrieve configurations that belong to a specific service using the `ListAllConfigsByService` method. Provide the service name as a parameter to get the configurations associated with that service.

```go
serviceName := "your-service-name"
configList, err := client.ListAllConfigsByService(serviceName)
if err != nil {
    // Handle the error
}
```

### Listing Configurations by Service and Context

To retrieve configurations based on a specific service and context, use the `ListAllConfigsByServiceAndContext` method. Provide the service and context as parameters to get the configurations that match these criteria.

```go
serviceName := "your-service-name"
contextEnv := "your-context"
configList, err := client.ListAllConfigsByServiceAndContext(serviceName, contextEnv)
if err != nil {
    // Handle the error
}
```

### Listing Configurations by Dependent Job

You can retrieve configurations related to a specific service and source (dependent job) using the `ListAllConfigsByDependentJob` method. Provide the service and source as parameters to get the configurations associated with the given criteria.

```go
serviceName := "your-service-name"
source := "your-source"
configList, err := client.ListAllConfigsByDependentJob(serviceName, source)
if err != nil {
    // Handle the error
}
```

## Error Handling

In case of any errors during API requests, the methods return an error value that you can handle according to your application's needs.

## Contributing

Contributions to this library are welcome. If you find issues or have suggestions for improvements, feel free to create an issue or submit a pull request.

Feel free to customize this README to include information specific to your library and its usage.
