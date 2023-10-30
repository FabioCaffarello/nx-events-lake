from pylog.log import setup_logging
import aio_pika
from pyrabbitmq.base import BaseRabbitMQ

logger = setup_logging(__name__)


class RabbitMQConsumer(BaseRabbitMQ):
    """A RabbitMQ consumer class that extends BaseRabbitMQ.

    Args:
        None

    Attributes:
        None

    Methods:
        __init__: Initialize the RabbitMQConsumer object.
        listen: Asynchronously listen to a queue and call the callback function on message arrival.

    """
    def __init__(self):
        super().__init__()

    async def listen(self, queue: aio_pika.Queue, callback: callable) -> None:
        """Listen to a RabbitMQ queue and call the specified callback on message arrival.

        Args:
            queue (aio_pika.Queue): The queue to listen to.
            callback (callable): The callback function to execute on message arrival.

        Returns:
            None

        """
        async with queue.iterator() as queue_iter:
            message: aio_pika.AbstractIncomingMessage
            async for message in queue_iter:
                await callback(message)
