from pyrequest.factory import RateLimitedAsyncHttpClient
from pysd.service_discovery import new_from_env
from dto_schema_handler.output import SchemaDTO
from pyserializer.serializer import serialize_to_dataclass


class AsyncPySchemaHandlerClient:
    """
    A client for handling asynchronous interactions with a schema repository service.

    Args:
        base_url (str): The base URL of the schema repository service.

    Attributes:
        __max_calls (int): The maximum number of API calls allowed in a period.
        __period (int): The time period (in seconds) in which API calls are rate-limited.
        client (RateLimitedAsyncHttpClient): An instance of the rate-limited HTTP client.

    """
    def __init__(self, base_url):
        self.__max_calls = 100
        self.__period = 60
        self.client = RateLimitedAsyncHttpClient(base_url, self.__max_calls, self.__period)

    async def create(self, schema: SchemaDTO) -> SchemaDTO:
        """
        Create a schema.

        Args:
            schema (SchemaDTO): The schema to create.

        Returns:
            SchemaDTO: The created schema in the form of a data class.

        """
        endpoint = "/schemas"
        result = await self.client.make_request("POST", endpoint, schema)
        return serialize_to_dataclass(result, SchemaDTO)

    async def list_one_schema_by_service_n_source_n_schema_type(self, service_name: str, source_name: str, schema_type: str) -> SchemaDTO:
        """
        Retrieve a specific schema by service name, source name, and schema type.

        Args:
            service_name (str): The name of the service.
            source_name (str): The name of the data source.
            schema_type (str): The type of the schema.

        Returns:
            SchemaDTO: The requested schema in the form of a data class.

        """
        endpoint = "/schemas/service/{service_name}/source/{source_name}/schema-type/{schema_type}".format(
            service_name=service_name,
            source_name=source_name,
            schema_type=schema_type
        )
        result = await self.client.make_request("GET", endpoint)
        return serialize_to_dataclass(result, SchemaDTO)

    async def list_one_schema_by_service_n_source_n_context_n_schema_type(self, context: str, service_name: str, source_name: str, schema_type: str) -> SchemaDTO:
        """
        Retrieve a specific schema by service name, source name, context, and schema type.

        Args:
            context (str): The context associated with the service.
            service_name (str): The name of the service.
            source_name (str): The name of the data source.
            schema_type (str): The type of the schema.

        Returns:
            SchemaDTO: The requested schema in the form of a data class.

        """
        endpoint = "/schemas/service/{service_name}/source/{source_name}/context/{context}/schema-type/{schema_type}".format(
            service_name=service_name,
            source_name=source_name,
            context=context,
            schema_type=schema_type
        )
        result = await self.client.make_request("GET", endpoint)
        return serialize_to_dataclass(result, SchemaDTO)




def async_py_schema_handler_client():
    sd = new_from_env()
    return AsyncPySchemaHandlerClient(sd.services_schemas_handler_endpoint())
