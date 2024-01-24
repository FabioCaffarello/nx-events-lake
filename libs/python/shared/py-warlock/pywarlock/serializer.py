import warlock
from typing import Dict

def serialize_to_dataclass(schema_parser_class: Dict[str, any], input_data: Dict[str, any]) -> type[warlock.model.Model]:
    Input_dataclass = warlock.model_factory(schema_parser_class)
    return Input_dataclass(**input_data)
