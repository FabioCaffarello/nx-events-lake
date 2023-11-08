import asyncio
import os
from typing import List

from config_loader.loader import fetch_configs
from source_watcher.consumer.consumer import EventConsumer
from pydotenv.loader import DotEnvLoader
from pylog.log import setup_logging
from pyrabbitmq.consumer import RabbitMQConsumer
from pysd.service_discovery import ServiceDiscovery, new_from_env

logger = setup_logging(__name__, log_level="DEBUG")

QUEUE_ACTIVE_JOBS = asyncio.Queue()
ENVIRONMENT = os.getenv("ENVIRONMENT")


async def create_consumers_channel(sd: ServiceDiscovery, service_name: str, context_env: str) -> List[asyncio.Task]:
    """
    Create consumers for processing data from various configurations.

    Args:
        sd (ServiceDiscovery): An instance of the ServiceDiscovery class.
        service_name (str): The name of the service.
        context_env (str): The context environment.

    Returns:
        List[asyncio.Task]: A list of asyncio tasks for processing data.

    """
    configs = await fetch_configs(service_name, context_env)
    rabbitmq_service = RabbitMQConsumer()
    await rabbitmq_service.connect()
    tasks = list()

    for _, context_configs in configs.items():
        for _, config in context_configs.items():
            logger.info(f"Creating consumer for config: {config.id}")
            tasks.append(
                asyncio.create_task(
                    EventConsumer(sd, rabbitmq_service, config, QUEUE_ACTIVE_JOBS).run()
                )
            )
    return tasks


async def main():
    """
    The main entry point of the service.

    This function initializes the necessary configurations, creates consumers for data processing, and runs the asyncio tasks.

    """
    envs = DotEnvLoader(environment=ENVIRONMENT)
    service_name = envs.get_variable("SERVICE_NAME")
    context_env = envs.get_variable("CONTEXT_ENV")
    logger.info(f"Service name: {service_name}")

    sd = new_from_env()
    tasks = await create_consumers_channel(sd, service_name, context_env)

    await asyncio.gather(*tasks)


if __name__ == "__main__":
    asyncio.run(main())
