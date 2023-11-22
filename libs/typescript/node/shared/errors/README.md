# errors

The `errors` library provides custom error classes for handling various error scenarios in TypeScript applications.

## Usage

### NotFoundError

The `NotFoundError` class is designed to be thrown when an entity is not found, typically in cases where a query for an entity by its ID fails.

Example:

```typescript
import { NotFoundError } from '@nodelib/shared/ddd-utils/errors';
import { Entity } from '@nodelib/shared/ddd-utils/entity';

class MyEntity extends Entity {
  // Your entity-specific properties and methods go here
}

const entityId = 123;
try {
  throw new NotFoundError(entityId, MyEntity);
} catch (error) {
  if (error instanceof NotFoundError) {
    console.error(error.message); // Outputs: "MyEntity Not Found using ID 123"
  }
}
```

### Note

- The `NotFoundError` class expects the entity's ID and its class as parameters during instantiation.
- The error message is constructed based on the provided entity class and ID.

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
