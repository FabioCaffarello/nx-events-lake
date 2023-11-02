
from typing import Dict
from pylog.log import setup_logging
from cli_config_handler.client import async_py_config_handler_client
from dto_config_handler.output import ConfigDTO

logger = setup_logging(__name__)


mapping_config: Dict[str, Dict[str, ConfigDTO]] = dict()


async def fetch_configs(service: str, context_env: str) -> Dict[str, Dict[str, ConfigDTO]]:
    """
    Fetch configurations for a given service and context environment.

    Args:
        service (str): The name of the service for which configurations are to be fetched.
        context_env (str): The context environment for which configurations are to be fetched.

    Returns:
        Dict[str, Dict[str, ConfigDTO]]: A dictionary containing configurations organized by context and ID.

    Note:
        The `mapping_config` dictionary will be populated by this function.

    """
    await ConfigLoader().fetch_configs_for_service(service_name=service, context_env=context_env)
    return mapping_config


class ConfigLoader:
    def __init__(self) -> None:
        """
        Initialize the ConfigLoader.

        Note:
            This class is used to load configurations from the config handler API client.

        """
        self.__config_handler_api_client = async_py_config_handler_client()
        super().__init__()

    async def fetch_configs_for_service(self, service_name: str, context_env: str) -> None:
        """
        Fetch configurations for a specific service and context environment.

        Args:
            service_name (str): The name of the service for which configurations are to be fetched.
            context_env (str): The context environment for which configurations are to be fetched.

        Returns:
            None

        Note:
            This method fetches configurations and registers them using the `register_config` function.

        """
        configs = await self.__config_handler_api_client.list_all_configs_by_service_and_context(service_name, context_env)
        for config in configs:
            register_config(
                config.context,
                config.id,
                config
            )


def register_config(context: str, config_id: str, config: ConfigDTO) -> None:
    """
    Register a configuration in the mapping_config dictionary.

    Args:
        context (str): The context for which the configuration is being registered.
        config_id (str): The ID of the configuration.
        config (ConfigDTO): The configuration object to be registered.

    Returns:
        None

    Raises:
        Exception: If a configuration with the same ID already exists
    """
    if context not in mapping_config:
        mapping_config[context] = dict()
    if config_id in mapping_config[context]:
        raise Exception(f"Duplicate config ID '{config_id}' for context '{context}'. Overwriting existing config.")
    mapping_config[context][config_id] = config
