from dataclasses import dataclass, field
from typing import Dict


@dataclass
class SchemaDTO:
    id: str = field(metadata={"json_name": "id"})
    schema_type: str = field(metadata={"json_name": "schema_type"})
    service: str = field(metadata={"json_name": "service"})
    source: str = field(metadata={"json_name": "source"})
    context: str = field(metadata={"json_name": "context"})
    json_schema: Dict[str, any] = field(metadata={"json_name": "json_schema"})
    schema_id: str = field(metadata={"json_name": "schema_id"})
    created_at: str = field(metadata={"json_name": "created_at"}, repr=False)
    updated_at: str = field(metadata={"json_name": "updated_at"}, repr=False)
