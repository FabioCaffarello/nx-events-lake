import asyncio
import subprocess
import os

class SubprocessD:
    def __init__(self, subprocess_args, log_file_path):
        if not os.path.isdir(log_file_path):
            os.makedirs(log_file_path)
        self._logs = open(log_file_path + "/logs.txt", "w")
        self._subprocessd = subprocess.Popen(subprocess_args, stdout=self._logs, stderr=self._logs)

    def stop(self):
        if self._subprocessd is None:
            return
        self._subprocessd.kill()
        self._subprocessd.wait()
        self._subprocessd = None
        self._logs.close()


class SubprocessDAsync:
    def __init__(self, subprocess_args, log_file_path):
        self.subprocess_args = subprocess_args
        self.log_file_path = log_file_path
        self._subprocess = None

    async def start(self):
        if not os.path.isdir(self.log_file_path):
            os.makedirs(self.log_file_path)
        self._logs = open(os.path.join(self.log_file_path, "logs.txt"), "w")
        self._subprocess = await asyncio.create_subprocess_exec(
            *self.subprocess_args,
            stdout=self._logs,
            stderr=self._logs
        )

    async def stop(self):
        if self._subprocess is None:
            return
        self._subprocess.kill()
        await self._subprocess.wait()
        self._subprocess = None
        self._logs.close()
