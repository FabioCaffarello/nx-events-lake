# py-serializer

`py-serializer` is a Python library that provides functionality to serialize and deserialize dataclass objects to and from JSON and dictionaries. It simplifies the process of converting dataclass objects into a format that can be easily stored or transmitted, such as JSON.

## Installation

You can install `py-serializer` using `nx`:

```sh
npx nx add <project> --name python-shared-py-serializer --local
```

## Usage

### Importing the Library

```python
from pyserializer.serializer import serialize_to_json, serialize_to_dict, serialize_to_dataclass

```

### Serializing a Dataclass Object to JSON

```python
# Serialize a dataclass object to a JSON string
dataclass_obj = YourDataclass(...)
json_string = serialize_to_json(dataclass_obj)
```

### Serializing a Dataclass Object to a Dictionary

```python
# Serialize a dataclass object to a dictionary
dataclass_obj = YourDataclass(...)
data_dict = serialize_to_dict(dataclass_obj)
```

### Deserializing Data to a Dataclass Object

```python
# Deserialize data from a dictionary to a dataclass object
data = {...}  # Your data in dictionary form
dataclass_obj = serialize_to_dataclass(data, YourDataclass)
```

## Examples

Here are some examples of how to use the library:

```python
from dataclasses import dataclass
from pyserializer.serializer import serialize_to_json, serialize_to_dict, serialize_to_dataclass

# Define a sample dataclass
@dataclass
class Person:
    name: str
    age: int

# Create a Person object
person = Person(name="Alice", age=30)

# Serialize the object to JSON
json_data = serialize_to_json(person)
print(json_data)  # Output: {"name": "Alice", "age": 30}

# Serialize the object to a dictionary
dict_data = serialize_to_dict(person)
print(dict_data)  # Output: {"name": "Alice", "age": 30}

# Deserialize data to a dataclass object
data = {"name": "Bob", "age": 25}
new_person = serialize_to_dataclass(data, Person)
print(new_person)  # Output: Person(name='Bob', age=25)

```

## Developing

### Run tests

```sh
npx nx test python-shared-py-serializer
```
