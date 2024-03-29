import asyncio
import json
import time
from datetime import datetime
from typing import Dict, Union

import warlock
from cli_schema_handler.client import async_py_schema_handler_client
from dto_config_handler.output import ConfigDTO
from dto_events_handler.input import ServiceFeedbackDTO
from dto_events_handler.shared import MetadataDTO, MetadataInputDTO
from dto_input_handler.output import InputDTO
from pylog.log import setup_logging
import pywarlock.serializer
from pyrabbitmq.consumer import RabbitMQConsumer
import mod_debug.debug as debug
from pyserializer.serializer import serialize_to_dataclass, serialize_to_json

logger = setup_logging(__name__)

_REPOSITORY_SCHEMA_TYPE = "service-input"


class Controller:
    """
    Base class for handling event data processing.

    Args:
        config (ConfigDTO): The configuration data.
        queue_active_jobs (asyncio.Queue): An asyncio queue for active jobs.

    Attributes:
        _config (ConfigDTO): The configuration data.
        _config_id (str): The ID associated with the configuration.
        _service_name (str): The service name from the configuration.
        _source_name (str): The source name from the configuration.
        _context_env (str): The context environment from the configuration.
        _repository_schema_type (str): The repository schema type for service input.
        _queue_active_jobs (asyncio.Queue): An asyncio queue for active jobs.
        _active (bool): The activation status based on the configuration.
        _schema_handler_client: The schema handler client for retrieving JSON schemas.
        _input_body_dto: The input data DTO.

    Methods:
        __init__(self, config: ConfigDTO, queue_active_jobs: asyncio.Queue) -> None:
            Initializes a Controller instance with the provided configuration and active jobs queue.

        _should_cotroller_active(self) -> bool:
            Check if the controller should be active based on the configuration.

        async _get_event_parser(self) -> Dict[str, any]:
            Get the event parser JSON schema for data processing.

        async _parse_event(self, message: str) -> type[warlock.model.Model]:
            Parse the incoming event message and transform it into the appropriate data format.

        _get_metadata(self) -> MetadataDTO:
            Generate metadata information for the processed event data.

        async job_dispatcher(self, event_input, job_handler) -> ServiceFeedbackDTO:
            Dispatch a job to process the event input data and collect the results.
    """
    def __init__(self, config: ConfigDTO, queue_active_jobs: asyncio.Queue, dbg: Union[debug.EnabledDebug, debug.DisabledDebug]):
        """
        Initializes a Controller instance with the provided configuration and active jobs queue.

        Args:
            config (ConfigDTO): The configuration data.
            queue_active_jobs (asyncio.Queue): An asyncio queue for active jobs.

        Returns:
            None
        """
        self._dbg = dbg
        self._config = config
        self._config_id = config.id
        self._service_name = config.service
        self._source_name = config.source
        self._context_env = config.context
        self._repository_schema_type = _REPOSITORY_SCHEMA_TYPE
        self._queue_active_jobs = queue_active_jobs
        self._active = config.active
        self._schema_handler_client = async_py_schema_handler_client()
        self._input_body_dto = None

    def _should_cotroller_active(self) -> bool:
        """
        Check if the controller should be active based on the configuration.

        Returns:
            bool: True if the controller is active, False otherwise.

        """
        if self._active:
            return True
        return False

    async def _get_event_parser(self) -> Dict[str, any]:
        """
        Get the event parser JSON schema for data processing.

        Returns:
            Dict[str, any]: The JSON schema for data processing.

        """
        self.schema_input = await self._schema_handler_client.list_one_schema_by_service_n_source_n_context_n_schema_type(
            context=self._context_env,
            service_name=self._service_name,
            source_name=self._source_name,
            schema_type=self._repository_schema_type
        )
        json_schema = self.schema_input.json_schema
        return json_schema

    async def _parse_event(self, message: str) -> type[warlock.model.Model]:
        """
        Parse the incoming event message and transform it into the appropriate data format.

        Args:
            message (str): The incoming event message.

        Returns:
            object: The parsed event data in the required data format.

        Raises:
            ValueError: If the message body cannot be parsed.

        """
        message_body = message.body.decode()
        event_parser_class = await self._get_event_parser()
        try:
            input_body = json.loads(message_body)
            self._input_body_dto = serialize_to_dataclass(input_body, InputDTO)

            input_data = self._input_body_dto.data
            return pywarlock.serializer.serialize_to_dataclass(event_parser_class, input_data)
        except json.JSONDecodeError as e:
            logger.error(f"Failed to parse message body: {e}")
            raise ValueError("Invalid message body")

    def _get_metadata(self) -> MetadataDTO:
        """
        Generate metadata information for the processed event data.

        Returns:
            MetadataDTO: Metadata information for the event data.

        """
        return MetadataDTO(
            input=MetadataInputDTO(
                id=self._input_body_dto.id,
                data=self._input_body_dto.data,
                processing_id=self._input_body_dto.metadata["processing_id"],
                processing_timestamp=self._input_body_dto.metadata["processing_timestamp"],
                input_schema_id=self.schema_input.schema_id
            ),
            context=self._config.context,
            service=self._config.service,
            source=self._config.source,
            processing_timestamp=datetime.now().strftime("%Y-%m-%dT%H:%M:%SZ"),
            job_frequency=self._config.frequency,
        )

    async def job_dispatcher(self, event_input, job_handler: callable) -> ServiceFeedbackDTO:
        """
        Dispatch a job to process the event input data and collect the results.

        Args:
            event_input: The input data for the job.
            job_handler (callable): The job handler class responsible for processing the data.

        Returns:
            ServiceFeedbackDTO: Feedback and result information from the job processing.

        """
        await self._queue_active_jobs.put(1)
        job_data, status_data = await job_handler(self._config, self._dbg).run(event_input)
        return ServiceFeedbackDTO(
            data=job_data,
            metadata=self._get_metadata(),
            status=status_data,
        )


