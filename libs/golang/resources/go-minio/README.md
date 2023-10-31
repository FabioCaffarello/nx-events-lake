# go-minio

`go-minio` is a Go library that simplifies interaction with the MinIO object storage server. It provides a convenient way to work with MinIO buckets and objects, including uploading, downloading, and creating buckets.

## Usage

### Importing the Package

Import the `go-minio` package into your Go code:

```golang
import gominio "libs/golang/resources/go-minio/client"
```

### Creating a MinioClient
You can create a `MinioClient` by providing the MinIO endpoint, access key, and secret key.

```golang 
minioEndpoint := "minio:9000"
minioAccessKey := "your-access-key"
minioSecretKey := "your-secret-key"

client := gominio.NewMinioClient(minioEndpoint, minioAccessKey, minioSecretKey)

```

### Downloading an Object
To download an object from a MinIO bucket, use the `DownloadFile` method. Provide the URI of the object you want to download.

```golang
uri := "minio.example.com/mybucket/myobject"
content, err := client.DownloadFile(uri)
if err != nil {
    // Handle the error
}
```

### Uploading a File
To upload a file to a MinIO bucket, use the `UploadFile` method. Provide the bucket name, file name, partition, and the file content as a byte slice.

```golang
bucketName := "mybucket"
fileName := "myobject.jpg"
partition := "2023-10-31"
fileContent := []byte("Your file content here")

path, err := client.UploadFile(bucketName, fileName, partition, fileContent)
if err != nil {
    // Handle the error
}

```
