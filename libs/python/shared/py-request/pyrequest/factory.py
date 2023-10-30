from typing import Dict
import httpx
import asyncio

class RateLimitedAsyncHttpClient:
    """
    An asynchronous HTTP client with rate limiting.

    This class allows you to make HTTP requests with rate limiting to prevent
    exceeding a maximum number of requests within a specified time period.

    Args:
        base_url (str): The base URL for the HTTP requests.
        max_calls (int): The maximum number of allowed calls within the specified period.
        period (int): The time period (in seconds) during which the maximum calls are allowed.

    Attributes:
        base_url (str): The base URL for the HTTP requests.
        max_calls (int): The maximum number of allowed calls within the specified period.
        period (int): The time period (in seconds) during which the maximum calls are allowed.
        semaphore (asyncio.Semaphore): An asyncio semaphore used for rate limiting.

    """
    def __init__(self, base_url: str, max_calls: int, period: int) -> None:
        """
        Initialize the RateLimitedAsyncHttpClient with the specified parameters.

        Args:
            base_url (str): The base URL for the HTTP requests.
            max_calls (int): The maximum number of allowed calls within the specified period.
            period (int): The time period (in seconds) during which the maximum calls are allowed.
        """
        self.base_url = base_url
        self.max_calls = max_calls
        self.period = period
        self.semaphore = asyncio.Semaphore(max_calls)

    async def make_request(self, method: str, endpoint: str, data: Dict[str, any] = None, params: Dict[str, any] = None) -> Dict[str, any]:
        """
        Make an asynchronous HTTP request with rate limiting.

        This method sends an HTTP request using the specified method, endpoint, data, and parameters.
        Rate limiting is enforced to prevent exceeding the maximum number of calls within the specified period.

        Args:
            method (str): The HTTP request method (e.g., 'GET', 'POST').
            endpoint (str): The endpoint to request, relative to the base URL.
            data (dict, optional): A dictionary of data to send in the request body (as JSON).
            params (dict, optional): A dictionary of query parameters to include in the request.

        Returns:
            dict: A dictionary representing the JSON response from the HTTP request.

        Raises:
            httpx.HTTPStatusError: If the HTTP request results in an error response.

        """
        url = self.base_url + endpoint
        async with self.semaphore:
            async with httpx.AsyncClient() as client:
                response = await client.request(method, url, json=data, params=params)
                response.raise_for_status()
                return response.json()
