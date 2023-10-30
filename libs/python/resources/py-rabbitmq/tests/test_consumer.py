import os
import time
import unittest
import asyncio
from pyrabbitmq.consumer import RabbitMQConsumer
from aio_pika.connection import Connection
from aio_pika.exceptions import AMQPConnectionError

def export_variables() -> None:
    os.environ["RABBITMQ_PORT_6572_TCP"] = "tcp://guest:guest@gateway_host:5672/"

class RabbitMQConsumerTests(unittest.TestCase):
    def setUp(self):
        export_variables()
        self.consumer = RabbitMQConsumer()
        self.queue_name = "test_queue"
        self.exchange_name = "test_exchange"
        self.routing_key = "test_routing_key"

    def tearDown(self):
        async def _close_connection():
            if self.consumer.connection is not None:
                await self.consumer.close_connection()
        self.run_async(_close_connection())

    def run_async(self, coro):
        self.loop = asyncio.get_event_loop()
        return self.loop.run_until_complete(coro)

    def test_connection(self):
        async def connection_test():
            await self.consumer.connect()
            self.assertIsInstance(self.consumer.connection, Connection)
        self.run_async(connection_test())

    def test_create_queue(self):
        async def create_queue_test():
            await self.consumer.connect()
            channel = await self.consumer.create_channel()
            queue = await self.consumer.create_queue(channel, self.queue_name, self.exchange_name, self.routing_key)
            self.assertIn(self.queue_name, queue.name)

        self.run_async(create_queue_test())

    def test_close_connection(self):
        async def close_connection_test():
            await self.consumer.connect()
            await self.consumer.close_connection()
            self.assertIsNone(self.consumer.connection)

        self.run_async(close_connection_test())

    def test_declare_exchange(self):
      async def declare_exchange_test():
          await self.consumer.connect()
          channel = await self.consumer.create_channel()
          await self.consumer.declare_exchange(channel, self.exchange_name)
          self.assertEqual(self.consumer.exchange.name, self.exchange_name)

      self.run_async(declare_exchange_test())

    def test_publish_message(self):
        async def publish_message_test():
            await self.consumer.connect()
            message = "Test message"
            await self.consumer.publish_message(self.exchange_name, self.routing_key, message)
            self.assertTrue(True)

        self.run_async(publish_message_test())

if __name__ == '__main__':
    unittest.main()
