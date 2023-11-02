from typing import List
import asyncio
import os
from pylog.log import setup_logging
from pydotenv.loader import DotEnvLoader
from config_loader.loader import fetch_configs
from file_downloader.consumer.consumer import EventConsumer
from pyrabbitmq.consumer import RabbitMQConsumer
from pysd.service_discovery import ServiceDiscovery, new_from_env

logger = setup_logging(__name__, log_level="DEBUG")

QUEUE_ACTIVE_JOBS = asyncio.Queue()
ENVIRONMENT = os.getenv("ENVIRONMENT")

async def create_consumers_channel(sd: ServiceDiscovery, service_name: str, context_env: str) -> List[asyncio.Task]:
    configs = await fetch_configs(service_name, context_env)
    rabbitmq_service = RabbitMQConsumer(url=sd.rabbitmq_endpoint())
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
    envs = DotEnvLoader(environment=ENVIRONMENT)
    service_name = envs.get_variable("SERVICE_NAME")
    context_env = envs.get_variable("CONTEXT_ENV")
    logger.info(f"Service name: {service_name}")

    sd = new_from_env()
    tasks = await create_consumers_channel(sd, service_name, context_env)

    await asyncio.gather(*tasks)


if __name__ == "__main__":
    asyncio.run(main())
