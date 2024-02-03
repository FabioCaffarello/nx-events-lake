import unittest
import pytest
from service_fixture.fixtures import FileDownloaderTestsFixture
from pylog.log import setup_logging


logger = setup_logging(__name__)


class DefaultHandlerIntegrationTests(FileDownloaderTestsFixture):
    service_name = "file-downloader"

    async def asyncSetUp(self):
        return await super().asyncSetUp()

    async def asyncTearDown(self):
        return await super().asyncTearDown()

    @pytest.mark.asyncio
    async def test_should_download_file_for_specific_source_ceaf(self):
        valid_input = {
            "documentUri": "http://minio:9000/process-input-br-source-ceaf/20240131.zip",
            "partition": "20240131"
        }
        source = "ceaf"
        await self.push_job(valid_input, "br", self.service_name, source)
        logger.info("Job pushed to queue")
        await self.pop_job("br", self.service_name, source, 30)
        self.assertFalse(self.queue.empty())
        while not self.queue.empty():
            result = await self.queue.get()
            logger.info(f"Job result: {result}")
            self.assertEqual(result["status"]["code"], 200)
            self.assertIsNotNone(result)

    @pytest.mark.asyncio
    async def test_should_download_file_for_specific_source_cnep(self):
        valid_input = {
            "documentUri": "http://minio:9000/process-input-br-source-cnep/20240131.zip",
            "partition": "20240131"
        }
        source = "cnep"
        await self.push_job(valid_input, "br", self.service_name, source)
        logger.info("Job pushed to queue")
        await self.pop_job("br", self.service_name, source, 30)
        self.assertFalse(self.queue.empty())
        while not self.queue.empty():
            result = await self.queue.get()
            logger.info(f"Job result: {result}")
            self.assertEqual(result["status"]["code"], 200)
            self.assertIsNotNone(result)

if __name__ == "__main__":
    unittest.main()
