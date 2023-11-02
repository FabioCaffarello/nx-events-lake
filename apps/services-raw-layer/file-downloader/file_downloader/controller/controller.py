import asyncio
import json
import time
from datetime import datetime

import warlock
from dto_config_handler.output import ConfigDTO
from dto_events_handler.input import ServiceFeedbackDTO
from dto_events_handler.shared import MetadataDTO, MetadataInputDTO
from dto_input_handler.output import InputDTO
from file_downloader.jobs.job_handler import JobHandler
from pylog.log import setup_logging
from pyrabbitmq.consumer import RabbitMQConsumer
from cli_schema_handler.client import async_py_schema_handler_client
from pyserializer.serializer import serialize_to_json, serialize_to_dataclass

logger = setup_logging(__name__)


_REPOSITORY_SCHEMA_TYPE = "service-input"


class Controller:
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

    def _should_cotroller_active(self):
        if self._active :
            return True
        return False

    async def _get_event_parser(self):
        self.schema_input = await self._schema_handler_client.list_one_schema_by_service_n_source_n_context_n_schema_type(
            context=self._context_env,
            service_name=self._service_name,
            source_name=self._source_name,
            schema_type=self._repository_schema_type
        )
        json_schema = self.schema_input.json_schema
        return json_schema

    async def _parse_event(self, message: str):
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

    def _get_metadata(self, target_endpoint: str):
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
        await self._queue_active_jobs.put(1)
        job_data, status_data, target_endpoint = JobHandler(self._config).run(event_input)
        return ServiceFeedbackDTO(
            data=job_data,
            metadata=self._get_metadata(target_endpoint),
            status=status_data,
        )

class EventController(Controller):
    def __init__(self, config: ConfigDTO, rabbitmq_service: RabbitMQConsumer, queue_active_jobs: asyncio.Queue):
        self._rabbitmq_service = rabbitmq_service
        super().__init__(config, queue_active_jobs)

    async def run(self, message):
        if not self._should_cotroller_active():
            logger.info(f"Controller for config_id {self._config_id} is not active")
            return

        event_input = await self._parse_event(message)
        job_result = await self.job_dispatcher(event_input)
        await self._queue_active_jobs.get()
        output = serialize_to_json(job_result)
        logger.info(f"sleeping for 20 seconds...")
        time.sleep(5)
        logger.info(f"Output: {output}")
        await self._rabbitmq_service.publish_message(
                "services",
                "feedback",
                output
            )
        await message.ack()
        logger.info("Published message to service")
