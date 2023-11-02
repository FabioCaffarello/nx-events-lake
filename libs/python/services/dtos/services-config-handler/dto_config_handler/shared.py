from dataclasses import dataclass, field

@dataclass
class JobDependencies:
    service: str = field(metadata={"json_name": "service"})
    source: str = field(metadata={"json_name": "source"})
