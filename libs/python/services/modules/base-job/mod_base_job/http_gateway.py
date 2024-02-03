
from typing import Union
from cli_jobs_handler.client import async_py_jobs_handler_client
from dto_jobs_handler.output import JobParamsHttpGateway
import mod_debug.debug as debug


class BaseJob:
    def __init__(self, service: str, source: str, context_env: str, dbg: Union[debug.DisabledDebug, debug.EnabledDebug]):
        self.__debug = dbg
        self._service = service
        self._source = source
        self._context_env = context_env
        self._jobs_handler_client = async_py_jobs_handler_client()

    async def get_jobs_params(self) -> JobParamsHttpGateway:
        return await self._jobs_handler_client.list_one_job_params_http_gateway_by_context_n_service_n_source(
            context=self._context_env,
            service_name=self._service,
            source_name=self._source
        )

    def debug_response(self, file_name, response_body):
        self.__debug.save_response(file_name, response_body)
