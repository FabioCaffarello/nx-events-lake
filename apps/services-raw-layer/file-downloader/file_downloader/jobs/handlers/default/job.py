from datetime import datetime

import requests
from dto_config_handler.output import ConfigDTO
from dto_events.shared import StatusDTO
from pylog.log import setup_logging
from pyminio.client import minio_client

logger = setup_logging(__name__)


class Job:
    def __init__(self, config: ConfigDTO, input_data):
        self._config = config
        self._source = config.source
        self._context = config.context
        self._input_data = input_data
        self._job_url = config.job_parameters["url"]
        self._patition = self._get_reference(input_data.reference)
        self._target_endpoint = self._get_endpoint()

    def _get_reference(self, reference):
        logger.info(f"Reference: {reference}")
        ref = datetime(reference["year"], reference["month"], reference["day"])
        return ref.strftime("%Y%m%d")

    def _get_endpoint(self):
        return self._job_url.format(self._patition)

    def _get_bucket_name(self):
        return "landing-{context}-source-{source}".format(
            context=self._context,
            source=self._source,
        )

    def _get_status(self, response) -> StatusDTO:
        return StatusDTO(
            code=response.status_code,
            detail=response.reason,
        )

    def make_request(self):
        logger.info(f"endpoint: {self._target_endpoint}")
        headers = {
            "Sec-Fetch-Site": "same-origin",
            "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
            "Accept-Encoding": "gzip, deflate, br",
            "Accept-Language": "pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7",
            "User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:92.0) Gecko/20100101 Firefox/92.0",
            "Cookie": "NSC_JOibrajdb4h3qgcckulqmuceplvn5eb=ffffffff09e9012945525d5f4f58455e445a4a4229a0; sede=ffffffff09e9012245525d5f4f58455e445a4a4229a0; JSESSIONID=Ny2VOX2CrK9rbUSCh-f72i2l5nSVN4tLObFvoLuh.dzp-jboss1-01"
        }
        return requests.get(
            self._target_endpoint,
            verify=False,
            headers=headers,
            timeout=10*60,
        )

    def run(self):
        logger.info(f"Job triggered with input: {self._input_data}")
        response = self.make_request()
        minio = minio_client()
        uri = minio.upload_bytes(self._get_bucket_name(), f"{self._patition}/{self._source}.zip", response.content)
        logger.info(f"File storage uri: {uri}")
        result = {"documentUri": uri, "partition": self._patition}
        logger.info(f"Job result: {result}")
        return result, self._get_status(response), self._target_endpoint
