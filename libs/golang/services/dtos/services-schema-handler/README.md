# services-schema-handler

The `services-schema-handler` library contains Data Transfer Objects (DTOs) for the `services-schema-handler` service. These DTOs are used to define the structure of data exchanged between different parts of the service, providing a clear and standardized representation of data.

## Usage

To use the DTOs provided by this library, import the necessary package in your Go code. You can then create instances of these DTOs to work with data in a structured and consistent way.

### Importing the Library

Import the library in your Go code as follows:

```go
import "libs/golang/services/dtos/services-schema-handler/input"
import "libs/golang/services/dtos/services-schema-handler/output"
```

### Using the DTOs

You can use the DTOs in your code to represent and work with data structures related to the `services-schema-handler` service.

#### Example:

```go
// Import the necessary packages
import (
    "libs/golang/services/dtos/services-schema-handler/input"
    "libs/golang/services/dtos/services-schema-handler/output"
)

// Create instances of DTOs
schemaInputDTO := input.SchemaDTO{
    SchemaType: "example",
    Service: "sample-service",
    Source: "sample-source",
    Context: "sample-context",
    JsonSchema: map[string]interface{}{
        "key1": "value1",
        "key2": "value2",
    },
}

schemaOutputDTO := output.SchemaDTO{
    ID: "12345",
    SchemaType: "example",
    Service: "sample-service",
    Source: "sample-source",
    Context: "sample-context",
    JsonSchema: map[string]interface{}{
        "key1": "value1",
        "key2": "value2",
    },
    SchemaID: "67890",
    CreatedAt: "2023-11-01T12:00:00Z",
    UpdatedAt: "2023-11-01T12:30:00Z",
}

// Use the DTOs as needed in your code
```

## DTOs Provided

The library provides the following DTOs for use in the `services-schema-handler` service:

### `SchemaDTO` (input and output)

This DTO represents schema information and includes attributes such as `SchemaType`, `Service`, `Source`, `Context`, and `JsonSchema`.

### `SchemaVersionData` (input and output)

This DTO represents data related to schema versions and includes attributes like `SchemaID` and a reference to the `SchemaDTO`.

### `SchemaVersionDTO` (input and output)

This DTO represents schema versions and includes attributes like `ID` and a list of `SchemaVersionData`.

Please refer to the Go code and documentation for further details on the structure and usage of these DTOs.

## Contributions

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request.
