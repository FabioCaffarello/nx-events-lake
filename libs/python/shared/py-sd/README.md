# py-sd

`py-sd` is a Python library for service discovery and environment variable management. It provides a convenient way to access service endpoints and environment variables needed for a distributed system.

## Installation

You can install `py-sd` using `nx`:

```sh
npx nx add <project> --name python-shared-py-sd --local
```

## Usage

### Importing the Library

```python
from pysd.service_discovery import new_from_env

sd = new_from_env()

# Get the RabbitMQ endpoint
rabbitmq_endpoint = sd.rabbitmq_endpoint()
```

### Handling Exceptions
`py-sd` provides two custom exceptions:

- `UnrecoverableError`: Raised when environment variables are not set.
- `ServiceUnavailableError`: Raised when a required environment variable is not set.
You can handle these exceptions in your code to provide appropriate error handling.


## Examples

Here are some examples of how to use the py-sd library:

```python
import os
from pysd import ServiceDiscovery, UnrecoverableError, ServiceUnavailableError

# Create a ServiceDiscovery instance using environment variables
service_discovery = ServiceDiscovery(os.environ)

try:
    # Get the RabbitMQ endpoint
    rabbitmq_endpoint = service_discovery.rabbitmq_endpoint()
    print(f"RabbitMQ endpoint: {rabbitmq_endpoint}")

    # Get the Minio access key
    minio_access_key = service_discovery.minio_access_key()
    print(f"Minio access key: {minio_access_key}")

except UnrecoverableError as e:
    print(f"Unrecoverable error: {str(e)}")

except ServiceUnavailableError as e:
    print(f"Service unavailable: {str(e)}")

```
