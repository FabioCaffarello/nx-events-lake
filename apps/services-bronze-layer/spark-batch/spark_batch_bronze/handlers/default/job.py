import os
from typing import Tuple, List
from dataclasses import dataclass, field
import warlock
from cli_file_catalog_handler.client import async_py_file_catalog_handler_client

from pyserializer.serializer import serialize_to_dataclass
from pyspark.sql import SparkSession
from domain_text_rename_camel_case.domain import Domain
from pyspark.sql.types import StructType, StructField, StringType
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
        logger.info("[SPARK] Creating Spark Session...")
        NESSIE_URI = os.environ.get("NESSIE_URI") ## Nessie Server URI
        WAREHOUSE = os.environ.get("WAREHOUSE") ## BUCKET TO WRITE DATA TOO
        AWS_ACCESS_KEY = os.environ.get("AWS_ACCESS_KEY") ## AWS CREDENTIALS
        AWS_SECRET_KEY = os.environ.get("AWS_SECRET_KEY") ## AWS CREDENTIALS
        AWS_S3_ENDPOINT= os.environ.get("AWS_S3_ENDPOINT") ## MINIO ENDPOINT
        logger.info(f"NESSIE_URI: {NESSIE_URI}")
        logger.info(f"WAREHOUSE: {WAREHOUSE}")
        logger.info(f"AWS_ACCESS_KEY: {AWS_ACCESS_KEY}")
        logger.info(f"AWS_SECRET_KEY: {AWS_SECRET_KEY}")
        logger.info(f"AWS_S3_ENDPOINT: {AWS_S3_ENDPOINT}")
        bucket_name = self._get_bucket_name()
        spark_session = (
            SparkSession
            .builder
            .appName("spark-batch-bronze")
            .config('spark.jars.packages', 'org.apache.iceberg:iceberg-spark-runtime-3.3_2.12:1.3.1,org.projectnessie.nessie-integrations:nessie-spark-extensions-3.3_2.12:0.67.0,software.amazon.awssdk:bundle:2.17.178,software.amazon.awssdk:url-connection-client:2.17.178')
            .config('spark.sql.extensions', 'org.apache.iceberg.spark.extensions.IcebergSparkSessionExtensions,org.projectnessie.spark.extensions.NessieSparkSessionExtensions')
            .config('spark.sql.catalog.nessie', 'org.apache.iceberg.spark.SparkCatalog')
            .config('spark.sql.catalog.nessie.uri', NESSIE_URI)
            .config('spark.sql.catalog.nessie.ref', 'main')
            .config('spark.sql.catalog.nessie.authentication.type', 'NONE')
            .config('spark.sql.catalog.nessie.catalog-impl', 'org.apache.iceberg.nessie.NessieCatalog')
            .config('spark.sql.catalog.nessie.s3.endpoint', AWS_S3_ENDPOINT)
            .config('spark.sql.catalog.nessie.warehouse', WAREHOUSE)
            .config('spark.sql.catalog.nessie.io-impl', 'org.apache.iceberg.aws.s3.S3FileIO')
            .config('spark.hadoop.fs.s3a.access.key', AWS_ACCESS_KEY)
            .config('spark.hadoop.fs.s3a.secret.key', AWS_SECRET_KEY)
            .getOrCreate()
        )
        logger.info("[SPARK-CONNECT] Spark Running...")
        logger.info(f"Job triggered with input: {self._input_data}")


# Define the schema
        schema = StructType([
            StructField("CADASTRO", StringType(), True),
            StructField("CÓDIGO DA SANÇÃO", StringType(), True),
            StructField("TIPO DE PESSOA", StringType(), True),
            StructField("CPF OU CNPJ DO SANCIONADO", StringType(), True),
            StructField("NOME DO SANCIONADO", StringType(), True),
            StructField("CATEGORIA DA SANÇÃO", StringType(), True),
            StructField("NÚMERO DO DOCUMENTO", StringType(), True),
            StructField("NÚMERO DO PROCESSO", StringType(), True),
            StructField("DATA INÍCIO SANÇÃO", StringType(), True),
            StructField("DATA FINAL SANÇÃO", StringType(), True),
            StructField("DATA PUBLICAÇÃO", StringType(), True),
            StructField("PUBLICAÇÃO", StringType(), True),
            StructField("DETALHAMENTO", StringType(), True),
            StructField("DATA DO TRÂNSITO EM JULGADO", StringType(), True),
            StructField("ABRANGÊNCIA DEFINIDA EM DECISÃO JUDICIAL", StringType(), True),
            StructField("CARGO EFETIVO", StringType(), True),
            StructField("FUNÇÃO OU CARGO DE CONFIANÇA", StringType(), True),
            StructField("ÓRGÃO DE LOTAÇÃO", StringType(), True),
            StructField("ÓRGÃO SANCIONADOR", StringType(), True),
            StructField("UF ÓRGÃO SANCIONADOR", StringType(), True),
            StructField("FUNDAMENTAÇÃO LEGAL", StringType(), True)
        ])

        # Create the dummy data as a list of tuples
        data = [
            ("CEAF", "192230", "F", "***.836.688-**", "SILVIO RICARDO SANCHEZ DE SOUZA", "Demissão", "Portaria 115",
            "50500341374201938", "17/08/2020", "", "17/08/2020", "Diário Oficial da União", "", "", "Em todos os Poderes da Esfera do Órgão sancionador",
            "Sem Informação", "Sem Informação", "AGÊNCIA NACIONAL DE TRANSPORTES TERRESTRES", "AGÊNCIA NACIONAL DE TRANSPORTES TERRESTRES", "", "ESTATUTO DOS SERVIDORES PÚBLICOS CIVIS DA UNIÃO - ART. 132, II - ABANDONO DE CARGO"),
            # Add more rows as needed
        ]

        # Create a DataFrame using the schema and data
        spark_df = spark_session.createDataFrame(data, schema=schema)
        spark_df = Domain(spark_df).transform()
        logger.info(f"Dummy dataframe lazy: {spark_df}")
        logger.info(f"Dummy dataframe: {spark_df.show()}")

        logger.info("Spark Running")
        ## Create a Table
        # spark_session.sql("CREATE TABLE nessie.names (name STRING) USING iceberg;").show()
        # ## Insert Some Data
        # spark_session.sql("INSERT INTO nessie.names VALUES ('Alex Merced'), ('Dipankar Mazumdar'), ('Jason Hughes')").show()
        # ## Query the Data
        # logger.info(f'loging nessie df: {spark_session.sql("SELECT * FROM nessie.names;").show()}')



        # df = spark_session.read.csv(uri, header=True, inferSchema=True)
        # logger.info(f"dataframe lazy: {df}")
        # logger.info(f"dataframe: {df.show()}")



        result = {"documentUri": "uri", "partition": self._partition}
        spark_session.stop()
        return result, self._get_status(), uri
