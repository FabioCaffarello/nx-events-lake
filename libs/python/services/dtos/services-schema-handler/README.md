# services-schema-handler

The `services-schema-handler` library contains Data Transfer Objects (DTOs) for the `services-schema-handler` service. These DTOs are used to define the structure of data exchanged between different parts of the service, providing a clear and standardized representation of data.

## Usage

To use the DTOs provided by this library, import the necessary package in your Python code. You can then create instances of these DTOs to work with data in a structured and consistent way.

### Importing the Library

Import the library in your python code as follows:

```python
from dto_schema_handler.output import SchemaDTO
```

### Using the DTOs

You can use the DTOs in your code to represent and work with data structures related to the `services-schema-handler` service.

Here's an example of how to use the `SchemaDTO` from the output package:

```python
from dto_schema_handler.output import SchemaDTO

# Create an instance of SchemaDTO with sample data
schema_data = SchemaDTO(
    id="12345",
    schema_type="data_schema",
    service="services-schema-handler",
    source="user",
    context="example",
    json_schema={
        "type": "object",
        "properties": {
            "name": {"type": "string"},
            "age": {"type": "integer"}
        }
    },
    schema_id="67890"
)

# Access the properties within the SchemaDTO
print("ID:", schema_data.id)
print("Schema Type:", schema_data.schema_type)
print("Service:", schema_data.service)
print("Source:", schema_data.source)
print("Context:", schema_data.context)
print("JSON Schema:", schema_data.json_schema)
print("Schema ID:", schema_data.schema_id)
```

## DTOs Provided

The library provides the following DTOs for use in the `services-schema-handler` service:

### `SchemaDTO` (output)

This DTO represents the main data structure exchanged by the service, including data.


## Contributions

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
