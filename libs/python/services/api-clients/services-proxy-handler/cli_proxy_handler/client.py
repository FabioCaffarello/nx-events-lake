from pyrequest.factory import RateLimitedAsyncHttpClient
from pysd.service_discovery import new_from_env
from dto_proxy_handler.output import TorProxyDTO
from pyserializer.serializer import serialize_to_dataclass


class AsyncPyProxyHandlerClient:
    def __init__(self, base_url):
        self.__max_calls = 100
        self.__period = 60
        self.client = RateLimitedAsyncHttpClient(base_url, self.__max_calls, self.__period)

    async def get_tor_proxy_ip_rotate(self):
        endpoint = "/tor"
        result = await self.client.make_request("GET", endpoint)
        return serialize_to_dataclass(result, TorProxyDTO)

def async_py_proxy_handler_client():
    sd = new_from_env()
    return AsyncPyProxyHandlerClient(sd.services_proxy_handler_endpoint())
