# service-file-catalog-handler

`service-file-catalog-handler` is a Python library that provides a client for interacting with a service file catalog management system. It allows you to create, list, and retrieve file catalogs data for various services.

## Installation

You can install `service-file-catalog-handler` using `nx`:

```sh
npx nx add <project> --name python-services-api-clients-services-file-catalog-handler --local
```

## Usage

### Initializing the Client

You can initialize the client by providing the base URL of your service configuration management system.

```python
file_catalog_handler_client = async_py_file_catalog_handler_client()
```

### Listing a File Catalog by ID

To retrieve a specific configuration by its ID, use the `list_one_file_catalog_by_id` method.

```python
file_catalog_id = "your_file_catalog_id"
file_catalog = await file_catalog_handler_client.list_one_file_catalog_by_id(file_catalog_id)
```

### Listing a File Catalog by Service and Source

To retrieve one file catalog associated with a specific service and source, use the `list_one_file_catalog_by_service_source` method.

```python
service_name = "your_service_name"
source_name = "your_source_name"
service_configs = await file_catalog_handler_client.list_one_file_catalog_by_service_source(service_name, source_name)
```

## Service Discovery

The library also provides a convenient function `async_py_file_catalog_handler_client` to create a client using service discovery. This function automatically retrieves the base URL from your environment using `new_from_env()`.

```python
file_catalog_handler_client = async_py_file_catalog_handler_client()
```

Make sure to set up your service discovery environment variables before using this function.

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
