import asyncio
from datetime import datetime
from typing import List, Tuple
import grequests
import requests
from py_external_request.factory_request import create_request
import warlock
from dto_config_handler.output import ConfigDTO
from dto_events_handler.shared import StatusDTO
from source_watcher.handlers.annual_reports.html_utils import get_href_data_from_html, get_document_download_target
from pylog.log import setup_logging
from pyminio.client import minio_client, MinioClient

logger = setup_logging(__name__)

class Job:
    """
    Represents a job that performs tasks related to data retrieval and processing.

    Attributes:
        _config (ConfigDTO): The configuration data for the job.
        _source (str): The source identifier for the job.
        _context (str): The context identifier for the job.
        _input_data (type[warlock.model.Model]): The input data for the job.
        _job_url (str): The URL associated with the job.
        _domain_url (str): The domain URL associated with the job.
        _target_endpoint (str): The target endpoint URL for the job.
        _partition (str): The formatted reference string for the job's partition.
        downloads_exceptions (List[str]): List of URLs with download exceptions.
        target_documents (List[str]): List of target document names.
        document_uris (List[str]): List of URIs for stored documents.
        companies_partition (List[str]): List of company names associated with the partition.

    Methods:
        _get_reference(reference) -> str:
            Extracts and formats the reference data.

        _get_endpoint() -> str:
            Generates the target endpoint URL.

        _get_bucket_name() -> str:
            Generates the bucket name for Minio storage.

        _get_status(response) -> StatusDTO:
            Extracts the status information from an HTTP response.

        make_request() -> requests.Response:
            Makes an HTTP GET request to the target endpoint.

        async make_request_company(minio: MinioClient, company_url: str, company_name: str) -> None:
            Asynchronously makes an HTTP GET request to a company's URL.

        async parse_response_body(response: requests.Response, minio: MinioClient) -> List[str]:
            Parses the response body to extract the document URLs.

        async run() -> Tuple[dict, StatusDTO, str]:
            Runs the job, making the HTTP request and handling the response.
    """
    def __init__(self, config: ConfigDTO, input_data: type[warlock.model.Model]) -> None:
        """
        Initializes a Job instance with the provided configuration and input data.

        Args:
            config (ConfigDTO): The configuration data.
            input_data (type[warlock.model.Model]): The input data.

        Returns:
            None
        """
        self._config = config
        self._source = config.source
        self._context = config.context
        self._input_data = input_data
        self._job_url = config.job_parameters["url"]
        self._domain_url = config.job_parameters["domain_url"]
        self._target_endpoint = self._get_endpoint()
        self._partition = self._get_reference(input_data.reference)
        self.downloads_exceptions = []
        self.target_documents = []
        self.document_uris = []
        self.companies_partition = []

    def _get_reference(self, reference) -> str:
        """
        Extracts and formats the reference data.

        Args:
            reference: The reference data.

        Returns:
            str: The formatted reference string.
        """
        logger.info(f"Reference: {reference}")
        ref = datetime(reference.year, reference.month, reference.day)
        return ref.strftime("%Y%m%d")

    def _get_endpoint(self) -> str:
        """
        Generates the target endpoint URL.

        Returns:
            str: The target endpoint URL.
        """
        return self._job_url

    def _get_bucket_name(self) -> str:
        """
        Generates the bucket name for Minio storage.

        Returns:
            str: The bucket name.
        """
        return "process-input-{context}-source-{source}".format(
            context=self._context,
            source=self._source,
        )

    def _get_status(self, response) -> StatusDTO:
        """
        Extracts the status information from an HTTP response.

        Args:
            response: The HTTP response.

        Returns:
            StatusDTO: The status information.
        """
        return StatusDTO(
            code=response.status_code,
            detail=response.reason,
        )

    def make_request(self) -> requests.Response:
        """
        Makes an HTTP GET request to the target endpoint.

        Returns:
            requests.Response: The HTTP response.
        """
        logger.info(f"endpoint: {self._target_endpoint}")
        return create_request(
            url=self._target_endpoint,
            verify=False,
            timeout=10*60,
        )

    async def make_request_company(self, minio: MinioClient, company_url: str, company_name: str) -> None:
        """
        Asynchronously makes an HTTP GET request to a company's URL.

        Args:
            minio (MinioClient): The Minio client.
            company_url (str): The company's URL.
            company_name (str): The company's name.

        Returns:
            None
        """
        headers = {
            "Sec-Fetch-Site": "same-origin",
            "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
            "Accept-Encoding": "gzip, deflate, br",
            "Accept-Language": "pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7",
            "Referer": f"https://www.annualreports.com/Company/{company_name}",
            "User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:92.0) Gecko/20100101 Firefox/92.0",
        }
        response = requests.get(
            url=company_url,
            verify=False,
            headers=headers,
            timeout=10*60,
        )
        if response.status_code >= 400:
            self.downloads_exceptions.append(company_url)
        else:
            document_target = get_document_download_target(response.content)
            if document_target is None:
                self.downloads_exceptions.append(company_url)
            else:
                page_template = f"{document_target.split('.')[0]}.html"
                file_path =  "{company}/{partition}/{document_target}".format(
                    company=company_name,
                    partition=self._partition,
                    document_target=page_template
                )
                uri = minio.upload_bytes(self._get_bucket_name(), file_path, response.content)
                self.companies_partition.append(company_name)
                self.target_documents.append(document_target)
                self.document_uris.append(uri)

    async def parse_response_body(self, response: requests.Response, minio: MinioClient) -> List[str]:
        """
        Parses the response body to extract the document URLs.

        Returns:
            List[str]: The list of document URLs.
        """
        hrefs = get_href_data_from_html(response.content)
        if hrefs is None or len(hrefs) == 0:
            raise Exception("No data found")
        companies_data = [(f"{self._domain_url}{href}", href.split("/")[-1]) for href in hrefs][:10]
        tasks = [asyncio.create_task(self.make_request_company(minio, company_url, company_name)) for company_url, company_name in companies_data]

        await asyncio.gather(*tasks)

    async def run(self) -> Tuple[dict, StatusDTO]:
        """
        Runs the job, making the HTTP request and handling the response.

        Returns:
            tuple: A tuple containing result data, status information, and the target endpoint.
        """
        logger.info(f"Job triggered with input: {self._input_data}")
        response = self.make_request()
        minio = minio_client()
        await self.parse_response_body(response, minio)
        logger.info(f"File storage uris: {self.document_uris}")
        results = []
        for _result in zip(self.companies_partition, self.document_uris):
            results.append({
                "partition": _result[0],
                "documentUri": _result[1],
            })
        logger.info(f"Job result: {results}")
        return results, self._get_status(response)


