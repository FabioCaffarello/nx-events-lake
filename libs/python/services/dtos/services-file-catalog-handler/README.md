# services-file-catalog-handler

The `services-file-catalog-handler` library contains Data Transfer Objects (DTOs) for the `services-file-catalog-handler` service. These DTOs are used to define the structure of data exchanged between different parts of the service, providing a clear and standardized representation of data.

## Usage

To use the DTOs provided by this library, import the necessary package in your Python code. You can then create instances of these DTOs to work with data in a structured and consistent way.

### Importing the Library

Import the library in your python code as follows:

```python
from dto_file_catalog_handler.output import SchemaDTO
```

### Using the DTOs

You can use the DTOs in your code to represent and work with data structures related to the `services-file-catalog-handler` service.

Here's an example of how to use the `FileCatalogDTO` from the output package:

```python
from dataclasses import dataclass, field
from typing import Dict
from dto_file_catalog_handler.output import FileCatalogDTO

# Create an instance of FileCatalogDTO
file_catalog_data = {
    "id": "123",
    "service": "file-catalog-handler",
    "source": "example_source",
    "context": "example_context",
    "lake_layer": "raw",
    "schema_type": "avro",
    "catalog_id": "456",
    "catalog": {"field1": "value1", "field2": "value2"},
    "created_at": "2023-11-11T12:00:00Z",
    "updated_at": "2023-11-11T13:30:00Z"
}

file_catalog_instance = FileCatalogDTO(**file_catalog_data)

# Accessing attributes
print(f"File ID: {file_catalog_instance.id}")
print(f"Service: {file_catalog_instance.service}")
print(f"Source: {file_catalog_instance.source}")
print(f"Context: {file_catalog_instance.context}")
print(f"Lake Layer: {file_catalog_instance.lake_layer}")
print(f"Schema Type: {file_catalog_instance.schema_type}")
print(f"Catalog ID: {file_catalog_instance.catalog_id}")
print(f"Catalog Data: {file_catalog_instance.catalog}")
print(f"Created At: {file_catalog_instance.created_at}")
print(f"Updated At: {file_catalog_instance.updated_at}")
```

## DTOs Provided

The library provides the following DTOs for use in the `services-file-catalog-handler` service:

### `FileCatalogDTO` (output)

This DTO represents the main data structure exchanged by the service, including data.


## Contributions

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
