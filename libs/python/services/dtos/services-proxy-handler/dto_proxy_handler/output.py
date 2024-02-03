from dataclasses import dataclass, field

@dataclass
class TorProxyDTO:
    origin: str = field(metadata={"json_name": "origin"})
