# py-minio


`py-minio` is a Python library that provides a client class for interacting with a Minio server. It allows you to create buckets, upload and download objects, list buckets and objects, and generate URIs for accessing objects on a Minio server.

## Features

- Create new buckets on the Minio server.
- List all available buckets on the Minio server.
- Upload files to a specified bucket.
- Upload bytes data to a specified bucket.
- Download files from a specified bucket and save locally.
- Download files from a specified bucket as bytes.
- List objects in a specified bucket.
- Generate URIs for accessing objects on the Minio server.

## Installation

You can install `py-minio` using `nx`:

```sh
npx nx add <project> --name python-resources-py-minio --local
```

## Examples

Here's an example of how to use py-minio:

```python
from pyminio.client import MinioClient

# Initialize the Minio client
minio = MinioClient(endpoint="http://minio.example.com", access_key="your_access_key", secret_key="your_secret_key")

# Create a new bucket
minio.create_bucket("my_bucket")

# Upload a file to the bucket
minio.upload_file("my_bucket", "example.txt", "path/to/local/file.txt")

# List objects in the bucket
objects = minio.list_objects("my_bucket")
print(objects)

```


### Configuration
Before using the `py-minio` library, make sure to configure the Minio server connection with valid credentials and an endpoint URL. You can pass these details when initializing the `MinioClient`.


## API Reference

`create_bucket(bucket_name)`: Create a new bucket on the Minio server.
`list_buckets()`: List all buckets available on the Minio server.
`upload_file(bucket_name, object_name, file_path)`: Upload a file to a specified bucket on the Minio server.
`upload_bytes(bucket_name, object_name, bytes_data)`: Upload bytes data to a specified bucket on the Minio server.
`download_file(bucket_name, object_name, file_path)`: Download a file from a specified bucket on the Minio server and save it locally.
`download_file_as_bytes(bucket_name, object_name)`: Download a file from a specified bucket on the Minio server as bytes.
`list_objects(bucket_name)`: List objects in a specified bucket on the Minio server.


## Note
Make sure to configure the Minio server connection with valid credentials and an endpoint URL before using the methods of this class.
