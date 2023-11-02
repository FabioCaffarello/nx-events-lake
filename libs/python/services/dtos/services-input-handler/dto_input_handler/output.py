from dataclasses import dataclass, field
from typing import Dict
from dto_input_handler.shared import StatusDTO, MetadataDTO

@dataclass
class InputDTO:
    id : str = field(metadata={"json_name": "id"})
    data: Dict[str, any] = field(metadata={"json_name": "data"})
    metadata: MetadataDTO = field(metadata={"json_name": "metadata"})
    status: StatusDTO = field(metadata={"json_name": "status"})
