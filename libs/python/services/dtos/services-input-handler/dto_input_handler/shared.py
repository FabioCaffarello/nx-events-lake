from dataclasses import dataclass, field
from typing import Dict


@dataclass
class StatusDTO:
    code: int = field(metadata={"json_name": "code"})
    detail: str = field(metadata={"json_name": "detail"})


@dataclass
class MetadataDTO:
    processing_id: str = field(metadata={"json_name": "processing_id"})
    processing_timestamp: str = field(metadata={"json_name": "processing_timestamp"})
    service: str = field(metadata={"json_name": "service"})
    source: str = field(metadata={"json_name": "source"})
    context: str = field(metadata={"json_name": "context"})
