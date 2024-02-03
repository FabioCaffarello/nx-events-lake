import os
import base_fixture.fixtures  as base_fixtures


class SourceWatcherTestsFixture(base_fixtures.BaseTestsFixture):
    service_name = "source-watcher"
    cmd = "python"

    async def asyncSetUp(self):
        return await super().asyncSetUp()

    async def asyncTearDown(self):
        return await super().asyncTearDown()

    def _get_service_process_args(self):
        args = [
            self.cmd,
            f"/app/source_watcher/main.py",
        ]
        return args


class FileDownloaderTestsFixture(base_fixtures.BaseTestsFixture):
    service_name = "file-downloader"
    cmd = "python"

    async def asyncSetUp(self):
        return await super().asyncSetUp()

    async def asyncTearDown(self):
        return await super().asyncTearDown()

    def _get_service_process_args(self):
        args = [
            self.cmd,
            "/app/file_downloader/main.py",
        ]
        return args
