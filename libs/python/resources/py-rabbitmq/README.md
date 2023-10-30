# py-rabbitmq

`py-rabbitmq` is a Python library that simplifies interaction with RabbitMQ using asynchronous Python libraries. It provides a base class for handling RabbitMQ connections and a consumer class for consuming messages from RabbitMQ queues.

## Installation

You can install `py-rabbitmq` using `nx`:

```sh
npx nx add <project> --name python-resources-py-rabbitmq --local
```

## Usage

The `RabbitMQConsumer` class is used for consuming messages from RabbitMQ queues. It extends the `BaseRabbitMQ` class. The `BaseRabbitMQ` class is a base class for interacting with RabbitMQ. It provides methods for connecting to RabbitMQ, creating channels, declaring exchanges, creating queues, and publishing messages.


## Examples

```python
import asyncio
from pylog.log import setup_logging
import aio_pika
from pyrabbitmq.consumer import RabbitMQConsumer

logger = setup_logging(__name__)

# Define a callback function to process incoming messages
async def process_message(message):
    body = message.body.decode()
    logger.info(f"Received message: {body}")

# Create a RabbitMQConsumer instance
async def main():
    consumer = RabbitMQConsumer()

    # Establish a connection to RabbitMQ
    await consumer.connect()

    # Create a channel
    channel = await consumer.create_channel()

    # Define the queue name and routing key
    queue_name = "my_queue"
    exchange_name = "my_exchange"
    routing_key = "my_routing_key"

    # Create a queue and bind it to the exchange
    queue = await consumer.create_queue(channel, queue_name, exchange_name, routing_key)

    # Start listening for messages
    await consumer.listen(queue, process_message)

if __name__ == "__main__":
    asyncio.run(main())

```
