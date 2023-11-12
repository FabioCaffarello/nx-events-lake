from typing import Tuple

import requests
import warlock
from dto_config_handler.output import ConfigDTO
from dto_events_handler.shared import StatusDTO
from pylog.log import setup_logging
from pyminio.client import minio_client


logger = setup_logging(__name__)


class Job:
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

    def __init__(self, config: ConfigDTO, input_data: type[warlock.model.Model]) -> None:
        self._config = config
        self._source = config.source
        self._context = config.context
        self._input_data = input_data
        self._job_url = config.job_parameters["url"]
        self._partition = input_data.partition
        self._target_document = input_data.targetDocument
        self._target_endpoint = self._get_endpoint()

    def _get_endpoint(self) -> str:
        """
        Generates the target endpoint URL.

        Returns:
            str: The target endpoint URL.
        """
        return self._job_url.format(self._target_document)

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

    def make_request(self) -> requests.Response:
        """
        Makes an HTTP GET request to the target endpoint.

        Returns:
            requests.Response: The HTTP response.
        """
        logger.info(f"endpoint: {self._target_endpoint}")
        headers = {
            "Sec-Fetch-Site": "same-origin",
            "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
            "Accept-Encoding": "gzip, deflate, br",
            "Accept-Language": "pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7",
            "Referer": f"https://www.annualreports.com/Company/{self._partition}",
            "User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:92.0) Gecko/20100101 Firefox/92.0",
        }
        return requests.get(
            self._target_endpoint,
            verify=False,
            headers=headers,
            timeout=10*60,
        )

    def run(self) -> Tuple[dict, StatusDTO, str]:
        """
        Runs the job, making the HTTP request and handling the response.

        Returns:
            tuple: A tuple containing result data, status information, and the target endpoint.
        """
        logger.info(f"Job triggered with input: {self._input_data}")
        response = self.make_request()
        minio = minio_client()
        uri = minio.upload_bytes(self._get_bucket_name(), f"{self._partition}/{self._target_document}", response.content)
        logger.info(f"File storage uri: {uri}")
        result = {"documentUri": uri, "partition": self._partition}
        logger.info(f"Job result: {result}")
        return result, self._get_status(response), self._target_endpoint
