import importlib

from dto_config_handler.output import ConfigDTO
from pylog.log import setup_logging

logger = setup_logging(__name__)

class JobHandler:
    def __init__(self, config: ConfigDTO):
        self._config = config
        self._job_handler = config.service_parameters["job_handler"]
        self._config_id = config.id
        self._module = self._import_job_handler_as_module()

    def _import_job_handler_as_module(self):
        return importlib.import_module(f"jobs.handlers.{self._job_handler}.job")

    def run(self, source_input):
        logger.info(f"[RUNNING JOB] - Config ID: {self._config_id} - handler: {self._job_handler}")
        job_data, job_status, target_endpoint = self._module.Job(self._config, source_input).run()
        return job_data, job_status, target_endpoint


