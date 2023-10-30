import asyncio
import aio_pika
import urllib.parse
from pylog.log import setup_logging
from pysd.service_discovery import new_from_env

logger = setup_logging(__name__)

class BaseRabbitMQ:
    """A base class for interacting with RabbitMQ.

    Args:
        None

    Attributes:
        _sd: The service discovery instance.
        url: The RabbitMQ connection URL.
        connection: The RabbitMQ connection.
        exchange: The RabbitMQ exchange.

    Methods:
        __init__: Initialize the BaseRabbitMQ object.
        _connect: Connect to RabbitMQ.
        connect: Retry connecting to RabbitMQ until successful.
        on_connection_error: Handle connection errors.
        create_channel: Create a new channel in the RabbitMQ connection.
        declare_exchange: Declare a RabbitMQ exchange.
        create_queue: Create a RabbitMQ queue and bind it to an exchange.
        close_connection: Close the RabbitMQ connection.
        publish_message: Publish a message to a RabbitMQ exchange.

    """
    def __init__(self) -> None:
        self._sd = new_from_env()
        self.url = self._sd.rabbitmq_endpoint()
        self.connection = None
        self.exchange = None

    async def _connect(self) -> None:
        """Connect to RabbitMQ.

        Args:
            None

        Returns:
            None

        """
        parsed_url = urllib.parse.urlparse(self.url)
        self.connection = await aio_pika.connect(
            host=parsed_url.hostname,
            port=parsed_url.port,
            login=parsed_url.username,
            password=parsed_url.password,
        )

    async def connect(self) -> None:
        """Retry connecting to RabbitMQ until successful.

        Args:
            None

        Returns:
            None

        """
        while True:
            try:
                await self._connect()
                break
            except Exception as err:
                logger.error('[CONNECTION] - Could not connect to RabbitMQ, retrying in 2 seconds...')
                self.on_connection_error(err)
                await asyncio.sleep(2)

    def on_connection_error(self, error: Exception) -> None:
        """Handle connection errors.

        Args:
            error (Exception): The connection error.

        Returns:
            None

        """
        logger.error(f"Connection error: {error}")
        logger.error(f"Connection parameters: {self.url}")

    async def create_channel(self) -> aio_pika.Channel:
        channel = await self.connection.channel()
        await channel.set_qos(prefetch_count=1)
        return channel

    async def declare_exchange(self, channel: aio_pika.Channel, exchange_name: str) -> None:
        """Declare a RabbitMQ exchange.

        Args:
            channel (aio_pika.Channel): The channel to declare the exchange on.
            exchange_name (str): The name of the exchange.

        Returns:
            None

        """
        self.exchange = await channel.declare_exchange(
            exchange_name, aio_pika.ExchangeType.TOPIC, durable=True
        )

    async def create_queue(self, channel: aio_pika.Channel, queue_name: str, exchange_name: str, routing_key: str) -> aio_pika.Queue:
        """Create a RabbitMQ queue and bind it to an exchange.

        Args:
            channel (aio_pika.Channel): The channel to create the queue on.
            queue_name (str): The name of the queue.
            exchange_name (str): The name of the exchange to bind the queue to.
            routing_key (str): The routing key to use for binding.

        Returns:
            aio_pika.Queue: The created queue.

        """
        await self.declare_exchange(channel, exchange_name)
        queue = await channel.declare_queue(queue_name, durable=True)
        await queue.bind(self.exchange, routing_key)

        return queue

    async def close_connection(self) -> None:
        """Close the RabbitMQ connection.

        Args:
            None

        Returns:
            None

        """
        if self.connection is not None:
            await self.connection.close()
            self.connection = None

    async def publish_message(self, exchange_name: str, routing_key: str, message: str) -> None:
        """Publish a message to a RabbitMQ exchange.

        Args:
            exchange_name (str): The name of the exchange to publish to.
            routing_key (str): The routing key for the message.
            message (str): The message to publish.

        Returns:
            None

        """
        try:
            await self.exchange.publish(
                aio_pika.Message(
                    body=message.encode(),
                    delivery_mode=aio_pika.DeliveryMode.PERSISTENT,
                ),
                routing_key=routing_key,
            )
            logger.info(f"Published message to exchange '{exchange_name}' with routing key '{routing_key}'")
        except Exception as e:
            logger.error(f"Error while publishing message: {e}")
