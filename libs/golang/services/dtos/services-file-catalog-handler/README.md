# services-file-catalog-handler

The `services-file-catalog-handler` library contains Data Transfer Objects (DTOs) for the `services-file-catalog-handler` service. These DTOs are used to define the structure of data exchanged between different parts of the service, providing a clear and standardized representation of data.

## Usage

To use the DTOs provided by this library, import the necessary package in your Go code. You can then create instances of these DTOs to work with data in a structured and consistent way.

### Importing the Library

Import the library in your Go code as follows:

```go
import "libs/golang/services/dtos/services-file-catalog-handler/input"
import "libs/golang/services/dtos/services-file-catalog-handler/output"
```

### Using the DTOs

You can use the DTOs in your code to represent and work with data structures related to the `services-file-catalog-handler` service.

Here's an example of how to use the `FileCatalogDTO` from the input package:

```go
import (
	"fmt"
	inputDTO "libs/golang/services/dtos/services-file-catalog-handler/input"
)

func main() {
	// Example of using the FileCatalogDTO from the input package
	inputFileCatalog := inputDTO.FileCatalogDTO{
		Service:    "example-service",
		Source:     "example-source",
		Context:    "example-context",
		LakeLayer:  "example-lake-layer",
		SchemaType: "example-schema-type",
		Catalog: map[string]interface{}{
			"key1": ["value1"],
			"key2": ["value2"],
			"key3": ["value2"],
		},
	}

	fmt.Println("Input File Catalog DTO:")
	fmt.Printf("%+v\n", inputFileCatalog)
}
```

## DTOs Provided

The library provides the following DTOs for use in the `services-file-catalog-handler` service:

### `FileCatalogDTO` (input and output)

This DTO represents the main data structure exchanged by the service, including data.

Please refer to the Go code and documentation for further details on the structure and usage of these DTOs.

## Contributions

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
