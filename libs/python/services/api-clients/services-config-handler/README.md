# service-config-handler

`service-config-handler` is a Python library that provides a client for interacting with a service configuration management system. It allows you to create, list, and retrieve configuration data for various services.

## Installation

You can install `service-config-handler` using `nx`:

```sh
npx nx add <project> --name python-services-api-clients-services-config-handler --local
```

## Usage

### Initializing the Client

You can initialize the client by providing the base URL of your service configuration management system.

```python
config_handler_client = async_py_config_handler_client()
```

### Creating a Configuration

To create a new configuration, use the `create_config` method, passing a dictionary of configuration data.

```python
data = {
    "key": "value",
    "another_key": "another_value"
}
new_config = await config_handler_client.create_config(data)
```

### Listing All Configurations

To list all configurations, use the `list_all_configs` method.

```python
configs = await config_handler_client.list_all_configs()
```

### Listing a Configuration by ID

To retrieve a specific configuration by its ID, use the `list_one_config_by_id` method.

```python
config_id = "your_config_id"
config = await config_handler_client.list_one_config_by_id(config_id)
```

### Listing Configurations by Service

To retrieve all configurations associated with a specific service, use the `list_all_configs_by_service` method.

```python
service_name = "your_service_name"
service_configs = await config_handler_client.list_all_configs_by_service(service_name)
```

### Listing Configurations by Service and Context

To retrieve configurations associated with a specific service and context, use the `list_all_configs_by_service_and_context` method.

```python
service_name = "your_service_name"
context = "your_context"
service_context_configs = await config_handler_client.list_all_configs_by_service_and_context(service_name, context)
```

## Service Discovery

The library also provides a convenient function `async_pycontroller_client` to create a client using service discovery. This function automatically retrieves the base URL from your environment using `new_from_env()`.

```python
config_handler_client = async_pycontroller_client()
```

Make sure to set up your service discovery environment variables before using this function.

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
