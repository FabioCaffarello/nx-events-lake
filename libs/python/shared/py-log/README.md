# py-log

`py-log` is a Python library that provides a simple way to set up logging with a JSON format. It is designed to be easy to use and configure, allowing you to quickly integrate structured logging into your Python applications.

## Installation

You can install `py-log` using `nx`:

```sh
npx nx run <project>:add --name python-shared-py-log --local
```

## Usage

### Importing the Library

```python
from pylog.log import setup_logging

# Set up logging
logger = setup_logging(module_name="your_module_name")

# Log a message
logger.info("This is an info message")

# Log an error
try:
    # some code that may raise an exception
    result = 10 / 0
except ZeroDivisionError as e:
    logger.error("An error occurred", exc_info=True)

# Change the log level (optional)
logger.setLevel(logging.WARNING)
logger.warning("This is a warning message")

# For more advanced configuration, you can also set up custom logging handlers and formatters.

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
