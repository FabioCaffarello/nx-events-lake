# go-events

The `go-events` package provides a simple event handling mechanism for your Go applications. It allows you to define and manage events, event handlers, and event listeners. This can be particularly useful for implementing event-driven architectures and decoupling different parts of your application.

## Usage

### Event Interface

The EventInterface defines the methods that should be implemented by event objects:

- `GetName() string`: Returns the name of the event.
- `GetDateTime() time.Time`: Returns the date and time when the event occurred.
- `GetPayload() interface{}`: Returns the event's payload.
- `SetPayload(payload interface{})`: Sets the event's payload.

### Event Handler Interface

The `EventHandlerInterface` defines the method for handling events:

- `Handle(event EventInterface, wg *sync.WaitGroup, exchangeName string, routingKey string)`: Handles the event, and you can specify the exchange name and routing key for further customization.

### Event Listener Interface

The `EventListenerInterface` defines the method for listening to events:

- `Handle(event EventInterface, wg *sync.WaitGroup)`: Handles the event and can be used in scenarios where you don't need to specify exchange names or routing keys.

### Event Dispatcher Interface

The `EventDispatcherInterface` defines methods for registering, dispatching, removing, and checking event handlers:

- `Register(eventName string, handler EventHandlerInterface) error`: Registers an event handler for a specific event.
- `Dispatch(event EventInterface, exchangeName string, routingKey string) error`: Dispatches an event to all registered handlers, optionally specifying the exchange name and routing key.
- `Remove(eventName string, handler EventHandlerInterface) error`: Removes a specific event handler.
- `Has(eventName string, handler EventHandlerInterface) bool`: Checks if a specific event handler is registered for an event.
- `Clear()`: Removes all registered event handlers.

### EventDispatcher

The `EventDispatcher` struct implements the `EventDispatcherInterface` and provides a simple way to manage event handlers and dispatch events.

```golang

import "libs/golang/shared/go-events/events"

ed := events.NewEventDispatcher()

// Register an event handler
ed.Register("myEvent", myHandler)

// Dispatch an event
ed.Dispatch(myEvent, "myExchange", "myRoutingKey")

// Remove an event handler
ed.Remove("myEvent", myHandler)

// Check if an event handler is registered
if ed.Has("myEvent", myHandler) {
    // Handler is registered
}

// Clear all event handlers
ed.Clear()
```
