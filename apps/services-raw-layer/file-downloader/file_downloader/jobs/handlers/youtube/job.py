import io
from typing import Tuple

from pyyoutube.client import download_to_buffer
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
        _partition (str): The partition based on video id.
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
        self._partition = input_data.videoId
        self._target_endpoint = self._get_endpoint()

    def _get_endpoint(self) -> str:
        """
        Generates the target endpoint URL.

        Returns:
            str: The target endpoint URL.
        """
        return self._job_url.format(self._partition)

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

    def _get_status(self) -> StatusDTO:
        """
        Extracts the status information from an HTTP response.

        Args:
            response: The HTTP response.

        Returns:
            StatusDTO: The status information.
        """
        return StatusDTO(
            code=200,
            detail="Success",
        )

    def make_request(self) -> io.BytesIO:
        """
        Makes a video download from youtube by the id provided.

        Returns:
            io.BytesIO: The video bytes.
        """
        logger.info(f"endpoint: {self._target_endpoint}")
        return download_to_buffer(self._target_endpoint)


    def run(self) -> Tuple[dict, StatusDTO, str]:
        """
        Runs the job, making the HTTP request and handling the response.

        Returns:
            tuple: A tuple containing result data, status information, and the target endpoint.
        """
        logger.info(f"Job triggered with input: {self._input_data}")
        video = self.make_request()
        minio = minio_client()

        uri = minio.upload_bytes(self._get_bucket_name(), f"{self._partition}/video.mp4", video)


        logger.info(f"File storage uri: {uri}")
        result = {"documentUri": uri, "partition": self._partition}
        logger.info(f"Job result: {result}")
        return result, self._get_status(), self._target_endpoint
