# services-output-handler

The `services-staging-handler` API Client is a Go library that allows you to interact with a remote HTTP API for managing staging data of processing jobs. This client provides a set of methods to perform various operations on staging data of processing jobs, such as creating new staging data for a dependent job og the actual job running, listing one staging data of processing jobs by id, remove a staging data of processing jobs and update an status of a the running job in the staging data of processing jobs.

## Getting Started

### Creating a Client

To create a new instance of the `services-staging-handler` API Client, use the `NewClient` function. It initializes the client with default configuration:

```go
client := stagingClient.NewClient()
```

You can also customize the client's configuration by setting a custom base URL or providing a context.

### Creating a Processing Job Dependencies

You can create a new staging processing job dependency using the `CreateProcessingJobDependencies` method. It takes a `ProcessingJobDependenciesDTO` as input and sends a POST request to the API to create a new processing job dependencies.

```go
processingDepData := inputDTO.ProcessingJobDependenciesDTO{
    // Set your processing job dependencies details here
}

processingDep, err := client.CreateProcessingJobDependencies(processingDepData)
if err != nil {
    // Handle the error
}
```

### List One Processing Job Dependencies by ID

```go
processingDeps, err := client.ListOneProcessingJobDependenciesById("your-processing-job-dependencies-id")
if err != nil {
    // Handle the error
}
```

### Remove a Processing Job Dependencies


```go
processingDeps, err := client.RemoveProcessingJobDependencies("your-processing-job-dependencies-id")
if err != nil {
    // Handle the error
}
```

### Update Job Dependencies of a Processing Job Dependencies


```go
processingDepData := sharedDTO.ProcessingJobDependencies{
    // Set your processing job dependencies details here
}

processingDeps, err := client.UpdateProcessingJobDependencies("your-processing-job-dependencies-id")
if err != nil {
    // Handle the error
}
```

## Error Handling

In case of any errors during API requests, the methods return an error value that you can handle according to your application's needs.

## Contributing

Contributions to this library are welcome. If you find issues or have suggestions for improvements, feel free to create an issue or submit a pull request.

Feel free to customize this README to include information specific to your library and its usage.
