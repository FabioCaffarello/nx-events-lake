from typing import List
from pyrequest.factory import RateLimitedAsyncHttpClient
from pysd.service_discovery import new_from_env
from dto_config_handler.output import ConfigDTO
from pyserializer.serializer import from_data_to_dataclass

class AsyncPyConfigHandlerClient:
    """
    A client for handling asynchronous interactions with a configuration service.

    Args:
        base_url (str): The base URL of the configuration service.

    Attributes:
        __max_calls (int): The maximum number of API calls allowed in a period.
        __period (int): The time period (in seconds) in which API calls are rate-limited.
        client (RateLimitedAsyncHttpClient): An instance of the rate-limited HTTP client.

    """

    def __init__(self, base_url):
        self.__max_calls = 100
        self.__period = 60
        self.client = RateLimitedAsyncHttpClient(base_url, self.__max_calls, self.__period)

    async def create_config(self, data: dict) -> ConfigDTO:
        """
        Create a new configuration using the provided data.

        Args:
            data (dict): The configuration data to be created.

        Returns:
            ConfigDTO: The created configuration in the form of a data class.

        """
        endpoint = "/configs"
        config = await self.client.make_request("POST", endpoint, data=data)
        return from_data_to_dataclass(config, ConfigDTO)

    async def list_all_configs(self) -> List[ConfigDTO]:
        """
        Retrieve a list of all configurations.

        Returns:
            List[ConfigDTO]: A list of all configurations in the form of data classes.

        """
        endpoint = "/configs"
        configs = await self.client.make_request("GET", endpoint)
        return [from_data_to_dataclass(config, ConfigDTO) for config in configs]

    async def list_one_config_by_id(self, config_id: str)  -> ConfigDTO:
        """
        Retrieve a specific configuration by its ID.

        Args:
            config_id (str): The unique identifier of the configuration.

        Returns:
            ConfigDTO: The requested configuration in the form of a data class.

        """
        endpoint = f"/configs/{config_id}"
        config = await self.client.make_request("GET", endpoint)
        return from_data_to_dataclass(config, ConfigDTO)

    async def list_all_configs_by_service(self, service_name: str) -> List[ConfigDTO]:
        """
        Retrieve a list of configurations associated with a specific service.

        Args:
            service_name (str): The name of the service.

        Returns:
            List[ConfigDTO]: A list of configurations for the specified service in the form of data classes.

        """
        endpoint = f"/configs/service/{service_name}"
        configs = await self.client.make_request("GET", endpoint)
        return [from_data_to_dataclass(config, ConfigDTO) for config in configs]

    async def list_all_configs_by_service_and_context(self, service_name: str, context: str) -> list[ConfigDTO]:
        """
        Retrieve a list of configurations associated with a specific service and context.

        Args:
            service_name (str): The name of the service.
            context (str): The context associated with the service.

        Returns:
            list[ConfigDTO]: A list of configurations for the specified service and context in the form of data classes.

        """
        endpoint = f"/configs/service/{service_name}/context/{context}"
        configs = await self.client.make_request("GET", endpoint)
        return [from_data_to_dataclass(config, ConfigDTO) for config in configs]

def async_py_config_handler_client():
    """
    Create an instance of the AsyncPyConfigHandlerClient using service discovery information.

    Returns:
        AsyncPyConfigHandlerClient: An instance of the configuration handler client.

    """
    sd = new_from_env()
    return AsyncPyConfigHandlerClient(sd.services_config_handler_endpoint())
