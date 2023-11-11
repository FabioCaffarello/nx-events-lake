import io
import os
from typing import Tuple
import tempfile
import warlock
from moviepy.editor import VideoFileClip
from dto_config_handler.output import ConfigDTO
from dto_events_handler.shared import StatusDTO
from pylog.log import setup_logging
from pyminio.client import minio_client, MinioClient


logger = setup_logging(__name__)


class Job:
    """
    Represents a job that makes HTTP requests and handles the response.

    Args:
        config (ConfigDTO): The configuration data for the job.
        input_data (type[warlock.model.Model]): The input data for the job.

    Attributes:
        _config (ConfigDTO): The configuration data for the job.
        _source (str): The source information from the configuration.
        _context (str): The context information from the configuration.
        _input_data (type[warlock.model.Model]): The input data for the job.
        _partition (str): The partition based on video id.
        _target_endpoint (str): The final endpoint URL.

    Methods:
        _get_bucket_name(self, layer: str) -> str:
            Generates the bucket name for Minio storage.

        _get_status(self) -> StatusDTO:
            Extracts the status information from an HTTP response.

        make_request(self, minio: MinioClient, audio_path: str) -> None:
            Download a video from Minio, convert it to audio, and save the audio file.

        convert_video_bytes_to_audio(self, video_bytes: bytes, audio_path: str) -> None:
            Convert a video in bytes format to audio and save it as a separate file.

        run(self) -> Tuple[dict, StatusDTO, str]:
            Runs the job, making the HTTP request and handling the response.

    """

    def __init__(self, config: ConfigDTO, input_data: type[warlock.model.Model]) -> None:
        """
        Initialize the Job instance.

        Args:
            config (ConfigDTO): The configuration data for the job.
            input_data (type[warlock.model.Model]): The input data for the job.

        Returns:
            None
        """
        self._config = config
        self._source = config.source
        self._context = config.context
        self._input_data = input_data
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

        Args:
            response: The HTTP response.

        Returns:
            StatusDTO: The status information.
        """
        return StatusDTO(
            code=200,
            detail="Success",
        )

    def make_request(self, minio: MinioClient, audio_path: str) -> None:
        """
        Download a video from Minio, convert it to audio, and save the audio file.

        Args:
            minio (MinioClient): An instance of the MinioClient for interacting with Minio.
            audio_path (str): The local path where the audio file will be saved.

        Returns:
            None

        """
        logger.info(f"endpoint: {self._target_endpoint}")
        file_bytes = minio.download_file_as_bytes(self._get_bucket_name(layer="landing"), f"{self._partition}/video.mp4")
        self.convert_video_bytes_to_audio(file_bytes, audio_path)

    def convert_video_bytes_to_audio(self, video_bytes: bytes, audio_path: str) -> None:
        """
        Convert a video in bytes format to audio and save it as a separate file.

        Args:
            video_bytes (bytes): The video content in bytes.
            audio_path (str): The local path where the audio file will be saved.

        Raises:
            Exception: If there is an error converting the video to audio.
        """
        try:
            # Create a temporary file to write the video bytes
            with tempfile.NamedTemporaryFile(delete=False) as temp_video_file:
                temp_video_file.write(video_bytes)

            # Use VideoFileClip with the temporary file
            video_clip = VideoFileClip(temp_video_file.name)

            audio_clip = video_clip.audio
            audio_clip.write_audiofile(audio_path)
            audio_clip.close()
            video_clip.close()

            os.remove(temp_video_file.name)
            logger.info("Video converted to audio successfully.")
        except Exception as err:
            raise Exception(f"Error converting video to audio: {err}")


    def run(self) -> Tuple[dict, StatusDTO, str]:
        """
        Convert video content in bytes format to audio and save it as a separate file.

        Args:
            video_bytes (bytes): The video content in bytes.
            audio_path (str): The local path where the audio file will be saved.

        Raises:
            Exception: If there is an error converting the video to audio.

        """
        logger.info(f"Job triggered with input: {self._input_data}")
        minio = minio_client()
        audio_path = f"video-audio.mp3"
        video_audio = self.make_request(minio, audio_path)
                # Upload the audio file to the remote bucket
        with open(audio_path, 'rb') as audio_file:
            audio_data = audio_file.read()
        uri = minio.upload_bytes(self._get_bucket_name(layer="raw"), f"{self._partition}/audio.mp3", audio_data)
        os.remove(audio_path)
        logger.info(f"File storage uri: {uri}")
        result = {"documentUri": uri, "partition": self._partition}
        logger.info(f"Job result: {result}")
        return result, self._get_status(), self._target_endpoint
