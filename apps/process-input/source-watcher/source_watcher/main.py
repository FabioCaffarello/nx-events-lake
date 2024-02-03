import asyncio
import os
import time
from typing import List, Union

from config_loader.loader import fetch_configs
from mod_consumer.consumer import EventConsumer
from mod_controller.controller import EventController
from mod_jobs.job_handler import JobHandler
from pydotenv.loader import DotEnvLoader
from pylog.log import setup_logging
from pyrabbitmq.consumer import RabbitMQConsumer
from pysd.service_discovery import ServiceDiscovery, new_from_env
import pyargparse.argsparser as pyargparser
import mod_debug.debug as debug

logger = setup_logging(__name__, log_level="DEBUG")

QUEUE_ACTIVE_JOBS = asyncio.Queue()
ENVIRONMENT = os.getenv("ENVIRONMENT")


async def create_consumers_channel(sd: ServiceDiscovery, service_name: str, context_env: str, dbg: Union[debug.EnabledDebug, debug.DisabledDebug]) -> List[asyncio.Task]:
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
                    EventConsumer(sd, rabbitmq_service, config, QUEUE_ACTIVE_JOBS, dbg).run(EventController, JobHandler)
                )
            )
    return tasks


def _parse_args():
    parser = pyargparser.new("source-watcher args parser")
    return parser.parse_args()


def setup_debug_storage(args):
    if args.enable_debug_storage:
        os.environ["DEBUG_STORAGE_ENABLED"] = "True"
        os.environ["DEBUG_STORAGE_DIR"] = args.debug_storage_dir


def cast_string_to_boolean(input_string):
    if input_string.lower() == "true":
        return True
    else:
        return False


async def main():
    """
    The main entry point of the service.

    This function initializes the necessary configurations, creates consumers for data processing, and runs the asyncio tasks.

    """
    envs = DotEnvLoader(environment=ENVIRONMENT)
    service_name = envs.get_variable("SERVICE_NAME")
    context_env = envs.get_variable("CONTEXT_ENV")
    logger.info(f"Service name: {service_name}")

    args = _parse_args()
    setup_debug_storage(args)

    dbg = debug.new(cast_string_to_boolean(envs.get_variable("DEBUG_STORAGE_ENABLED")), envs.get_variable("DEBUG_STORAGE_DIR"))

    sd = new_from_env()
    tasks = await create_consumers_channel(sd, service_name, context_env, dbg)

    await asyncio.gather(*tasks)


if __name__ == "__main__":
    try:
        asyncio.run(main())
    except:
        time.sleep(30)
        asyncio.run(main())