class EventController(Controller):
    """
    EventController class for processing event data.

    Args:
        config (ConfigDTO): The configuration data.
        rabbitmq_service (RabbitMQConsumer): An instance of the RabbitMQConsumer class.
        queue_active_jobs (asyncio.Queue): An asyncio queue for active jobs.

    Methods:
        __init__(self, config: ConfigDTO, rabbitmq_service: RabbitMQConsumer, queue_active_jobs: asyncio.Queue) -> None:
            Initializes an EventController instance with the provided configuration, RabbitMQ service, and active jobs queue.

        async run(self, message) -> None:
            Run the EventController to process event data.

            This method initiates the processing of incoming event data using the specified controller logic.

            Args:
                message: The incoming event message.
                job_handler (callable): The job handler class responsible for processing the data.
    """
    def __init__(self, config: ConfigDTO, rabbitmq_service: RabbitMQConsumer, queue_active_jobs: asyncio.Queue, job_handler: callable, dbg: Union[debug.EnabledDebug, debug.DisabledDebug]):
        """
        Initializes an EventController instance with the provided configuration, RabbitMQ service, and active jobs queue.

        Args:
            config (ConfigDTO): The configuration data.
            rabbitmq_service (RabbitMQConsumer): An instance of the RabbitMQConsumer class.
            queue_active_jobs (asyncio.Queue): An asyncio queue for active jobs.

        Returns:
            None
        """
        self._rabbitmq_service = rabbitmq_service
        self._job_handler = job_handler
        super().__init__(config, queue_active_jobs, dbg)

    async def run(self, message) -> None:
        """
        Run the EventController to process event data.

        This method initiates the processing of incoming event data using the specified controller logic.

        Args:
            message: The incoming event message.
            job_handler (callable): The job handler class responsible for processing the data.

        Returns:
            None
        """
        if not self._should_cotroller_active():
            logger.info(f"Controller for config_id {self._config_id} is not active")
            return

        await self._rabbitmq_service.publish_message(
            "services",
            "input-processing",
            json.dumps(json.loads(message.body.decode()))
        )

        event_input = await self._parse_event(message)
        job_result = await self.job_dispatcher(event_input, self._job_handler)
        await self._queue_active_jobs.get()
        output = serialize_to_json(job_result)
        logger.info(f"sleeping for 5 seconds...")
        time.sleep(5)
        logger.info(f"Output: {output}")
        await self._rabbitmq_service.publish_message(
            "services",
            "feedback",
            output
        )
        await message.ack()
        logger.info("Published message to service")
