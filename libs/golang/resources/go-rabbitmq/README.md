# go-rabbitmq

`go-rabbitmq` is a Go library that simplifies interaction with RabbitMQ using the official RabbitMQ Go client library (amqp091-go). It provides a convenient way to create, configure, and manage RabbitMQ connections, channels, and message handling in your Go applications.

## Usage

### Importing the Package

Import the `go-rabbitmq` package into your Go code:

```golang
import "github.com/yourusername/go-rabbitmq/queue"
```

### Creating a RabbitMQ Instance
You can create a RabbitMQ instance using the `NewRabbitMQ` function. Provide the user, password, host, port, vhost, consumer queue name, consumer name, dead-letter exchange name, and protocol.

```golang
user := "your-username"
password := "your-password"
host := "localhost"
port := "5672"
vhost := "/"
consumerQueueName := "your-queue"
consumerName := "your-consumer"
dlxName := "your-dlx"
protocol := "amqp"

rabbitMQ := rabbitmq.NewRabbitMQ(user, password, host, port, vhost, consumerQueueName, consumerName, dlxName, protocol)
```

### Connecting to RabbitMQ
To establish a connection with RabbitMQ, use the `Connect` method. This method will make multiple attempts to connect, and you can specify the number of attempts and the interval between them. This helps in handling potential connection issues.

```golang
channel, err := rabbitMQ.Connect()
if err != nil {
    // Handle the connection error
}
defer rabbitMQ.Close()
```

### Declaring an Exchange
You can declare an exchange using the `DeclareExchange` method. This method checks if the exchange has been declared already and declares it if not.

```golang
exchangeName := "your-exchange"
exchangeType := "direct"
rabbitMQ.DeclareExchange(exchangeName, exchangeType)
```

### Consuming Messages
To consume messages from RabbitMQ, use the `Consume` method. This method sets up a consumer that receives messages from the specified queue and sends them to the provided message channel.

```golang
messageChannel := make(chan amqp.Delivery)
exchangeName := "your-exchange"
bindingKey := "your-binding-key"
queueName := "your-queue"
consumerName := "your-consumer"
go rabbitMQ.Consume(messageChannel, exchangeName, bindingKey, queueName, consumerName)

for message := range messageChannel {
    // Handle the incoming message
}
```

### Publishing Messages
To publish messages to RabbitMQ, use the `Notify` method. This method allows you to send a message to a specified exchange with the given routing key.


```golang
message := []byte("Your message content")
contentType := "text/plain"
exchange := "your-exchange"
routingKey := "your-routing-key"

err := rabbitMQ.Notify(message, contentType, exchange, routingKey)
if err != nil {
    // Handle the publishing error
}
```

## Example
Here is a simple example of how to use `go-rabbitmq` to connect to RabbitMQ and consume and publish messages:

```golang
package main

import (
    rabbitmq "github.com/yourusername/go-rabbitmq/queue"
    "github.com/rabbitmq/amqp091-go"
    "log"
)

func main() {
    user := "your-username"
    password := "your-password"
    host := "localhost"
    port := "5672"
    vhost := "/"
    consumerQueueName := "your-queue"
    consumerName := "your-consumer"
    dlxName := "your-dlx"
    protocol := "amqp"

    rabbitMQ := rabbitmq.NewRabbitMQ(user, password, host, port, vhost, consumerQueueName, consumerName, dlxName, protocol)

    channel, err := rabbitMQ.Connect()
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
        return
    }
    defer rabbitMQ.Close()

    exchangeName := "your-exchange"
    exchangeType := "direct"
    rabbitMQ.DeclareExchange(exchangeName, exchangeType)

    messageChannel := make(chan amqp.Delivery)
    bindingKey := "your-binding-key"
    queueName := "your-queue"
    consumerName := "your-consumer"
    go rabbitMQ.Consume(messageChannel, exchangeName, bindingKey, queueName, consumerName)

    // Publish a message
    message := []byte("Hello, RabbitMQ!")
    contentType := "text/plain"
    exchange := "your-exchange"
    routingKey := "your-routing-key"

    err = rabbitMQ.Notify(message, contentType, exchange, routingKey)
    if err != nil {
        log.Printf("Failed to publish message: %v", err)
    }

    for message := range messageChannel {
        log.Printf("Received message: %s", string(message.Body))
    }
}
```
