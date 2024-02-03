from typing import Tuple, Union

import requests
import warlock
from dto_config_handler.output import ConfigDTO
from dto_events_handler.shared import StatusDTO
from mod_base_job.http_gateway import BaseJob
from pylog.log import setup_logging
from dto_jobs_handler.output import JobParamsHttpGateway
from py_external_request.factory_request import create_request
from pyminio.client import minio_client
import mod_debug.debug as debug


logger = setup_logging(__name__)


class Job(BaseJob):
    """
    Represents a job that makes HTTP requests and handles the response.

    Args:
        config (ConfigDTO): The configuration data for the job.
        input_data: The input data for the job.

    Attributes:
        _config (ConfigDTO): The configuration data for the job.
        _source  (str): The source information from the configuration.
        _context (str): The context information from the configuration.
        _input_data (type[warlock.model.Model]): The input data for the job.
        _job_url (str): The URL for the job.
        _partition (str): The partition based on input data reference.
        _target_endpoint (str): The final endpoint URL.

    Methods:

        _get_endpoint(self):
            Generates the target endpoint URL.

        _get_bucket_name(self):
            Generates the bucket name for Minio storage.

        _get_status(self, response) -> StatusDTO:
            Extracts the status information from an HTTP response.

        make_request(self):
            Makes an HTTP GET request to the target endpoint.

        run(self):
            Runs the job, making the HTTP request and handling the response.

    """

    def __init__(self, config: ConfigDTO, input_data: type[warlock.model.Model],  dbg: Union[debug.DisabledDebug, debug.EnabledDebug]) -> None:
        self._config = config
        self._source = config.source
        self._context = config.context
        self._input_data = input_data
        self._partition = input_data.partition
        self._target_endpoint = ""
        super().__init__(service=config.service, source=config.source, context_env=config.context, dbg=dbg)

    def _get_endpoint(self, job_params: JobParamsHttpGateway) -> str:
        return job_params.base_url.format(self._partition)

    def _get_bucket_name(self) -> str:
        """
        Generates the bucket name for Minio storage.

        Returns:
            str: The bucket name.
        """
        return "landing-{context}-source-{source}".format(
            context=self._context,
            source=self._source,
        )

    def _get_status(self, response) -> StatusDTO:
        """
        Extracts the status information from an HTTP response.

        Args:
            response: The HTTP response.

        Returns:
            StatusDTO: The status information.
        """
        return StatusDTO(
            code=response.status_code,
            detail=response.reason,
        )

    def make_request(self, job_params: JobParamsHttpGateway) -> requests.Response:
        """
        Makes an HTTP GET request to the target endpoint.

        Returns:
            requests.Response: The HTTP response.
        """
        logger.info(f"endpoint: {self._target_endpoint}")
        return create_request(
            url=self._target_endpoint,
            verify=False,
            headers=job_params.headers,
            timeout=10*60,
        )

    async def run(self) -> Tuple[dict, StatusDTO]:
        """
        Runs the job, making the HTTP request and handling the response.

        Returns:
            tuple: A tuple containing result data and status information.
        """
        logger.info(f"Job triggered with input: {self._input_data}")
        job_params = await self.get_jobs_params()
        self._target_endpoint = self._get_endpoint(job_params)
        response = self.make_request(job_params)
        self.debug_response("file-source.zip", response.content)
        minio = minio_client()
        uri = minio.upload_bytes(self._get_bucket_name(), f"{self._partition}/{self._source}.zip", response.content)
        logger.info(f"File storage uri: {uri}")
        result = {"documentUri": uri, "partition": self._partition}
        logger.info(f"Job result: {result}")
        return result, self._get_status(response)
