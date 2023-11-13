from typing import Tuple
import warlock
from dto_config_handler.output import ConfigDTO
from dto_events_handler.shared import StatusDTO
from pylog.log import setup_logging
from transformers import pipeline
from pyminio.client import minio_client, MinioClient


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
        _transcription_pipeline (pipeline): The pipeline for transcription.
        _partition (str): The partition based on video id.
        _target_endpoint (str): The final endpoint URL.

    Methods:
        _get_bucket_name(self, layer: str) -> str:
            Generates the bucket name for Minio storage.

        _get_status(self) -> StatusDTO:
            Extracts the status information from an HTTP response.

        make_request(self, minio: MinioClient, audio_path: str) -> bytes:
            Makes a request to Minio for audio data.

        run(self) -> Tuple[dict, StatusDTO, str]:
            Runs the job, making the HTTP request and handling the response.

    """
    def __init__(self, config: ConfigDTO, input_data: type[warlock.model.Model], transcription_pipeline: pipeline) -> None:
        """
        Initialize the Job instance.

        Args:
            config (ConfigDTO): The configuration data for the job.
            input_data (type[warlock.model.Model]): The input data for the job.
            transcription_pipeline (pipeline): The pipeline for transcription.

        Returns:
            None
        """
        self._config = config
        self._source = config.source
        self._context = config.context
        self._input_data = input_data
        self._transcription_pipeline = transcription_pipeline
        self._partition = input_data.partition
        self._target_endpoint = input_data.documentUri

    def _get_bucket_name(self, layer: str) -> str:
        """
        Generates the bucket name for Minio storage.

        Args:
            layer (str): The layer of the bucket.

        Returns:
            str: The bucket name.
        """
        return "{layer}-{context}-source-{source}".format(
            layer=layer,
            context=self._context,
            source=self._source,
        )

    def _get_status(self) -> StatusDTO:
        """
        Extracts the status information from an HTTP response.

        Returns:
            StatusDTO: The status information.
        """
        return StatusDTO(
            code=200,
            detail="Success",
        )

    def make_request(self, minio: MinioClient) -> None:
        """
        Makes a request to Minio for audio data.

        Args:
            minio (MinioClient): An instance of the MinioClient for interacting with Minio.

        Returns:
            bytes: The audio data in bytes.
        """
        logger.info(f"endpoint: {self._target_endpoint}")
        file_bytes = minio.download_file_as_bytes(self._get_bucket_name(layer="raw"), f"{self._partition}/audio.mp3")
        return file_bytes


    def run(self) -> Tuple[dict, StatusDTO]:
        """
        Runs the job, making the HTTP request and handling the response.

        Returns:
            Tuple[dict, StatusDTO]: A tuple containing job_data and job_status.
        """
        logger.info(f"Job triggered with input: {self._input_data}")
        minio = minio_client()
        audio = self.make_request(minio)
        transcription = self._transcription_pipeline(audio)
        transcription_text = transcription["text"]

        uri = minio.upload_bytes(self._get_bucket_name(layer="bronze"), f"{self._partition}/transcription.txt", transcription_text.encode("utf-8"))
        logger.info(f"File storage uri: {uri}")
        result = {"documentUri": uri, "partition": self._partition}
        logger.info(f"Job result: {result}")
        return result, self._get_status()
