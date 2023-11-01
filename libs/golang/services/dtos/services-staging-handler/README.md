# services-staging-handler

The `services-staging-handler` library contains Data Transfer Objects (DTOs) for the `services-staging-handler` service. These DTOs are used to define the structure of data exchanged between different parts of the service, providing a clear and standardized representation of data.

## Usage

To use the DTOs provided by this library, import the necessary package in your Go code. You can then create instances of these DTOs to work with data in a structured and consistent way.

### Importing the Library

Import the library in your Go code as follows:

```go
import "libs/golang/services/dtos/services-staging-handler/input"
import "libs/golang/services/dtos/services-staging-handler/staging"
import "libs/golang/services/dtos/services-staging-handler/shared"
```

### Using the DTOs

You can use the DTOs in your code to represent and work with data structures related to the `services-staging-handler` service.

Here's an example of how to use the `ProcessingJobDependenciesDTO` from the input package:

```go
import (
    inputDTO "libs/golang/services/dtos/services-staging-handler/input"
    "encoding/json"
    "fmt"
)

func main() {
    // Create a ProcessingJobDependenciesDTO instance
    jobDependencies := []sharedDTO.ProcessingJobDependencies{
        {
            Service:             "dependency-service-1",
            Source:              "source-1",
            ProcessingId:        "process-1",
            ProcessingTimestamp: "2023-11-01T12:00:00Z",
            StatusCode:          200,
        },
        {
            Service:             "dependency-service-2",
            Source:              "source-2",
            ProcessingId:        "process-2",
            ProcessingTimestamp: "2023-11-01T13:00:00Z",
            StatusCode:          404,
        },
    }

    processingJob := inputDTO.ProcessingJobDependenciesDTO{
        Service:         "my-service",
        Source:          "my-source",
        Context:         "my-context",
        JobDependencies: jobDependencies,
    }

    // Convert to JSON for output or further processing
    jsonBytes, err := json.Marshal(processingJob)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(string(jsonBytes))
}
```

## DTOs Provided

The library provides the following DTOs for use in the `services-staging-handler` service:

### `ProcessingJobDependenciesDTO` (input and staging)

This DTO is used for input and staging and represents the data structure exchanged by the service. It includes information about the service, source, context, and an array of job dependencies.

### `ProcessingJobDependencies` (shared)

This DTO represents the shared information about job dependencies, including service, source, processing ID, processing timestamp, and status code.

### `Metadata` (shared)

This shared DTO may be used for additional metadata when working with the `services-staging-handler` service.

## Contributions

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
