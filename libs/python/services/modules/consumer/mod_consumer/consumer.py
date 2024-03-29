import asyncio
from typing import Union

from dto_config_handler.output import ConfigDTO
from pylog.log import setup_logging
from pyrabbitmq.consumer import RabbitMQConsumer
from pysd.service_discovery import ServiceDiscovery
import mod_debug.debug as debug

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

    Attributes:
        _config (ConfigDTO): The configuration data.
        _rabbitmq_service (RabbitMQConsumer): The RabbitMQ consumer service.
        _queue_active_jobs (asyncio.Queue): The asyncio queue for active jobs.
        _exchange_name (str): The name of the RabbitMQ exchange.
        _queue_name (str): The name of the RabbitMQ queue.
        _routing_key (str): The routing key for the RabbitMQ queue.

    Methods:
        __init__(self, sd: ServiceDiscovery, rabbitmq_service: RabbitMQConsumer, config: ConfigDTO, queue_active_jobs: asyncio.Queue) -> None:
            Initializes a Consumer instance with the provided parameters.

        async _run(self, controller: callable) -> None:
    """
    def __init__(self, sd: ServiceDiscovery, rabbitmq_service: RabbitMQConsumer, config: ConfigDTO, queue_active_jobs: asyncio.Queue, dbg: Union[debug.EnabledDebug, debug.DisabledDebug]):
        """
        Initializes a Consumer instance with the provided ServiceDiscovery, RabbitMQConsumer, configuration, and active jobs queue.

        Args:
            sd (ServiceDiscovery): An instance of the ServiceDiscovery class.
            rabbitmq_service (RabbitMQConsumer): An instance of the RabbitMQConsumer class.
            config (ConfigDTO): The configuration data.
            queue_active_jobs (asyncio.Queue): An asyncio queue for active jobs.

        Returns:
            None

        Raises:
            None
        """
        self._dbg = dbg
        self._config = config
        self._rabbitmq_service = rabbitmq_service
        self._queue_active_jobs = queue_active_jobs
        self._exchange_name = sd.services_rabbitmq_exchange()
        self._queue_name = _get_queue_name(config)
        self._routing_key = _get_routing_key(config)

    async def _run(self, controller: callable, job_handler: callable) -> None:
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
        await self._rabbitmq_service.listen(queue, controller(self._config, self._rabbitmq_service, self._queue_active_jobs, job_handler, self._dbg).run)


class EventConsumer(Consumer):
    """
    The EventConsumer class for processing event data.

    Args:
        sd (ServiceDiscovery): An instance of the ServiceDiscovery class.
        rabbitmq_service (RabbitMQConsumer): An instance of the RabbitMQConsumer class.
        config (ConfigDTO): The configuration data.
        queue_active_jobs (asyncio.Queue): An asyncio queue for active jobs.

    """

    def __init__(self, sd: ServiceDiscovery, rabbitmq_service: RabbitMQConsumer, config: ConfigDTO, queue_active_jobs: asyncio.Queue, dbg: Union[debug.EnabledDebug, debug.DisabledDebug]):
        super().__init__(sd, rabbitmq_service, config, queue_active_jobs, dbg)

    async def run(self, controller: callable, job_handler: callable) -> None:
        """
        Run the EventConsumer to process event data.

        This method triggers the processing of incoming event data using the specified controller.
        Args:
            controller (callable): The controller class responsible for processing data.

        """
        await self._run(controller, job_handler)
