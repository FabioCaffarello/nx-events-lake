# py-request

`py-request` is a Python library that provides an asynchronous HTTP client with rate limiting. It allows you to make HTTP requests while ensuring that you do not exceed a maximum number of requests within a specified time period.

## Installation


You can install `py-dotenv` using `nx`:

```sh
npx nx add <project> --name python-shared-py-request --local
```

## Usage

To use `py-request`, you first need to import the `RateLimitedAsyncHttpClient` class from the library. Here's an example of how to use it:

```python
from pyrequest.factory import RateLimitedAsyncHttpClient

# Initialize the client with your base URL, maximum calls, and period
base_url = "http://example.com"
max_calls = 2
period = 1
client = RateLimitedAsyncHttpClient(base_url, max_calls, period)

# Make an asynchronous HTTP request
response = await client.make_request("GET", "/api/endpoint")

# The response is a dictionary representing the JSON response from the HTTP request
print(response)

```
