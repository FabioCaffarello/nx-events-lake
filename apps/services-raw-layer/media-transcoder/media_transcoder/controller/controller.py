import asyncio
import json
import time
from datetime import datetime
from typing import Dict

import warlock
from cli_schema_handler.client import async_py_schema_handler_client
from dto_config_handler.output import ConfigDTO
from dto_events_handler.input import ServiceFeedbackDTO
from dto_events_handler.shared import MetadataDTO, MetadataInputDTO
from dto_input_handler.output import InputDTO
from media_transcoder.jobs.job_handler import JobHandler
from pylog.log import setup_logging
from pyrabbitmq.consumer import RabbitMQConsumer
from pyserializer.serializer import serialize_to_dataclass, serialize_to_json

logger = setup_logging(__name__)

_REPOSITORY_SCHEMA_TYPE = "service-input"

class Controller:
    """
    Base class for handling event data processing.

    Args:
        config (ConfigDTO): The configuration data.
        queue_active_jobs (asyncio.Queue): An asyncio queue for active jobs.

    """
    def __init__(self, config: ConfigDTO, queue_active_jobs: asyncio.Queue):
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
            Input_dataclass = warlock.model_factory(event_parser_class)
            return Input_dataclass(**input_data)
        except json.JSONDecodeError as e:
            logger.error(f"Failed to parse message body: {e}")
            raise ValueError("Invalid message body")

    def _get_metadata(self, target_endpoint: str) -> MetadataDTO:
        """
        Generate metadata information for the processed event data.

        Args:
            target_endpoint (str): The target endpoint for the event data.

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
            job_config_id=self._config.config_id,
        )

    async def job_dispatcher(self, event_input) -> ServiceFeedbackDTO:
        """
        Dispatch a job to process the event input data and collect the results.

        Args:
            event_input: The input data for the job.

        Returns:
            ServiceFeedbackDTO: Feedback and result information from the job processing.

        """
        await self._queue_active_jobs.put(1)
        job_data, status_data, target_endpoint = JobHandler(self._config).run(event_input)
        return ServiceFeedbackDTO(
            data=job_data,
            metadata=self._get_metadata(target_endpoint),
            status=status_data,
        )

class EventController(Controller):
    """
    EventController class for processing event data.

    Args:
        config (ConfigDTO): The configuration data.
        rabbitmq_service (RabbitMQConsumer): An instance of the RabbitMQConsumer class.
        queue_active_jobs (asyncio.Queue): An asyncio queue for active jobs.

    """
    def __init__(self, config: ConfigDTO, rabbitmq_service: RabbitMQConsumer, queue_active_jobs: asyncio.Queue) -> None:
        self._rabbitmq_service = rabbitmq_service
        super().__init__(config, queue_active_jobs)

    async def run(self, message) -> None:
        """
        Run the EventController to process event data.

        This method initiates the processing of incoming event data using the specified controller logic.

        Args:
            message: The incoming event message.

        """
        logger.info(f"Processing message: {message}")
        if not self._should_cotroller_active():
            logger.warning(f"Controller for config_id {self._config_id} is not active")
            return

        await self._rabbitmq_service.publish_message(
            "services",
            "input-processing",
            json.dumps(json.loads(message.body.decode()))
        )

        event_input = await self._parse_event(message)
        job_result = await self.job_dispatcher(event_input)
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
        await self._queue_active_jobs.get()
        logger.info("Published message to service")
