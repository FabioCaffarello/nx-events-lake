import importlib
from typing import Tuple, Union, Dict, Any, List

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
        _import_job_handler_as_module(self):
            Imports the job handler module based on the provided configuration.

        _apply_result_mask(self, job_data):
            Applies a result mask to the given job data.

        run(self, source_input):
            Runs the job associated with the configuration.

            Args:
                source_input: The input data for the job.

            Returns:
                tuple: A tuple containing job_data, job_status, and target_endpoint.

    """

    def __init__(self, config: ConfigDTO) -> None:
        self._config = config
        self._job_handler = config.service_parameters["job_handler"]
        self._config_id = config.id
        self._module = self._import_job_handler_as_module()

    def _import_job_handler_as_module(self):
        """
        Import the job handler module based on the specified job handler name.

        Returns:
            module: The imported module for the job handler.
        """
        return importlib.import_module(f"handlers.{self._job_handler}.job")

    def _apply_result_mask(self, job_data: Union[Dict[str, Any], List[Dict[str, Any]]]) -> Dict[str, Any]:
        """
        Apply a result mask to the given job data.

        Parameters:
        - job_data (Union[Dict[str, Any], List[Dict[str, Any]]]): The input data to be masked.
          It can be either a single dictionary or a list of dictionaries.

        Returns:
        - Dict[str, Any]: A dictionary containing the masked result with the key 'result'.
        """
        return {
            "result": job_data,
        }

    async def run(self, source_input: type[warlock.model.Model]) -> Tuple[dict, StatusDTO]:
        """
        Run the job associated with the configuration.

        Args:
            source_input: The input data for the job.

        Returns:
            tuple: A tuple containing job_data and job_status.
        """
        logger.info(f"[RUNNING JOB] - Config ID: {self._config_id} - handler: {self._job_handler}")
        job_data, job_status = await self._module.Job(self._config, source_input).run()
        return self._apply_result_mask(job_data), job_status
