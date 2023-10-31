# go-id

The `go-id` library provides utilities for generating and working with various types of identifiers (IDs) used in your Go applications. These IDs are designed to uniquely identify objects, schemas, or configurations within your systems.

## Usage

### Configuration ID
The `config` package provides tools for creating configuration-related IDs. You can use it as follows:

```golang
import "libs/golang/shared/go-id/config"

// Create a new configuration ID
id := config.NewID("my-service", "my-source")
```
### MD5 ID

The `md5` package allows you to generate MD5-based IDs for data and provides additional functionality for including source and service information:

```golang
import "libs/golang/shared/go-id/md5"

// Create an MD5 ID from a map of data
data := map[string]interface{}{"key1": "value1", "key2": "value2"}
id := md5.NewID(data)

// Create an MD5 ID with source information
idWithSource := md5.NewWithSourceID(data, "my-source")

// Create an MD5 ID with both source and service information
idWithSourceAndService := md5.NewWithSourceAndServiceID(data, "my-source", "my-service")
```

### Schema ID

The `schema` package helps you generate schema-related IDs based on schema type, service, and source:

```golang
import "libs/golang/shared/go-id/schema"

// Create a schema ID
schemaID := schema.NewID("user", "my-service", "my-source")
```

### UUID ID

The `uuid` package provides tools for working with UUID-based IDs:

```golang
import "libs/golang/shared/go-id/uuid"

// Create a new UUID ID
id := uuid.NewID()

// Parse a UUID ID from a string
parsedID, err := uuid.ParseID("7c9e6679-327c-43f8-9b2c-0d33d6d735f7")
```
