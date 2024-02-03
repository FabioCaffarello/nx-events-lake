import asyncio
import json
import os
import unittest
import uuid
from datetime import datetime
from pathlib import Path
from typing import Dict, List

from cli_config_handler.client import async_py_config_handler_client
from cli_file_catalog_handler.client import async_py_file_catalog_handler_client
from cli_jobs_handler.client import async_py_jobs_handler_client
from cli_schema_handler.client import async_py_schema_handler_client
from dto_config_handler.output import ConfigDTO
from dto_input_handler.output import InputDTO
from dto_input_handler.shared import MetadataDTO, StatusDTO
from pydotenv.loader import DotEnvLoader
from pylog.log import setup_logging
from pyminio.client import MinioClient, minio_client
from pymongodb.client import drop_database
from pyrabbitmq.consumer import RabbitMQConsumer
from pyserializer.serializer import serialize_to_dataclass, serialize_to_json
from subprocessd.subprocessd import SubprocessD, SubprocessDAsync

logger = setup_logging(__name__)


def _clean_storage():
    drop_database("schemas")
    drop_database("configs")
    drop_database("jobs-configs")
    drop_database("file-catalog")

def _get_bucket_name(context: str, service: str, source: str):
    if service == "source-watcher":
        return "process-input-{context}-source-{source}".format(
                context=context,
                source=source,
            )
    elif service == "file-downloader":
        return "landing-{context}-source-{source}".format(
                context=context,
                source=source,
            )
    elif service == "spark-batch-bronze":
        return "raw-{context}-source-{source}".format(
                context=context,
                source=source,
            )
    else:
        return ""

def generate_input(input_data: Dict[str, any], context_env: str, service: str, source: str) -> InputDTO:
    return InputDTO(
        id=str(uuid.uuid4().hex),
        status=StatusDTO(
            code=0,
            detail=""
        ),
        metadata=MetadataDTO(
            processing_id=str(uuid.uuid4().hex),
            processing_timestamp=datetime.now().isoformat(),
            service=service,
            source=source,
            context=context_env
        ),
        data=input_data
    )

_JSON_EXTENTION_REGEX_PATTERN = "*.json"
ENVIRONMENT = os.getenv("ENVIRONMENT")


MAPPING_HANDLERS = {
    "configs": async_py_config_handler_client,
    "schemas": async_py_schema_handler_client,
    "job-params": async_py_jobs_handler_client,
    "file-catalog": async_py_file_catalog_handler_client
}

