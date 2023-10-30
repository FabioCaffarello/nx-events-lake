import os
from pathlib import Path
from dotenv import load_dotenv

class DotEnvLoader:
    """
    A utility class for loading environment variables from .env files.

    Args:
        environment (str): The environment for which to load environment variables.
        path (Path, optional): The path to the directory containing .env files. If not provided, assumes the current directory.

    Attributes:
        _environment (str): The environment for which environment variables are loaded.
        _path (Path): The path to the directory containing .env files.

    Methods:
        __init__(environment, path=None)
            Initializes a new DotEnvLoader instance.

        load()
            Loads environment variables from the corresponding .env file into the current environment.

        get_variable(key)
            Retrieves the value of an environment variable by its key.

    Example:
        loader = DotEnvLoader(environment="development", path=Path("/path/to/dotenv"))
        loader.load()
        value = loader.get_variable("SECRET_KEY")
    """

    def __init__(self, environment: str, path: Path = None) -> None:
        """
        Initialize a new DotEnvLoader instance.

        Args:
            environment (str): The environment for which to load environment variables.
            path (Path, optional): The path to the directory containing .env files. If not provided, assumes the current directory.
        """
        self._environment = environment
        self._path = path
        self.load()

    def _get_env(self) -> Path:
        """
        Get the path to the .env file corresponding to the specified environment.

        Returns:
            Path: The path to the .env file.
        """
        if self._path is not None:
            return self._path.joinpath(".env.{env}".format(env=self._environment))
        return Path(".env.{env}".format(env=self._environment))

    def load(self) -> None:
        """
        Load environment variables from the .env file into the current environment.
        """
        path = self._get_env()
        load_dotenv(path)

    def get_variable(self, key) -> str:
        """
        Retrieve the value of an environment variable by its key.

        Args:
            key (str): The key of the environment variable to retrieve.

        Returns:
            str: The value of the environment variable, or an empty string if not found.
        """
        return os.getenv(key, "")
