import importlib
from typing import Tuple
import warlock
from dto_config_handler.output import ConfigDTO
from dto_events_handler.shared import StatusDTO
from pylog.log import setup_logging

logger = setup_logging(__name__)

class JobHandler:
    """
    Handles the execution of jobs associated with a configuration.

    Args:
        config (ConfigDTO): The configuration data.
        embeddings: Embeddings for job processing.
        dimension: Dimension for embeddings.

    Attributes:
        _config (ConfigDTO): The configuration data.
        _job_handler (str): The name of the job handler.
        _config_id (str): The unique identifier for the configuration.
        _embeddings: Embeddings for job processing.
        _dimension: Dimension for embeddings.
        _module: The imported module for the job handler.

    Methods:
        __init__(self, config: ConfigDTO, embeddings, dimension) -> None:
            Initializes the JobHandler instance.

        _import_job_handler_as_module(self) -> module:
            Imports the job handler module based on the specified name.

        run(self, source_input: type[warlock.model.Model]) -> Tuple[dict, StatusDTO]:
            Runs the job associated with the configuration.

    """
    def __init__(self, config: ConfigDTO, embeddings, dimension) -> None:
        """
        Initializes the JobHandler instance.

        Args:
            config (ConfigDTO): The configuration data.
            embeddings: Embeddings for job processing.
            dimension: Dimension for embeddings.

        Returns:
            None
        """
        self._config = config
        self._job_handler = config.service_parameters["job_handler"]
        self._config_id = config.id
        self._embeddings = embeddings
        self._dimension = dimension
        self._module = self._import_job_handler_as_module()

    def _import_job_handler_as_module(self):
        """
        Imports the job handler module based on the specified job handler name.

        Returns:
            module: The imported module for the job handler.
        """
        return importlib.import_module(f"jobs.handlers.{self._job_handler}.job")

    def run(self, source_input: type[warlock.model.Model]) -> Tuple[dict, StatusDTO]:
        """
        Runs the job associated with the configuration.

        Args:
            source_input (type[warlock.model.Model]): The input data for the job.

        Returns:
            Tuple[dict, StatusDTO]: A tuple containing job_data and job_status.
        """
        logger.info(f"[RUNNING JOB] - Config ID: {self._config_id} - handler: {self._job_handler}")
        job_data, job_status = self._module.Job(self._config, source_input, self._embeddings, self._dimension).run()
        return job_data, job_status
