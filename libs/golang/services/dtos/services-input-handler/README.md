# services-input-handler

The `services-input-handler` library contains Data Transfer Objects (DTOs) for the `services-input-handler` service. These DTOs are used to define the structure of data exchanged between different parts of the service, providing a clear and standardized representation of data.

## Usage

To use the DTOs provided by this library, import the necessary package in your Go code. You can then create instances of these DTOs to work with data in a structured and consistent way.

### Importing the Library

Import the library in your Go code as follows:

```go
import "libs/golang/services/dtos/services-input-handler/input"
import "libs/golang/services/dtos/services-input-handler/output"
import "libs/golang/services/dtos/services-input-handler/shared"
```

### Using the DTOs

You can use the DTOs in your code to represent and work with data structures related to the `services-input-handler` service.

Here's an example of how to use the `InputDTO` from the input package:

```go
import (
    inputDTO "libs/golang/services/dtos/services-input-handler/input"
    "encoding/json"
    "fmt"
)

func main() {
    // Create an InputDTO instance
    data := map[string]interface{}{
        "key1": "value1",
        "key2": 42,
    }

    input := inputDTO.InputDTO{
        Data: data,
    }

    // Convert to JSON for output or further processing
    jsonBytes, err := json.Marshal(input)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(string(jsonBytes))
}
```

## DTOs Provided

The library provides the following DTOs for use in the `services-input-handler` service:

### `InputDTO` (input and output)

This DTO represents the main data structure exchanged by the service, including data.

### `Metadata` (shared)

This DTO includes metadata information, such as processing ID, processing timestamp, context, source, and service.

### `Status` (shared)

The `Status` DTO contains information about the status of the input, including a status code and a detail message.

Please refer to the Go code and documentation for further details on the structure and usage of these DTOs.

## Contributions

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
