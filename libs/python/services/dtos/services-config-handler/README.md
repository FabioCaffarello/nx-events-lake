# services-config-handler

The `services-config-handler` library contains Data Transfer Objects (DTOs) for the `services-config-handler` service. These DTOs are used to define the structure of data exchanged between different parts of the service, providing a clear and standardized representation of data.

## Usage

To use the DTOs provided by this library, import the necessary package in your Python code. You can then create instances of these DTOs to work with data in a structured and consistent way.

### Importing the Library

Import the library in your python code as follows:

```python
from dto_config_handler.output import InputDTO
from dto_config_handler.shared import StatusDTO, MetadataDTO
```

### Using the DTOs

You can use the DTOs in your code to represent and work with data structures related to the `services-config-handler` service.

Here's an example of how to use the `ConfigDTO` from the output package:

```python
from dto_config_handler.output import ConfigDTO
from dto_config_handler.shared import JobDependencies

# Create an instance of ConfigDTO with sample data
config_data = ConfigDTO(
    id="12345",
    name="SampleConfig",
    active=True,
    frequency="daily",
    service="services-config-handler",
    source="user",
    context="example",
    config_id="67890",
    depends_on=[
        JobDependencies(service="dependency-service", source="dependency-source"),
        JobDependencies(service="another-service", source="another-source"),
    ],
    service_parameters={
        "param1": "value1",
        "param2": "value2",
    },
    job_parameters={
        "job_param1": "job_value1",
        "job_param2": "job_value2",
    }
)

# Access the properties within the ConfigDTO
print("ID:", config_data.id)
print("Name:", config_data.name)
print("Active:", config_data.active)
print("Frequency:", config_data.frequency)
print("Service:", config_data.service)
print("Source:", config_data.source)
print("Context:", config_data.context)
print("Config ID:", config_data.config_id)
print("Depends On:", config_data.depends_on)
print("Service Parameters:", config_data.service_parameters)
print("Job Parameters:", config_data.job_parameters)
```

## DTOs Provided

The library provides the following DTOs for use in the `services-config-handler` service:

### `ConfigDTO` (output)

This DTO represents the main data structure exchanged by the service, including data.

### `JobDependencies` (shared)

This DTO includes the previous job dependency if exist

## Contributions

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
