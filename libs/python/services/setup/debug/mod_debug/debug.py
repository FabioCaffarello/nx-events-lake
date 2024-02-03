import os
from pylog.log import setup_logging
import shutil

logger = setup_logging(__name__, log_level="DEBUG")

def new(debug_enabled: bool, debug_dir: str):
    if debug_enabled:
        logger.info(f"Creating debug enabled storage at: {debug_dir}")
        return EnabledDebug(debug_dir)
    else:
        logger.info("Debug disabled, creating stub debug storage")
        return DisabledDebug()

class EnabledDebug:
    def __init__(self, debug_dir: str):
        self._debug_dir = debug_dir
        self._save_responses = {}
        EnabledDebug._create_dir(self._get_response_dir())

    def save_response(self, file_name: str, response_body: str):
        filename = EnabledDebug._get_filename(file_name, self._save_responses)
        EnabledDebug._write_file(self._get_response_dir(), filename, response_body)

    def _get_response_dir(self):
        return f"{self._debug_dir}/responses/"

    @staticmethod
    def _create_dir(path: str):
        shutil.rmtree(path, ignore_errors=True)
        os.makedirs(path)

    @staticmethod
    def _get_filename(file_name: str, saved_files: dict):
        if file_name in saved_files:
            saved_files[file_name] += 1
        else:
            saved_files[file_name] = 1
        count = saved_files[file_name]
        return f"{count}-{file_name}"

    @staticmethod
    def _write_file(dirname, file_name, file_to_write):
        logger.info(f"Writing file {file_name} to {dirname}")
        with open(dirname + file_name, "wb") as writer:
            writer.write(file_to_write)


class DisabledDebug:
    def save_response(self, file_name: str, response_body: str):
        pass