class BaseTestsFixture(unittest.IsolatedAsyncioTestCase):
    async def asyncSetUp(self):
        self.envs = DotEnvLoader(environment=ENVIRONMENT)
        self.rabbitmq = RabbitMQConsumer()
        await self.rabbitmq.connect()
        _clean_storage()
        self.__all_configs = []
        self.__mapping_handlers_client = MAPPING_HANDLERS
        self._service_name = self.envs.get_variable("SERVICE_NAME")
        await self.set_configs_by_service()
        await self.purge_all_queues()
        self.create_bucket()
        self.queue = asyncio.Queue()
        args = self._get_service_process_args()
        args.append("--enable-debug-storage")
        args.append("--debug-storage-dir")
        args.append(self.get_debug_storage_dir())
        self._subprocessd = SubprocessDAsync(args, self.get_debug_storage_dir())
        await self._subprocessd.start()
        # self._subprocessd = SubprocessD(args, self.get_debug_storage_dir())

    async def asyncTearDown(self):
        await self._subprocessd.stop()
        _clean_storage()
        await self.purge_all_queues()
        await self.rabbitmq.close_connection()

    async def set_configs_by_service(self):
        await self.load_config_and_post_to_api("configs")
        await self.load_config_and_post_to_api("schemas")
        if self._service_name == "source-watcher" or self._service_name == "file-downloader":
            await self.load_config_and_post_to_api("job-params", "http-gateway")
        if self._service_name == "spark-batch-bronze":
            await self.load_config_and_post_to_api("file-catalog")

    def _get_file_posix_path(self, config_type: str):
        return Path("/app/tests/.configs").joinpath(config_type)

    def _find_local_config_files(self, config_type: str) -> List[Path]:
        configs_path = self._get_file_posix_path(config_type)
        logger.info(f"Looking for {config_type} files in {configs_path}")
        return list(configs_path.glob(_JSON_EXTENTION_REGEX_PATTERN))

    def _read_config_file(self, config_path: Path):
        with open(config_path, "r") as config_file:
            return json.load(config_file)

    def _get_queue_name(self, context: str, service: str, source: str):
        return "{context}.{service}.inputs.{source}".format(
            context=context,
            service=service,
            source=source,
        )

    async def _purge_queue(self, context_env: str, service: str, source: str):
        queue_name = self._get_queue_name(context_env, service, source)
        logger.info(f"Purging queue {queue_name}")
        await self.rabbitmq.purge_queue(queue_name)

    async def purge_all_queues(self):
        for config in self.__all_configs:
            await self._purge_queue(config.context, config.service, config.source)

    async def load_config_and_post_to_api(self, config_type: str, job_params_type=""):
        all_configs = self._find_local_config_files(config_type)
        all_configs_ids = []
        logger.info(f"Found {len(all_configs)} {config_type} files")
        for config in all_configs:
            config_data = self._read_config_file(config)
            logger.info(f"Creating {config_type} with data: {config_data}")
            if config_type == "configs":
                all_configs_ids.append(serialize_to_dataclass(config_data, ConfigDTO))
            if config_type == "job-params":
                await self.__mapping_handlers_client[config_type]().create(config_data, job_params_type)
            else:
                await self.__mapping_handlers_client[config_type]().create(config_data)
        if config_type == "configs":
            self.__all_configs = all_configs_ids

    def get_suite_name(self):
        return self.__class__.__name__

    def get_test_name(self):
        return self._testMethodName

    def _create_bucket(self, minio: MinioClient, bucket_name: str):
        try:
            minio.create_bucket(bucket_name)
        except Exception as e:
            logger.warning(f"Error creating bucket {bucket_name}: {e}")

    def create_bucket(self):
        minio = minio_client()
        for config in self.__all_configs:
            bucket_name = _get_bucket_name(config.context, config.service, config.source)
            self._create_bucket(minio, bucket_name)

    def get_debug_storage_dir(self):
        return "/app/tests/debug/{suite_name}/{test_name}".format(
            suite_name=self.get_suite_name(),
            test_name=self.get_test_name()
        )

    async def push_job(self, input_data, context_env="", service="", source=""):
        input_data["__test__"] = f"{self.get_suite_name()}.{self.get_test_name()}"
        logger.info(f"Pushing job to {context_env}.{service}.inputs.{source}")

        queue_name = self._get_queue_name(context_env, service, source)
        channel = await self.rabbitmq.create_channel()
        _ = await self.rabbitmq.create_queue(
            channel=channel,
            queue_name=queue_name,
            exchange_name="services",
            routing_key=queue_name
        )
        input_dto = generate_input(input_data, context_env, service, source)
        await self.rabbitmq.publish_message(
            "services",
            self._get_queue_name(context_env, service, source),
            serialize_to_json(input_dto)
        )

    async def _callback(self, message):
        await self.queue.put(json.loads(message.body.decode()))

    async def pop_job(self, context_env="", service="", source="", timeout=60):
        queue_name = self._get_queue_name(context_env, service, source)
        channel = await self.rabbitmq.create_channel()
        queue = await self.rabbitmq.create_queue(
            channel=channel,
            queue_name=f"{queue_name}.results",
            exchange_name="services",
            routing_key="feedback"
        )

        await self.rabbitmq.listen(queue, self._callback, timeout=timeout)

    def _get_service_process_args(self):
        self.fail(
            "Classes inheriting from BaseTestsFixture must implement _get_service_process_args method"
        )
