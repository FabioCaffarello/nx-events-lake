import unittest
from dataclasses import dataclass, field
from typing import Optional, List
from pyserializer.serializer import serialize_to_json, serialize_to_dataclass, serialize_to_dict

@dataclass
class Config:
    other_field: Optional[str] = field(init=False, default=None, metadata={"json_name": "otherField"}, repr=False)
    subscriptions_active: Optional[List[str]] = field(default=None, metadata={"json_name": "subscriptionsActive"})


class TestSerializer(unittest.TestCase):

    def test_serialize_to_json(self):
        config = Config(subscriptions_active=["A", "B"])
        expected_json = '{"subscriptionsActive": ["A", "B"]}'
        serialized_json = serialize_to_json(config)
        self.assertEqual(serialized_json, expected_json)

    def test_serialize_to_dict(self):
        config = Config(subscriptions_active=["A", "B"])
        expected_json = {"subscriptionsActive": ["A", "B"]}
        serialized_json = serialize_to_dict(config)
        self.assertEqual(serialized_json, expected_json)

    def test_from_data_to_dataclass(self):
        data = {"subscriptionsActive": ["X", "Y"]}
        config = serialize_to_dataclass(data, Config)
        expected_config = Config(subscriptions_active=["X", "Y"])
        self.assertEqual(config, expected_config)

    def test_from_data_to_dataclass_with_missing_field(self):
        # Test when a field is missing in the data
        data = {}
        config = serialize_to_dataclass(data, Config)
        expected_config = Config()
        self.assertEqual(config, expected_config)

    def test_from_data_to_dataclass_with_extra_field(self):
        # Test when there's an extra field in the data
        data = {"subscriptionsActive": ["X", "Y"], "extraField": "value"}
        config = serialize_to_dataclass(data, Config)
        expected_config = Config(subscriptions_active=["X", "Y"])
        self.assertEqual(config, expected_config)

if __name__ == '__main__':
    unittest.main()
