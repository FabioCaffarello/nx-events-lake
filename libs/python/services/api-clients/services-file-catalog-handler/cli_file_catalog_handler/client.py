from pyrequest.factory import RateLimitedAsyncHttpClient
from pysd.service_discovery import new_from_env
from dto_file_catalog_handler.output import FileCatalogDTO
from pyserializer.serializer import serialize_to_dataclass


class AsyncPyFileCatalogHandlerClient:
    def __init__(self, base_url):
        self.__max_calls = 100
        self.__period = 60
        self.client = RateLimitedAsyncHttpClient(base_url, self.__max_calls, self.__period)

    async def list_one_file_catalog_by_id(self, file_catalog_id: str):
        endpoint = "/file-catalog/{file_catalog_id}".format(
            file_catalog_id=file_catalog_id
        )
        result = await self.client.make_request("GET", endpoint)
        return serialize_to_dataclass(result, FileCatalogDTO)

    async def list_one_file_catalog_by_service_source(self, service_name: str, source_name: str):
        endpoint = "/file-catalog/service/{service_name}/source/{source_name}".format(
            service_name=service_name,
            source_name=source_name
        )
        result = await self.client.make_request("GET", endpoint)
        return serialize_to_dataclass(result, FileCatalogDTO)

def async_py_file_catalog_handler_client():
    sd = new_from_env()
    return AsyncPyFileCatalogHandlerClient(sd.services_file_catalog_handler_endpoint())
