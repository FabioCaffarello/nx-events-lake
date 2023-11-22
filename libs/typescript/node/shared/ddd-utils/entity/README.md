# entity

The `entity` library provides an abstract class `Entity` that serves as a base for creating entities in TypeScript. Entities represent objects with a distinct identity and encapsulate business logic.

## Usage

To create an entity, extend the `Entity` class provided by this library:

```typescript
import { Entity } from "@nodelib/shared/ddd-utils/entity";
import { ValueObject } from '@nodelib/shared/value-object';

class MyEntity extends Entity {
  private _entity_id: ValueObject;

  constructor(entityId: ValueObject) {
    super();
    this._entity_id = entityId;
  }

  get entity_id(): ValueObject {
    return this._entity_id;
  }

  toJSON(): any {
    // Your custom logic for converting the entity to JSON
  }
}

// Example usage
const entityId = new ValueObject(/* pass values for initialization */);
const myEntity = new MyEntity(entityId);

console.log(myEntity.entity_id);
console.log(myEntity.toJSON());
```

### Note

- Ensure that your entity class has an `entity_id` property of type `ValueObject` and implements the `toJSON` method based on your application requirements.

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
