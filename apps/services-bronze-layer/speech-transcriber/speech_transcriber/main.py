import asyncio
import os
from typing import List
import torch
from transformers import AutoModelForSpeechSeq2Seq, AutoProcessor, pipeline
from config_loader.loader import fetch_configs
from speech_transcriber.consumer.consumer import EventConsumer
from pydotenv.loader import DotEnvLoader
from pylog.log import setup_logging
from optimum.bettertransformer import BetterTransformer
from pyrabbitmq.consumer import RabbitMQConsumer
from pysd.service_discovery import ServiceDiscovery, new_from_env

logger = setup_logging(__name__, log_level="DEBUG")

QUEUE_ACTIVE_JOBS = asyncio.Queue()
ENVIRONMENT = os.getenv("ENVIRONMENT")

async def create_consumers_channel(sd: ServiceDiscovery, service_name: str, context_env: str, transcription_pipeline: pipeline) -> List[asyncio.Task]:
    """
    Create consumers for processing data from various configurations.

    Args:
        sd (ServiceDiscovery): An instance of the ServiceDiscovery class.
        service_name (str): The name of the service.
        context_env (str): The context environment.
        transcription_pipeline (pipeline): The pipeline for automatic speech recognition.

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
                    EventConsumer(sd, rabbitmq_service, config, QUEUE_ACTIVE_JOBS, transcription_pipeline).run()
                )
            )
    return tasks


def setup_model(model_id, device, torch_dtype):
    """
    Set up the speech-to-text model.

    Args:
        model_id (str): The identifier for the speech-to-text model.
        device (str): The device to use (e.g., "cpu" or "cuda").
        torch_dtype: The torch data type (e.g., torch.float16 or torch.float32).

    Returns:
        BetterTransformer: The configured speech-to-text model.
    """
    model = AutoModelForSpeechSeq2Seq.from_pretrained(
        model_id, torch_dtype=torch_dtype, low_cpu_mem_usage=True, use_safetensors=True
    )
    model.to(device)
    return model.to_bettertransformer()

def create_transcription_pipeline(model, processor, torch_dtype, device):
    """
    Create a speech transcription pipeline.

    Args:
        model: The speech-to-text model.
        processor: The processor for the model.
        torch_dtype: The torch data type (e.g., torch.float16 or torch.float32).
        device (str): The device to use (e.g., "cpu" or "cuda").

    Returns:
        pipeline: The configured speech transcription pipeline.
    """
    return pipeline(
        "automatic-speech-recognition",
        model=model,
        tokenizer=processor.tokenizer,
        feature_extractor=processor.feature_extractor,
        max_new_tokens=128,
        chunk_length_s=15, #long form transcription
        batch_size=16,
        torch_dtype=torch_dtype,
        device=device,
    )

async def main():
    """
    The main entry point of the service.

    This function initializes the necessary configurations, creates consumers for data processing, and runs the asyncio tasks.
    """
    envs = DotEnvLoader(environment=ENVIRONMENT)
    service_name = envs.get_variable("SERVICE_NAME")
    context_env = envs.get_variable("CONTEXT_ENV")
    logger.info(f"Service name: {service_name}")


    device = "cuda" if torch.cuda.is_available() else "cpu"
    model_id = "distil-whisper/distil-medium.en"
    torch_dtype = torch.float16 if torch.cuda.is_available() else torch.float32
    model = setup_model(model_id, device, torch_dtype)
    processor = AutoProcessor.from_pretrained(model_id)
    transcription_pipeline = create_transcription_pipeline(model, processor, torch_dtype, device)

    sd = new_from_env()
    tasks = await create_consumers_channel(sd, service_name, context_env, transcription_pipeline)

    await asyncio.gather(*tasks)


if __name__ == "__main__":
    asyncio.run(main())
