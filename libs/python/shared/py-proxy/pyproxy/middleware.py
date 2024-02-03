
from functools import wraps


def proxy_middleware(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        kwargs["proxies"] = {'http': f'socks5://tor:9050', 'https': f'socks5://tor:9050'}
        return func(*args, **kwargs)
    return wrapper
