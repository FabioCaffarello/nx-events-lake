# services-events-handler

The `services-events-handler` library contains Data Transfer Objects (DTOs) for the `services-events-handler` service. These DTOs are used to define the structure of data exchanged between different parts of the service, providing a clear and standardized representation of data.

## Usage

To use the DTOs provided by this library, import the necessary package in your Python code. You can then create instances of these DTOs to work with data in a structured and consistent way.

### Importing the Library

Import the library in your python code as follows:

```python
from dto_events_handler.output import ServiceFeedbackDTO
from dto_input_hadto_events_handlerndler.shared import StatusDTO, MetadataInputDTO, MetadataDTO
```

### Using the DTOs

You can use the DTOs in your code to represent and work with data structures related to the `services-events-handler` service.

Here's an example of how to use the `ServiceFeedbackDTO` from the output package:

```python
from dto_events_handler.output import ServiceFeedbackDTO
from dto_events_handler.shared import StatusDTO, MetadataInputDTO, MetadataDTO

# Create instances of the DTOs
feedback_data = {
    "key1": "value1",
    "key2": "value2"
}
metadata_input = MetadataInputDTO(
    id="123",
    data={"input_key": "input_value"},
    processing_id="456",
    processing_timestamp="2023-11-02T14:30:00",
    input_schema_id="789"
)
metadata = MetadataDTO(
    input=metadata_input,
    service="example-service",
    source="example-source",
    context="example-context",
    processing_timestamp="2023-11-02T14:30:00",
    job_frequency="daily",
    job_config_id="101"
)
status = StatusDTO(code=200, detail="OK")

feedback_dto = ServiceFeedbackDTO(data=feedback_data, metadata=metadata, status=status)

# Now you can work with the feedback DTO
print(feedback_dto.data)
print(feedback_dto.metadata.input.id)
print(feedback_dto.status.code)
```

## DTOs Provided

The library provides the following DTOs for use in the `services-events-handler` service:

### `ServiceFeedbackDTO` (output)

This DTO represents the main data structure exchanged by the service, including data.

### `MetadataDTO` (shared)

This DTO includes metadata information around the input, such as processing ID, processing timestamp, context, source, service and job config version id.

### `MetadataInputDTO` (shared)

This DTO includes metadata information around the output, such as processing timestamp, context, source, service and schema input version id.

### `StatusDTO` (shared)

The `StatusDTO` DTO contains information about the status of the input, including a status code and a detail message.

## Contributions

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
