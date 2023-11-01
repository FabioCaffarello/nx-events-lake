# services-output-handler

The `services-output-handler` library contains Data Transfer Objects (DTOs) for the `services-output-handler` service. These DTOs are used to define the structure of data exchanged between different parts of the service, providing a clear and standardized representation of data.

## Usage

To use the DTOs provided by this library, import the necessary package in your Go code. You can then create instances of these DTOs to work with data in a structured and consistent way.

### Importing the Library

Import the library in your Go code as follows:

```go
import "libs/golang/services/dtos/services-output-handler/input"
import "libs/golang/services/dtos/services-output-handler/output"
import "libs/golang/services/dtos/services-output-handler/shared"
```

### Using the DTOs

You can use the DTOs in your code to represent and work with data structures related to the `services-output-handler` service.

Here's an example of how to use the `ServiceOutputDTO` from the input package:

```go
import (
    inputDTO "libs/golang/services/dtos/services-output-handler/input"
    "encoding/json"
    "fmt"
)

func main() {
    // Create a ServiceOutputDTO instance
    data := map[string]interface{}{
        "key1": "value1",
        "key2": 42,
    }

    metadata := sharedDTO.Metadata{
        InputId: "input123",
        Input: sharedDTO.MetadataInput{
            ID:                  "metadataId",
            Data:                map[string]interface{}{"metaKey": "metaValue"},
            ProcessingId:        "processing123",
            ProcessingTimestamp: "2023-11-01T12:00:00Z",
        },
        Service: "my-service",
        Source:  "my-source",
    }

    serviceOutput := inputDTO.ServiceOutputDTO{
        Data:     data,
        Metadata: metadata,
        Context:  "my-context",
    }

    // Convert to JSON for output or further processing
    jsonBytes, err := json.Marshal(serviceOutput)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(string(jsonBytes))
}
```

## DTOs Provided

The library provides the following DTOs for use in the `services-output-handler` service:

### `ServiceOutputDTO` (input and output)

This DTO is used for input and output and represents the main data structure exchanged by the service. It includes data, metadata, and context information.

### `MetadataInput` (shared)

This DTO represents the metadata for input data. It includes information like ID, data, processing ID, and processing timestamp.

### `Metadata` (shared)

This DTO represents the metadata associated with the service output. It includes input ID, input metadata, service, source, and more.

## Contributions

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request.
