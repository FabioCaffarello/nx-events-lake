from dataclasses import dataclass, field
from typing import List, Dict
from dto_config_handler.shared import JobDependencies

@dataclass
class ConfigDTO:
    id: str = field(metadata={"json_name": "id"})
    name: str = field(metadata={"json_name": "name"})
    active: bool = field(metadata={"json_name": "active"})
    frequency: str = field(metadata={"json_name": "frequency"})
    service: str = field(metadata={"json_name": "service"})
    source: str = field(metadata={"json_name": "source"})
    context: str = field(metadata={"json_name": "context"})
    config_id: str = field(metadata={"json_name": "config_id"})
    output_method: str = field(metadata={"json_name": "output_method"})
    depends_on: List[JobDependencies] = field(metadata={"json_name": "depends_on"})
    service_parameters: Dict[str, any] = field(metadata={"json_name": "service_parameters"})
    job_parameters: Dict[str, any] = field(metadata={"json_name": "job_parameters"})
    created_at: str = field(metadata={"json_name": "created_at"}, repr=False)
    updated_at: str = field(metadata={"json_name": "updated_at"}, repr=False)
