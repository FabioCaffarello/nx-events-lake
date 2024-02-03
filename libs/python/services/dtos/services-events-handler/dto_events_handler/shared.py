from dataclasses import dataclass, field
from typing import Dict


@dataclass
class StatusDTO:
    code: int = field(metadata={"json_name": "code"})
    detail: str = field(metadata={"json_name": "detail"})


@dataclass
class MetadataInputDTO:
    id: str = field(metadata={"json_name": "id"})
    data: Dict[str, any] = field(metadata={"json_name": "data"})
    processing_id: str = field(metadata={"json_name": "processing_id"})
    processing_timestamp: str = field(metadata={"json_name": "processing_timestamp"})
    input_schema_id: str = field(metadata={"json_name": "input_schema_id"})


@dataclass
class MetadataDTO:
    input: MetadataInputDTO = field(metadata={"json_name": "input"})
    service: str = field(metadata={"json_name": "service"})
    source: str = field(metadata={"json_name": "source"})
    context: str = field(metadata={"json_name": "context"})
    processing_timestamp: str = field(metadata={"json_name": "processing_timestamp"})
    job_frequency: str = field(metadata={"json_name": "job_frequency"})
