# config-loader

## Introduction

`config-loader` is a Python library designed to simplify the process of loading configurations for a given service and context environment. This library provides the tools to fetch, register, and organize configurations in a structured manner.

## Installation

You can install `config-loader` using `nx`:

```sh
npx nx add <project> --name python-services-setup-config-loader --local
```

## Usage

To use the library, you need to import the necessary components and follow the provided pattern for fetching and registering configurations.

### Initializing `mapping_config`

`mapping_config` is a dictionary that will store configurations organized by context and ID. It needs to be initialized as an empty dictionary.

```python
mapping_config: Dict[str, Dict[str, ConfigDTO]] = dict()
```

### Fetching Configurations

Use the `fetch_configs` function to fetch configurations for a given service and context environment.

```python
async def fetch_configs(service: str, context_env: str) -> Dict[str, Dict[str, ConfigDTO]]:
    # Fetch configurations for the specified service and context environment.
    await ConfigLoader().fetch_configs_for_service(service_name=service, context_env=context_env)
    return mapping_config
```

### The `ConfigLoader` Class

The `ConfigLoader` class is responsible for loading configurations from the config handler API client.

```python
class ConfigLoader:
    def __init__(self) -> None:
        # Initialize the ConfigLoader.
        # This class is used to load configurations from the config handler API client.
        self.__config_handler_api_client = async_py_config_handler_client()
        super().__init()

    async def fetch_configs_for_service(self, service_name: str, context_env: str) -> None:
        # Fetch configurations for a specific service and context environment.
        # This method fetches configurations and registers them using the `register_config` function.
        configs = await self.__config_handler_api_client.list_all_configs_by_service_and_context(service_name, context_env)
        for config in configs:
            register_config(
                config.context,
                config.id,
                config
            )
```

### Registering Configurations

To register a configuration, use the `register_config` function. It registers configurations in the `mapping_config` dictionary.

```python
def register_config(context: str, config_id: str, config: ConfigDTO) -> None:
    # Register a configuration in the mapping_config dictionary.
    if context not in mapping_config:
        mapping_config[context] = dict()
    if config_id in mapping_config[context]:
        raise Exception(f"Duplicate config ID '{config_id}' for context '{context}'. Overwriting existing config.")
    mapping_config[context][config_id] = config
```
