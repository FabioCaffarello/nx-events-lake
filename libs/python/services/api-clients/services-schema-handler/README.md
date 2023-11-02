# service-schema-handler

`service-schema-handler` is a Python library that provides a client for interacting with a service configuration management system. It allows you to create, list, and retrieve configuration data for various services.

## Installation

You can install `service-schema-handler` using `nx`:

```sh
npx nx add <project> --name python-services-api-clients-services-schema-handler --local
```

## Usage

### Initializing the Client

You can initialize the client by providing the base URL of your service configuration management system.

```python
schema_handler_client = async_py_schema_handler_client()
```


### Listing a schema by Service, Source and Schema Type

To retrieve a specific configuration by its ID, use the `list_one_schema_by_service_n_source_n_schema_type` method.

```python
service_name = "your_service_name"
source_name = "your_source_name"
schema_type = "your_schema_type"
config = await config_handler_client.list_one_schema_by_service_n_source_n_schema_type(
    service_name,
    source_name,
    schema_type
)
```

### Listing a schema by Context, Service, Source and Schema Type

To retrieve a specific configuration by its ID, use the `list_one_schema_by_service_n_source_n_context_n_schema_type` method.

```python
context_name = "your_context_name"
service_name = "your_service_name"
source_name = "your_source_name"
schema_type = "your_schema_type"
config = await config_handler_client.list_one_schema_by_service_n_source_n_context_n_schema_type(
    context_name,
    service_name,
    source_name,
    schema_type
)
```

## Service Discovery

The library also provides a convenient function `async_py_schema_handler_client` to create a client using service discovery. This function automatically retrieves the base URL from your environment using `new_from_env()`.

```python
config_handler_client = async_py_schema_handler_client()
```

Make sure to set up your service discovery environment variables before using this function.

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
