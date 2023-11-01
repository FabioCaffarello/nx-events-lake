# services-config-handler

The `services-config-handler` library contains Data Transfer Objects (DTOs) for the `services-config-handler` service. These DTOs are used to define the structure of data exchanged between different parts of the service, providing a clear and standardized representation of data.

## Usage

To use the DTOs provided by this library, import the necessary package in your Go code. You can then create instances of these DTOs to work with data in a structured and consistent way.

### Importing the Library

Import the library in your Go code as follows:

```go
import "libs/golang/services/dtos/services-config-handler/input"
import "libs/golang/services/dtos/services-config-handler/output"
import "libs/golang/services/dtos/services-config-handler/shared"
```

### Using the DTOs

You can use the DTOs in your code to represent and work with data structures related to the `services-config-handler` service.

#### Example:

```go
import (
    inputDTO "libs/golang/services/dtos/services-config-handler/input"
    "encoding/json"
    "fmt"
)

func main() {
    // Create a ConfigDTO instance
    config := inputDTO.ConfigDTO{
        Name:      "Sample Config",
        Active:    true,
        Frequency: "daily",
        Service:   "sample-service",
        // ... other fields
    }

    // Convert the ConfigDTO to JSON
    configJSON, err := json.Marshal(config)
    if err != nil {
        fmt.Println("Error marshaling JSON:", err)
        return
    }

    fmt.Println("ConfigDTO JSON:", string(configJSON))
}
```

## DTOs Provided

The library provides the following DTOs for use in the `services-config-handler` service:

### `ConfigDTO` (input and output packages)

This DTO represents the configuration for a service, including various attributes such as name, service parameters, and job parameters. It is used for both input and output data.

### `ConfigVersionData` (input and output packages)

This DTO represents data related to a specific configuration version, including the ConfigID and the corresponding ConfigDTO. It is used for both input and output data.

### `ConfigVersionDTO` (input and output packages)

This DTO represents a collection of configuration versions associated with a specific ID. It includes an ID and a list of `ConfigVersionData` instances. It is used for both input and output data.

### `JobDependencies` (shared package)

This DTO represents job dependencies with service and source attributes. It is used to define dependencies between jobs.

Please refer to the Go code and documentation for further details on the structure and usage of these DTOs.

## Contributions

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request.
