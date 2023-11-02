from dataclasses import dataclass, field
from typing import Dict
from dto_events_handler.shared import MetadataDTO, StatusDTO


@dataclass
class ServiceFeedbackDTO:
    data: Dict[str, any] = field(metadata={"json_name": "data"})
    metadata: MetadataDTO = field(metadata={"json_name": "metadata"})
    status: StatusDTO = field(metadata={"json_name": "status"})
