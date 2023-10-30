# py-dotenv

`py-dotenv` is a Python library for loading environment variables from .env files. It provides a convenient way to manage and access environment-specific configuration settings in your Python applications.

## Installation

You can install `py-dotenv` using `nx`:

```sh
npx nx add <project> --name python-shared-py-dotenv --local
```

## Usage

### Importing the Library

```python
from dotenv import load_dotenv
```

#### Class: DotEnvLoader

DotEnvLoader is a utility class provided by py-dotenv to manage the loading of environment variables from .env files.

Initialization
To begin using DotEnvLoader, you can create an instance by specifying the environment and an optional path to the directory containing .env files.

Retrieving Environment Variables
You can retrieve the value of an environment variable by its key using the get_variable method.

```python

loader_vars = DotEnvLoader(environment="development", path=Path("/path/to/dotenv"))

env_value = loader.get_variable("SECRET_KEY")

```


## Configuration
The setup_logging function is used to configure the logger. It takes the following parameters:

- `module_name` (str): The name of the module or application that is using the logger.
`propagate` (bool): Whether to propagate the logging to the parent logger.
`log_level` (str): The log level to set, which can be one of "DEBUG", "INFO", "WARNING", "ERROR", or "CRITICAL". The default is "INFO".

By default, `py-log` configures a JSON logger that writes log entries to the console. You can customize the log format and destination by modifying the setup_logging function to suit your specific needs.

## Example JSON Log Output

```json
{"levelname": "INFO", "filename": "example.py", "message": "This is an info message"}
{"levelname": "ERROR", "filename": "example.py", "message": "An error occurred", "exc_info": "Traceback (most recent call last):\n  File \"example.py\", line 11, in <module>\n    result = 10 / 0\nZeroDivisionError: division by zero"}
{"levelname": "WARNING", "filename": "example.py", "message": "This is a warning message"}
```
