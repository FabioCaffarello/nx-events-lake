# services-input-handler

The `services-input-handler` library contains Data Transfer Objects (DTOs) for the `services-input-handler` service. These DTOs are used to define the structure of data exchanged between different parts of the service, providing a clear and standardized representation of data.

## Usage

To use the DTOs provided by this library, import the necessary package in your Python code. You can then create instances of these DTOs to work with data in a structured and consistent way.

### Importing the Library

Import the library in your python code as follows:

```python
from dto_input_handler.output import InputDTO
from dto_input_handler.shared import StatusDTO, MetadataDTO
```

### Using the DTOs

You can use the DTOs in your code to represent and work with data structures related to the `services-input-handler` service.

Here's an example of how to use the `InputDTO` from the input package:

```python
from dto_input_handler.output import InputDTO
from dto_input_handler.shared import StatusDTO, MetadataDTO

# Create an InputDTO object
input_data = InputDTO(
    data={
        "key1": "value1",
        "key2": "value2",
    },
    status=StatusDTO(
        code=200,
        message="Success",
    ),
    metadata=MetadataDTO(
        processing_id=12345,
        processing_timestamp="2023-11-02T15:30:00",
        context="example",
        source="user",
        service="services-input-handler",
    )
)

# Access the data within the InputDTO
print("Data: ", input_data.data)
print("Status Code: ", input_data.status.code)
print("Status Message: ", input_data.status.message)
print("Processing ID: ", input_data.metadata.processing_id)
print("Processing Timestamp: ", input_data.metadata.processing_timestamp)
print("Context: ", input_data.metadata.context)
print("Source: ", input_data.metadata.source)
print("Service: ", input_data.metadata.service)
```

## DTOs Provided

The library provides the following DTOs for use in the `services-input-handler` service:

### `InputDTO` (output)

This DTO represents the main data structure exchanged by the service, including data.

### `MetadataDTO` (shared)

This DTO includes metadata information, such as processing ID, processing timestamp, context, source, and service.

### `StatusDTO` (shared)

The `StatusDTO` DTO contains information about the status of the input, including a status code and a detail message.

Please refer to the Go code and documentation for further details on the structure and usage of these DTOs.

## Contributions

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
