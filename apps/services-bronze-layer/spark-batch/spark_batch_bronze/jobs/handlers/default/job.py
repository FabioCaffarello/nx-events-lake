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
    """
    Data class representing fields for a DTO.

    Attributes:
        file_names (List[str]): List of file names.
        job_name (str): Name of the job.
    """
    file_names: List[str] = field(metadata={"json_name": "file_names"})
    job_name: str = field(metadata={"json_name": "job_name"})


@dataclass
class MappingHeaderDTO:
    """
    Data class representing mapping headers for a DTO.

    Attributes:
        properties (List[FieldsDTO]): List of FieldsDTO representing properties.
    """
    properties: List[FieldsDTO] = field(metadata={"json_name": "properties"})


class Job:
    """
    Class representing a job.

    Attributes:
        _config (ConfigDTO): Configuration data transfer object.
        _service (str): Service name.
        _source (str): Source name.
        _context (str): Context information.
        _input_data (warlock.model.Model): Input data model.
        _partition (str): Data partition.
        _file_catalog_handler_client: File catalog handler client.

    Methods:
        __init__: Initializes the Job instance.
        _get_mapping_headers_from_catalog_metadata: Retrieves mapping headers from catalog metadata.
        _get_bucket_name: Generates the bucket name for Minio storage.
        _get_status: Gets the status data transfer object.
        run: Executes the job and returns results.

    """

    def __init__(self, config: ConfigDTO, input_data: type[warlock.model.Model]) -> None:
        """
        Initializes the Job instance.

        Args:
            config (ConfigDTO): Configuration data transfer object.
            input_data (type[warlock.model.Model]): Input data model.
        """
        self._config = config
        self._service = config.service
        self._source = config.source
        self._context = config.context
        self._input_data = input_data
        self._partition = input_data.partition
        self._file_catalog_handler_client = async_py_file_catalog_handler_client()

    async def _get_mapping_headers_from_catalog_metadata(self):
        """
        Retrieves mapping headers from catalog metadata.

        Returns:
            MappingHeaderDTO: Mapping headers data transfer object.
        """
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
        """
        Gets the status data transfer object.

        Returns:
            StatusDTO: Status data transfer object.
        """
        return StatusDTO(
            code=200,
            detail="Success",
        )

    async def run(self) -> Tuple[dict, StatusDTO, str]:
        """
        Executes the job and returns results.

        Returns:
            Tuple[dict, StatusDTO, str]: Tuple containing results, status, and URI.
        """
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
