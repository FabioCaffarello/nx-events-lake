# uuid

The `uuid` library provides a TypeScript class `Uuid` that encapsulates UUID (Universally Unique Identifier) generation and validation. This class extends the `ValueObject` class from `@nodelib/shared/value-object` and ensures that UUIDs used in your application are valid.

## Usage

To use the `Uuid` class, import it into your TypeScript file and create instances as needed:

```typescript
import { Uuid } from "@nodelib/shared/value-objects/uuid";

// Example usage
const myUuid = new Uuid();
console.log(myUuid.toString()); // Outputs: a valid UUID string
```

The `Uuid` class provides a default constructor that generates a new UUID using the `uuidv4` function from the "uuid" library. You can also pass an existing UUID string to the constructor to create an instance with a specific UUID.

### Validation

The `Uuid` class validates the provided or generated UUID to ensure it adheres to the UUID standard. If an invalid UUID is detected, an `InvalidUuidError` is thrown.

```typescript
import { Uuid, InvalidUuidError } from "@nodelib/shared/value-objects/uuid";

try {
  const invalidUuid = new Uuid("invalid-uuid");
} catch (error) {
  if (error instanceof InvalidUuidError) {
    console.error(error.message); // Outputs: "ID must be a valid UUID"
  }
}
```

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
