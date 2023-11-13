import asyncio
import os
from typing import Any, List

from config_loader.loader import fetch_configs
from document_vectorizer.consumer.consumer import EventConsumer
from langchain.embeddings import (OllamaEmbeddings,
                                  SentenceTransformerEmbeddings)
from pydotenv.loader import DotEnvLoader
from pylog.log import setup_logging
from pyrabbitmq.consumer import RabbitMQConsumer
from pysd.service_discovery import ServiceDiscovery, new_from_env

logger = setup_logging(__name__, log_level="DEBUG")


QUEUE_ACTIVE_JOBS = asyncio.Queue()
ENVIRONMENT = os.getenv("ENVIRONMENT")

class BaseLogger:
    """
    A simple base logger class.

    This class is used for logging information.

    Methods:
        __init__(self) -> None:
            Initializes the BaseLogger instance.

    Attributes:
        info (callable): A callable method for logging information.

    """
    def __init__(self) -> None:
        self.info = print


async def create_consumers_channel(sd: ServiceDiscovery, service_name: str, context_env: str, embeddings, dimension) -> List[asyncio.Task]:
    """
    Create consumers for processing data from various configurations.

    Args:
        sd (ServiceDiscovery): An instance of the ServiceDiscovery class.
        service_name (str): The name of the service.
        context_env (str): The context environment.
        embeddings (Any): Embeddings for data processing.
        dimension (Any): Dimension for embeddings.

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
                    EventConsumer(sd, rabbitmq_service, config, QUEUE_ACTIVE_JOBS, embeddings, dimension).run()
                )
            )
    return tasks


async def main():
    """
    The main entry point of the service.

    This function initializes the necessary configurations, creates consumers for data processing,
    and runs the asyncio tasks.
    """
    envs = DotEnvLoader(environment=ENVIRONMENT)
    service_name = envs.get_variable("SERVICE_NAME")
    context_env = envs.get_variable("CONTEXT_ENV")
    logger.info(f"Service name: {service_name}")

    # Remapping for Langchain Neo4j integration
    os.environ["NEO4J_URL"] = envs.get_variable("NEO4J_URI")
    embedding_model_name = envs.get_variable("EMBEDDING_MODEL")
    ollama_base_url = envs.get_variable("OLLAMA_BASE_URL")


    embeddings, dimension = load_embedding_model(
        embedding_model_name,
        config={"ollama_base_url": ollama_base_url},
        logger=logger
    )

    sd = new_from_env()
    tasks = await create_consumers_channel(sd, service_name, context_env, embeddings, dimension)

    await asyncio.gather(*tasks)


def load_embedding_model(embedding_model_name: str, logger=BaseLogger(), config={}):
    """
    Load the embedding model based on the specified name.

    Args:
        embedding_model_name (str): The name of the embedding model.
        logger (BaseLogger, optional): The logger instance. Defaults to BaseLogger().
        config (dict, optional): Additional configuration for loading the model. Defaults to {}.

    Returns:
        Tuple: A tuple containing embeddings and dimension.
    """
    if embedding_model_name == "ollama":
        embeddings = OllamaEmbeddings(
            base_url=config["ollama_base_url"], model="llama2"
        )
        dimension = 4096
        logger.info("Embedding: Using Ollama")
    else:
        embeddings = SentenceTransformerEmbeddings(
            model_name="all-MiniLM-L6-v2", cache_folder="/embedding_model"
        )
        dimension = 384
        logger.info("Embedding: Using SentenceTransformer")
    return embeddings, dimension



if __name__ == "__main__":
    asyncio.run(main())
