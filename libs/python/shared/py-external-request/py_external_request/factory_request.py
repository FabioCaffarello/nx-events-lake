import requests
from pyproxy.middleware import proxy_middleware


@proxy_middleware
def create_request(url, proxies=None, *args, **kwargs):
    response = requests.get(url, proxies=proxies, *args, **kwargs)
    return response
