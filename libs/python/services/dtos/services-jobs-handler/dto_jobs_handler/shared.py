from dataclasses import dataclass, field
from typing import Dict

@dataclass
class UrlDomain:
    url: str = field(metadata={"json_name": "url"})
    name: str = field(metadata={"json_name": "name"})

@dataclass
class ProxyLoader:
    name: str = field(metadata={"json_name": "name"})
    priority: int = field(metadata={"json_name": "priority"})

@dataclass
class CaptchaSolver:
    name: str = field(metadata={"json_name": "name"})
    priority: int = field(metadata={"json_name": "priority"})

