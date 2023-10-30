import json
from typing import Dict, Type
from dataclasses import dataclass, fields, is_dataclass


def _get_serialized_object(obj: dataclass) -> Dict[str, any]:
    """
    Serializes a dataclass object to a dictionary.

    Args:
        obj: A dataclass object to be serialized.

    Returns:
        dict: A dictionary containing the serialized data from the dataclass object.
    """
    data = {}
    for field_obj in fields(obj):
        field_name = field_obj.name
        field_metadata = field_obj.metadata
        json_name = field_metadata.get("json_name", field_name)
        field_value = getattr(obj, field_name)

        # Check if the field has the 'repr' attribute and it's True, and it's not None
        if field_value is not None and getattr(field_obj.default, 'repr', True):
            if is_dataclass(field_value):
                # If the field is another dataclass, recursively serialize it
                data[json_name] = _get_serialized_object(field_value)
            else:
                data[json_name] = field_value
    return data

def serialize_to_json(obj: dataclass) -> str:
    """
    Serializes a dataclass object to a JSON string.

    Args:
        obj: A dataclass object to be serialized.

    Returns:
        str: A JSON string representing the serialized data from the dataclass object.
    """
    data = _get_serialized_object(obj)
    return json.dumps(data, sort_keys=True)

def serialize_to_dict(obj: dataclass) -> Dict[str, any]:
    """
    Serializes a dataclass object to a dictionary.

    Args:
        obj: A dataclass object to be serialized.

    Returns:
        dict: A dictionary containing the serialized data from the dataclass object.
    """
    return _get_serialized_object(obj)

def serialize_to_dataclass(data: Dict[str, any], cls: Type) -> dataclass:
    """
    Deserializes data from a dictionary into a dataclass object.

    Args:
        data: A dictionary containing data to be deserialized.
        cls: The dataclass type to which the data should be deserialized.

    Returns:
        cls: An instance of the specified dataclass type with data deserialized from the input dictionary.
    """
    args = {}
    for field_obj in fields(cls):
        field_name = field_obj.name
        field_metadata = field_obj.metadata
        json_name = field_metadata.get("json_name", field_name)

        if json_name in data:
            args[field_name] = data[json_name]

    return cls(**args)
