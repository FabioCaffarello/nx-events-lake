# go-mongo

`go-mongo` is a Go library that simplifies interaction with the MongoDB database using the official MongoDB Go driver. It provides a convenient way to create, configure, and manage database connections and access MongoDB features in your Go applications.

## Usage

### Importing the Package

Import the `go-mongo` package into your Go code:

```golang
import mongodb "libs/golang/resources/go-mongo/client"
```

### Creating a MongoDB Instance
You can create a MongoDB instance using the `NewMongoDB` function. Provide the driver, user, password, host, port, database name, and a context.

```golang 
driver := "mongodb"
user := "your-username"
password := "your-password"
host := "localhost"
port := "27017"
dbName := "your-database"
ctx := context.TODO()

mongoDB := mongodb.NewMongoDB(driver, user, password, host, port, dbName, ctx)
```

### Connecting to MongoDB
To establish a connection with MongoDB, use the `Connect` method. This method will make multiple attempts to connect, and you can specify the number of attempts and the interval between them. This helps in handling potential connection issues.

```golang
client, err := mongoDB.Connect()
if err != nil {
    // Handle the connection error
}
defer mongoDB.Disconnect(client)
```

### Accessing the MongoDB Client
Once connected, you can access the MongoDB client for performing database operations:

```golang
collection := client.Database(dbName).Collection("your-collection")
```

### Disconnecting from MongoDB
After using the MongoDB client, make sure to disconnect from the database to release resources. Use the `Disconnect` method for this purpose.

### Example
Here is a simple example of how to use `go-mongo` to connect to MongoDB and perform a basic operation:

```golang
package main

import (
    "context"
    "fmt"
    "github.com/yourusername/go-mongo"
)

func main() {
    driver := "mongodb"
    user := "your-username"
    password := "your-password"
    host := "localhost"
    port := "27017"
    dbName := "your-database"
    ctx := context.TODO()

    mongoDB := mongodb.NewMongoDB(driver, user, password, host, port, dbName, ctx)

    client, err := mongoDB.Connect()
    if err != nil {
        fmt.Printf("Failed to connect to MongoDB: %v\n", err)
        return
    }
    defer mongoDB.Disconnect(client)

    collection := client.Database(dbName).Collection("your-collection")
    // Perform database operations here
}
```
