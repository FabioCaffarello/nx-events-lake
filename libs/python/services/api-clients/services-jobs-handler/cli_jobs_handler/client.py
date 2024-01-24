from pyrequest.factory import RateLimitedAsyncHttpClient
from pysd.service_discovery import new_from_env
from dto_jobs_handler.output import JobParamsHttpGateway
from pyserializer.serializer import serialize_to_dataclass

class AsyncPyJobsHandlerClient:
    def __init__(self, base_url):
        self.__max_calls = 100
        self.__period = 60
        self.client = RateLimitedAsyncHttpClient(base_url, self.__max_calls, self.__period)

    async def list_one_job_params_http_gateway_by_context_n_service_n_source(self, context: str, service_name: str, source_name: str) -> JobParamsHttpGateway:

        endpoint = "/jobs-params/http-gateway/context/{context}/service/{service_name}/source/{source_name}".format(
            context=context,
            service_name=service_name,
            source_name=source_name
        )
        result = await self.client.make_request("GET", endpoint)
        return serialize_to_dataclass(result, JobParamsHttpGateway)

def async_py_jobs_handler_client():
    sd = new_from_env()
    return AsyncPyJobsHandlerClient(sd.services_jobs_handler_endpoint())
