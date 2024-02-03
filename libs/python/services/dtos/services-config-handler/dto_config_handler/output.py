from dataclasses import dataclass, field
from typing import Dict

@dataclass
class ConfigDTO:
    id: str = field(metadata={"json_name": "id"})
    name: str = field(metadata={"json_name": "name"})
    active: bool = field(metadata={"json_name": "active"})
    frequency: str = field(metadata={"json_name": "frequency"})
    service: str = field(metadata={"json_name": "service"})
    source: str = field(metadata={"json_name": "source"})
    context: str = field(metadata={"json_name": "context"})
    output_method: str = field(metadata={"json_name": "output_method"})
    service_parameters: Dict[str, any] = field(metadata={"json_name": "service_parameters"})
    job_parameters: Dict[str, any] = field(metadata={"json_name": "job_parameters"})
