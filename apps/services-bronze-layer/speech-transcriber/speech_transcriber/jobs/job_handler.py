import importlib
from typing import Tuple
from transformers import pipeline
import warlock
from dto_config_handler.output import ConfigDTO
from dto_events_handler.shared import StatusDTO
from pylog.log import setup_logging

logger = setup_logging(__name__)

class JobHandler:
    """
    Represents a job handler that runs a specific job based on configuration.

    Args:
        config (ConfigDTO): The configuration data for the job handler.

    Attributes:
        _config (ConfigDTO): The configuration data for the job handler.
        _job_handler (str): The name of the job handler.
        _config_id (int): The unique identifier for the configuration.
        _module (module): The imported module for the specified job handler.

    Methods:
        _import_job_handler_as_module(self) -> module:
            Imports the job handler module based on the provided configuration.

        run(self, source_input: type[warlock.model.Model]) -> Tuple[dict, StatusDTO, str]:
            Runs the job associated with the configuration.

            Args:
                source_input (type[warlock.model.Model]): The input data for the job.

            Returns:
                Tuple[dict, StatusDTO, str]: A tuple containing job_data, job_status, and target_endpoint.
    """
    def __init__(self, config: ConfigDTO, transcription_pipeline: pipeline) -> None:
        """
        Initialize the JobHandler instance.

        Args:
            config (ConfigDTO): The configuration data for the job handler.
            transcription_pipeline (pipeline): The pipeline for automatic speech recognition.

        Returns:
            None
        """
        self._config = config
        self._job_handler = config.service_parameters["job_handler"]
        self._config_id = config.id
        self._transcription_pipeline = transcription_pipeline
        self._module = self._import_job_handler_as_module()

    def _import_job_handler_as_module(self):
        """
        Import the job handler module based on the specified job handler name.

        Returns:
            module: The imported module for the job handler.
        """
        return importlib.import_module(f"jobs.handlers.{self._job_handler}.job")

    def run(self, source_input: type[warlock.model.Model]) -> Tuple[dict, StatusDTO]:
        """
        Run the job associated with the configuration.

        Args:
            source_input (type[warlock.model.Model]): The input data for the job.

        Returns:
            Tuple[dict, StatusDTO]: A tuple containing job_data and job_status.
        """
        logger.info(f"[RUNNING JOB] - Config ID: {self._config_id} - handler: {self._job_handler}")
        job_data, job_status = self._module.Job(self._config, source_input, self._transcription_pipeline).run()
        return job_data, job_status
