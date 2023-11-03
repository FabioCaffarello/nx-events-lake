import asyncio

from dto_config_handler.output import ConfigDTO
from source_watcher.controller.controller import EventController
from pylog.log import setup_logging
from pyrabbitmq.consumer import RabbitMQConsumer
from pysd.service_discovery import ServiceDiscovery

logger = setup_logging(__name__)


def _get_queue_name(config: ConfigDTO) -> str:
    """
    Get the queue name for a specific configuration.

    Args:
        config (ConfigDTO): The configuration data.

    Returns:
        str: The generated queue name.
    """
    return "{context}.{service}.inputs.{source}".format(
        context=config.context,
        service=config.service,
        source=config.source,
    )


def _get_routing_key(config: ConfigDTO) -> str:
    """
    Get the routing key for a specific configuration.

    Args:
        config (ConfigDTO): The configuration data.

    Returns:
        str: The generated routing key.
    """
    return "{context}.{service}.inputs.{source}".format(
        context=config.context,
        service=config.service,
        source=config.source,
    )



class Consumer:
    """
    The base class for creating data consumers.

    Args:
        sd (ServiceDiscovery): An instance of the ServiceDiscovery class.
        rabbitmq_service (RabbitMQConsumer): An instance of the RabbitMQConsumer class.
        config (ConfigDTO): The configuration data.
        queue_active_jobs (asyncio.Queue): An asyncio queue for active jobs.

    """

    def __init__(self, sd: ServiceDiscovery, rabbitmq_service: RabbitMQConsumer, config: ConfigDTO, queue_active_jobs: asyncio.Queue):
        self._config = config
        self._rabbitmq_service = rabbitmq_service
        self._queue_active_jobs = queue_active_jobs
        self._exchange_name = sd.services_rabbitmq_exchange()
        self._queue_name = _get_queue_name(config)
        self._routing_key = _get_routing_key(config)

    async def _run(self, controller: callable) -> None:
        """
        Run the consumer with the specified controller.

        Args:
            controller (callable): The controller class responsible for processing data.

        """
        channel = await self._rabbitmq_service.create_channel()
        queue = await self._rabbitmq_service.create_queue(
            channel,
            self._queue_name,
            self._exchange_name,
            self._routing_key
        )
        await self._rabbitmq_service.listen(queue, controller(self._config, self._rabbitmq_service, self._queue_active_jobs).run)


class EventConsumer(Consumer):
    """
    The EventConsumer class for processing event data.

    Args:
        sd (ServiceDiscovery): An instance of the ServiceDiscovery class.
        rabbitmq_service (RabbitMQConsumer): An instance of the RabbitMQConsumer class.
        config (ConfigDTO): The configuration data.
        queue_active_jobs (asyncio.Queue): An asyncio queue for active jobs.

    """

    def __init__(self, sd: ServiceDiscovery, rabbitmq_service: RabbitMQConsumer, config: ConfigDTO, queue_active_jobs: asyncio.Queue):
        super().__init__(sd, rabbitmq_service, config, queue_active_jobs)

    async def run(self) -> None:
        """
        Run the EventConsumer to process event data.

        This method triggers the processing of incoming event data using the specified controller.

        """
        await self._run(EventController)
