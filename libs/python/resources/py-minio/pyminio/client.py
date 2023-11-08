from io import BytesIO
from typing import List

from minio import Minio
from minio.error import S3Error
from pylog.log import setup_logging
from pysd.service_discovery import new_from_env

logger = setup_logging(__name__)


class MinioClient:
    """
    A client class for interacting with a Minio server.

    This class provides methods for creating buckets, uploading and downloading objects, listing
    buckets and objects, and generating URIs for accessing objects on a Minio server.

    Args:
        endpoint (str): The URL of the Minio server.
        access_key (str): Access key for authentication.
        secret_key (str): Secret key for authentication.
        secure (bool, optional): If True, use a secure (HTTPS) connection. Default is False.

    Methods:
        - create_bucket(bucket_name: str) -> None
        - list_buckets() -> List[str]
        - upload_file(bucket_name: str, object_name: str, file_path: str) -> str
        - upload_bytes(bucket_name: str, object_name: str, bytes_data: bytes) -> str
        - download_file(bucket_name: str, object_name: str, file_path: str) -> None
        - list_objects(bucket_name: str) -> List[str]
        - _get_uri(bucket_name: str, object_name: str) -> str

    The class allows for interaction with a Minio server, including creating and managing buckets,
    uploading and downloading objects, and more.

    Note:
    Make sure to configure the Minio server connection with valid credentials and endpoint URL
    before using the methods of this class.

    Example Usage:
    ```
    minio = MinioClient(endpoint="http://minio.example.com", access_key="your_access_key",
                        secret_key="your_secret_key")
    minio.create_bucket("my_bucket")
    minio.upload_file("my_bucket", "example.txt", "path/to/local/file.txt")
    objects = minio.list_objects("my_bucket")
    print(objects)
    ```
    """
    def __init__(self, endpoint: str, access_key: str, secret_key: str, secure: bool = False):
        """
        Initialize a MinioClient instance.

        Args:
            endpoint (str): The URL of the Minio server.
            access_key (str): Access key for authentication.
            secret_key (str): Secret key for authentication.
            secure (bool, optional): If True, use secure (HTTPS) connection. Default is False.

        Returns:
            MinioClient: A MinioClient instance.
        """
        self._endpoint = endpoint
        self.client = Minio(
            endpoint,
            access_key=access_key,
            secret_key=secret_key,
            secure=secure,
        )

    def create_bucket(self, bucket_name: str) -> None:
        """
        Create a new bucket on the Minio server.

        Args:
            bucket_name (str): The name of the bucket to be created.

        Raises:
            Exception: If there is an error creating the bucket.
        """
        try:
            self.client.make_bucket(bucket_name)
        except S3Error as err:
            raise Exception(f"Error creating bucket: {err}")
        logger.info(f"Bucket {bucket_name} created successfully")

    def list_buckets(self) -> List[str]:
        """
        List all buckets available on the Minio server.

        Returns:
            List[str]: A list of bucket names.

        Raises:
            Exception: If there is an error listing the buckets.
        """
        try:
            return self.client.list_buckets()
        except S3Error as err:
            raise Exception(f"Error listing buckets: {err}")

    def _get_uri(self, bucket_name: str, object_name: str) -> str:
        """
        Generate a URI for accessing an object in the Minio server.

        Args:
            bucket_name (str): The name of the bucket containing the object.
            object_name (str): The name of the object for which the URI is generated.

        Returns:
            str: A URI string for accessing the specified object.

        The generated URI follows the format: "http://<MinioServerEndpoint>/<bucket_name>/<object_name>".

        Note:
        The URI may not be a valid link to the object if the Minio server is not publicly accessible.
        """
        return f"http://{self._endpoint}/{bucket_name}/{object_name}"

    def upload_file(self, bucket_name: str, object_name: str, file_path: str) -> str:
        """
        Upload a file to a specified bucket on the Minio server.

        Args:
            bucket_name (str): The name of the target bucket.
            object_name (str): The name of the object in the bucket.
            file_path (str): The local path to the file to be uploaded.

        Returns:
            str: The URI of the uploaded file.

        Raises:
            Exception: If there is an error uploading the file.
        """
        try:
            self.client.fput_object(bucket_name, object_name, file_path)
        except S3Error as err:
            raise Exception(f"Error uploading file: {err}")
        return self._get_uri(bucket_name, object_name)

    def upload_bytes(self, bucket_name: str, object_name: str, bytes_data: bytes) -> str:
        """
        Upload bytes data to a specified bucket on the Minio server.

        Args:
            bucket_name (str): The name of the target bucket.
            object_name (str): The name of the object in the bucket.
            bytes_data (bytes): The bytes data to be uploaded.

        Returns:
            str: The URI of the uploaded data.

        Raises:
            Exception: If there is an error uploading the bytes data.
        """
        try:
            data_stream = BytesIO(bytes_data)
            data_size = len(bytes_data)
            self.client.put_object(bucket_name, object_name, data_stream, data_size)
        except S3Error as err:
            raise Exception(f"Error uploading bytes: {err}")
        return self._get_uri(bucket_name, object_name)

    def download_file(self, bucket_name: str, object_name: str, file_path: str) -> None:
        """
        Download a file from a specified bucket on the Minio server and save it locally.

        Args:
            bucket_name (str): The name of the source bucket.
            object_name (str): The name of the object to be downloaded.
            file_path (str): The local path where the downloaded file will be saved.

        Raises:
            Exception: If there is an error downloading the file.
        """
        try:
            self.client.fget_object(bucket_name, object_name, file_path)
        except S3Error as err:
            raise Exception(f"Error downloading file: {err}")

    def list_objects(self, bucket_name: str) -> List[str]:
        """
        List objects in a specified bucket on the Minio server.

        Args:
            bucket_name (str): The name of the bucket.

        Returns:
            List[str]: A list of object names in the bucket.

        Raises:
            Exception: If there is an error listing objects in the bucket.
        """
        try:
            return self.client.list_objects(bucket_name)
        except S3Error as err:
            raise Exception(f"Error listing objects: {err}")


def minio_client() -> MinioClient:
    """
    Create and return a MinioClient instance by retrieving configuration from environment variables.

    Returns:
        MinioClient: A MinioClient instance configured with information from environment variables.
    """
    sd = new_from_env()
    return MinioClient(
        endpoint=sd.minio_endpoint(),
        access_key=sd.minio_access_key(),
        secret_key=sd.minio_secret_key(),
        secure=False,
    )
