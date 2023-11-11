from typing import Tuple, List
from dataclasses import dataclass, field
import warlock
from cli_file_catalog_handler.client import async_py_file_catalog_handler_client
from pyserializer.serializer import serialize_to_dataclass
from pyspark.sql import SparkSession
from dto_config_handler.output import ConfigDTO
from dto_events_handler.shared import StatusDTO
from pylog.log import setup_logging
from pyminio.client import minio_client


logger = setup_logging(__name__)


@dataclass
class FieldsDTO:
    file_names: List[str] = field(metadata={"json_name": "file_names"})
    job_name: str = field(metadata={"json_name": "job_name"})

@dataclass
class MappingHeaderDTO:
    properties: List[FieldsDTO] = field(metadata={"json_name": "properties"})


class Job:
    """
    Represents a job that makes HTTP requests and handles the response.

    Args:
        config (ConfigDTO): The configuration data for the job.
        input_data: The input data for the job.

    Attributes:
        _config (ConfigDTO): The configuration data for the job.
        _source  (str): The source information from the configuration.
        _context (str): The context information from the configuration.
        _input_data (type[warlock.model.Model]): The input data for the job.
        _job_url (str): The URL for the job.
        _partition (str): The partition based on input data reference.
        _target_endpoint (str): The final endpoint URL.

    Methods:

        _get_endpoint(self):
            Generates the target endpoint URL.

        _get_bucket_name(self):
            Generates the bucket name for Minio storage.

        _get_status(self, response) -> StatusDTO:
            Extracts the status information from an HTTP response.

        make_request(self):
            Makes an HTTP GET request to the target endpoint.

        run(self):
            Runs the job, making the HTTP request and handling the response.

    """

    def __init__(self, config: ConfigDTO, input_data: type[warlock.model.Model]) -> None:
        self._config = config
        self._service = config.service
        self._source = config.source
        self._context = config.context
        self._input_data = input_data
        self._partition = input_data.partition
        self._file_catalog_handler_client = async_py_file_catalog_handler_client()

    async def _get_mapping_headers_from_catalog_metadata(self):
        job_catalog = await self._file_catalog_handler_client.list_one_file_catalog_by_service_source(
            service_name=self._service,
            source_name=self._source
        )
        mapping_header = job_catalog.catalog
        return serialize_to_dataclass(mapping_header, MappingHeaderDTO)

    def _get_bucket_name(self) -> str:
        """
        Generates the bucket name for Minio storage.

        Returns:
            str: The bucket name.
        """
        return "s3a://bronze-{context}-source-{source}".format(
            context=self._context,
            source=self._source,
        )

    def _get_status(self) -> StatusDTO:
        return StatusDTO(
            code=200,
            detail="Success",
        )

    async def run(self) -> Tuple[dict, StatusDTO, str]:
        mapping_headers = await self._get_mapping_headers_from_catalog_metadata()
        logger.info(f"Mapping headers: {mapping_headers}")
        uri = self._input_data.documentUri
        logger.info(f"Input data uri: {uri}")
        logger.info(f"Input data partition: {self._partition}")
        logger.info("[SPARK-CONNECT] Creating REMOTE Spark Session...")
        bucket_name = self._get_bucket_name()
        spark_session = (
            SparkSession
            .builder
            .remote("sc://spark:15002")
            .config("log4j.logger.org.apache.hadoop.metrics2", "WARN")
            .getOrCreate()
        )
        logger.info("[SPARK-CONNECT] Spark Running...")
        logger.info(f"Job triggered with input: {self._input_data}")

        # Sample data for the PySpark DataFrame
        data = [
            {"file_names": "file1.txt", "job_name": "example_job_1"},
            {"file_names": "file2.txt", "job_name": "example_job_2"},
            {"file_names": "file3.txt", "job_name": "example_job_3"},
        ]

        df = spark_session.createDataFrame(data)
        logger.info(f"Dummy dataframe lazy: {df}")
        logger.info(f"Dummy dataframe: {df.show()}")

        result = {"documentUri": "uri", "partition": self._partition}
        spark_session.stop()
        return result, self._get_status(), uri
