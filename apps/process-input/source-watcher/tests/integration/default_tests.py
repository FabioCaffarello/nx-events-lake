import asyncio
import unittest
import pytest
from service_fixture.fixtures import SourceWatcherTestsFixture
from pylog.log import setup_logging


logger = setup_logging(__name__)


class DefaultHandlerIntegrationTests(SourceWatcherTestsFixture):
    async def asyncSetUp(self):
        return await super().asyncSetUp()

    async def asyncTearDown(self):
        return await super().asyncTearDown()


    @pytest.mark.asyncio
    async def test_should_download_file_for_specific_source_ceaf(self):
        valid_input = {
            "reference": {
                "year": 2024,
                "month": 1,
                "day": 28
            }
        }
        await self.push_job(valid_input, "br", "source-watcher", "ceaf")
        logger.info("Job pushed to queue")
        await self.pop_job("br", "source-watcher", "ceaf", 30)
        self.assertFalse(self.queue.empty())
        while not self.queue.empty():
            result = await self.queue.get()
            logger.info(f"Job result: {result}")
            self.assertEqual(result["status"]["code"], 200)
            self.assertIsNotNone(result)


if __name__ == "__main__":
    unittest.main()
