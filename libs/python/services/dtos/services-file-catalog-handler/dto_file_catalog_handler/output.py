from dataclasses import dataclass, field
from typing import Dict

@dataclass
class FileCatalogDTO:
    id: str = field(metadata={"json_name": "id"})
    service: str = field(metadata={"json_name": "service"})
    source: str = field(metadata={"json_name": "source"})
    context: str = field(metadata={"json_name": "context"})
    lake_layer: str = field(metadata={"json_name": "lake_layer"})
    schema_type: str = field(metadata={"json_name": "schema_type"})
    catalog_id: str = field(metadata={"json_name": "catalog_id"})
    catalog: Dict[str, any] = field(metadata={"json_name": "catalog"})
    created_at: str = field(metadata={"json_name": "created_at"}, repr=False)
    updated_at: str = field(metadata={"json_name": "updated_at"}, repr=False)
