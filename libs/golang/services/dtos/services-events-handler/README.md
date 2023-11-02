# services-events-handler

The `services-events-handler` library contains Data Transfer Objects (DTOs) for the `services-events-handler` service. These DTOs are used to define the structure of data exchanged between different parts of the service, providing a clear and standardized representation of data.

## Usage

To use the DTOs provided by this library, import the necessary package in your Go code. You can then create instances of these DTOs to work with data in a structured and consistent way.

### Importing the Library

Import the library in your Go code as follows:

```go
import "libs/golang/services/dtos/services-events-handler/input"
import "libs/golang/services/dtos/services-events-handler/output"
import "libs/golang/services/dtos/services-events-handler/shared"
```

### Using the DTOs

You can use the DTOs in your code to represent and work with data structures related to the `services-events-handler` service.

#### Example:

```go
package main

import (
    "encoding/json"
    "fmt"
    inputDTO "libs/golang/services/dtos/services-events-handler/input"
    outputDTO "libs/golang/services/dtos/services-events-handler/output"
    sharedDTO "libs/golang/services/dtos/services-events-handler/shared"
)

func main() {
    // Create a sample input ServiceFeedbackDTO
    inputFeedback := inputDTO.ServiceFeedbackDTO{
        Data: map[string]interface{}{
            "param1": "value1",
            "param2": 42,
        },
        Metadata: sharedDTO.Metadata{
            Input: sharedDTO.MetadataInput{
                ID:                  "input123",
                Data:                map[string]interface{}{"inputParam": "inputValue"},
                ProcessingId:        "process123",
                ProcessingTimestamp: "2023-11-02T12:34:56",
                InputSchemaId:       "schema123",
            },
            Service:             "example-service",
            Source:              "example-source",
            Context:             "example-context",
            ProcessingTimestamp: "2023-11-02T12:34:56",
            JobFrequency:        "daily",
            JobConfigId:         "config123",
        },
        Status: sharedDTO.Status{
            Code: 200,
            Detail: "OK"
        },
    }

    // Convert the ConfigDTO to JSON
    inputFeedbackJSON, err := json.Marshal(inputFeedback)
    if err != nil {
        fmt.Println("Error marshaling JSON:", err)
        return
    }

    fmt.Println("ConfigDTO JSON:", string(inputFeedbackJSON))

}
```

## DTOs Provided

The library provides the following DTOs for use in the `services-events-handler` service:

### `ServiceFeedbackDTO` (input and output)

This DTO represents the feedback for a service, including various attributes such as name, service parameters, and job parameters. It is used for both input and output data.

### `MetadataInput` (shared)

This DTO represents metadata related to the input data, including ID, data, processing ID, processing timestamp, and input schema ID.

### `Metadata` (shared)

This DTO represents general metadata, including input metadata, service name, source, context, processing timestamp, job frequency, and job config ID.

### `Status` (shared)

This DTO represents the status of a service feedback, including a status code and a detailed description.

Please refer to the Go code and documentation for further details on the structure and usage of these DTOs.

## Contributions

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request.
