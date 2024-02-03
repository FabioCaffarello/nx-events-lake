from datetime import datetime
from typing import Tuple, Union

import requests
import warlock
from dto_config_handler.output import ConfigDTO
from dto_events_handler.shared import StatusDTO
from source_watcher.handlers.default.html_utils import get_target_input_data_by_regex_pattern
from pylog.log import setup_logging
from pyminio.client import minio_client
from mod_base_job.http_gateway import BaseJob
from dto_jobs_handler.output import JobParamsHttpGateway
from py_external_request.factory_request import create_request
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
        _target_endpoint (str): The final endpoint URL.

    Methods:
        _get_reference(self, reference):
            Extracts and formats the reference data.

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

    def __init__(self, config: ConfigDTO, input_data: type[warlock.model.Model], dbg: Union[debug.DisabledDebug, debug.EnabledDebug]):
        self._config = config
        self._source = config.source
        self._context = config.context
        self._input_data = input_data
        super().__init__(service=config.service, source=config.source, context_env=config.context, dbg=dbg)

    def _get_reference(self, reference) -> str:
        """
        Extracts and formats the reference data.

        Args:
            reference: The reference data.

        Returns:
            str: The formatted reference string.
        """
        logger.info(f"Reference: {reference}")
        ref = datetime(reference.year, reference.month, reference.day)
        return ref.strftime("%Y%m%d")

    def _get_bucket_name(self) -> str:
        """
        Generates the bucket name for Minio storage.

        Returns:
            str: The bucket name.
        """
        return "process-input-{context}-source-{source}".format(
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
        return create_request(
            url=job_params.base_url,
            verify=False,
            headers=job_params.headers,
            timeout=10*60,
        )

    def parse_response_body(self, response: requests.Response) -> str:
        self.debug_response("file-source.html", response.content)
        """
        Parses the response body.

        Args:
            response: The HTTP response.

        Returns:
            str: The response body.
        """
        pattern_search = {
            "year": b'"ano" : "([0-9]{4})"',
            "month": b'"mes" : "([0][0-9]|[1][012])"',
            "day": b'"dia" : "([0-9]{2})"',
        }
        input_search = get_target_input_data_by_regex_pattern(response.content, pattern_search)
        if input_search is None:
            raise Exception("No data found")
        return "{year}{month}{day}".format(**input_search)


    async def run(self) -> Tuple[dict, StatusDTO]:
        """
        Runs the job, making the HTTP request and handling the response.

        Returns:
            tuple: A tuple containing result data, status information, and the target endpoint.
        """
        logger.info(f"Job triggered with input: {self._input_data}")
        job_params = await self.get_jobs_params()
        response = self.make_request(job_params)
        input_data = self.parse_response_body(response)
        minio = minio_client()
        uri = minio.upload_bytes(self._get_bucket_name(), f"{input_data}/{self._source}.html", response.content)
        logger.info(f"File storage uri: {uri}")
        result = {"documentUri": uri, "partition": input_data}
        logger.info(f"Job result: {result}")
        return result, self._get_status(response)
