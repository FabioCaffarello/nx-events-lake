import asyncio

from controller.controller import EventController
from dto_config_handler.output import ConfigDTO
from pylog.log import setup_logging
from pyrabbitmq.consumer import RabbitMQConsumer
from pysd.service_discovery import ServiceDiscovery

logger = setup_logging(__name__)


def _get_queue_name(config: ConfigDTO):
    return "{context}.{service}.inputs.{source}".format(
        context=config.context,
        service=config.service,
        source=config.source,
    )


def _get_routing_key(config: ConfigDTO):
    return "{service}.inputs.{source}".format(
        service=config.service,
        source=config.source,
    )


class Consumer:
    def __init__(self, sd: ServiceDiscovery, rabbitmq_service: RabbitMQConsumer, config: ConfigDTO, queue_active_jobs: asyncio.Queue):
        self._config = config
        self._rabbitmq_service = rabbitmq_service
        self._queue_active_jobs = queue_active_jobs
        self._exchange_name = sd.services_rabbitmq_exchange()
        self._queue_name = _get_queue_name(config)
        self._routing_key = _get_routing_key(config)

    async def _run(self, controller):
        channel = await self._rabbitmq_service.create_channel()
        queue = await self._rabbitmq_service.create_queue(
            channel,
            self._queue_name,
            self._exchange_name,
            self._routing_key
        )
        await self._rabbitmq_service.listen(queue, controller(self._config, self._rabbitmq_service, self._queue_active_jobs).run)


class EventConsumer(Consumer):
    def __init__(self, sd: ServiceDiscovery, rabbitmq_service: RabbitMQConsumer, config: ConfigDTO, queue_active_jobs: asyncio.Queue):
        super().__init__(sd, rabbitmq_service, config, queue_active_jobs)

    async def run(self):
        await self._run(EventController)

