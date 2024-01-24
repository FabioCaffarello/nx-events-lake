from dataclasses import dataclass, field
from typing import Dict, List
from dto_jobs_handler.shared import UrlDomain, ProxyLoader, CaptchaSolver

@dataclass
class JobParamsHttpGateway:
    params_id: str = field(metadata={"json_name": "id"})
    service: str = field(metadata={"json_name": "service"})
    source: str = field(metadata={"json_name": "source"})
    context: str = field(metadata={"json_name": "context"})
    base_url: str = field(metadata={"json_name": "base_url"})
    url_domains: List[UrlDomain] = field(metadata={"json_name": "url_domains"})
    headers: Dict[str, str] = field(metadata={"json_name": "headers"})
    enable_proxy: bool = field(metadata={"json_name": "enable_proxy"})
    proxy_loaders: List[ProxyLoader] = field(metadata={"json_name": "proxy_loaders"})
    enable_captcha: bool = field(metadata={"json_name": "enable_captcha"})
    captcha_solvers: List[CaptchaSolver] = field(metadata={"json_name": "captcha_solvers"})



