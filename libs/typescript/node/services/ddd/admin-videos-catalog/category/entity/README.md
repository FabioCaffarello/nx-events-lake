# entity

The `entity` library is a TypeScript library designed to facilitate the creation and management of entities in a domain-driven design (DDD) context. It provides a base class for entities, value objects, and various utilities for creating and validating entities.

## Usage

### Entity Class

The core of the library is the `Entity` class, which serves as the base class for entities in your application. Entities are objects that have a distinct identity and are defined by their attributes.

#### Example

```typescript
import { Entity, Uuid } from '@nodelib/shared/value-objects/uuid';

export class Category extends Entity {
  category_id: Uuid;
  name: string;
  description: string | null;
  is_active: boolean;
  created_at: Date;

  constructor(props: CategoryConstructorProps) {
    super();
    // ... constructor implementation
  }

  // ... additional methods and properties
}
```

### ValueObject Class

The library also provides a `ValueObject` class for defining value objects in your domain model.

#### Example

```typescript
import { ValueObject } from '@nodelib/shared/value-object';

export class ExampleValueObject extends ValueObject {
  // ... value object implementation
}
```

### Validation

The library includes a validation mechanism using class-validator. You can define validation rules for your entities using decorators.

#### Example

```typescript
import { MaxLength } from 'class-validator';

export class CategoryRules {
  @MaxLength(255, { groups: ['name'] })
  name: string;

  // ... other validation rules
}
```

### Validator Class

To validate entities, the library provides a `ClassValidatorFields` class that extends the functionality of class-validator.

#### Example

```typescript
import { ClassValidatorFields, Notification } from "@nodelib/shared/validators";

export class CategoryValidator extends ClassValidatorFields {
  override validate(notification: Notification, data: any, fields?: string[]): boolean {
    // ... validation logic
  }
}
```

### Factory Class

A factory class, such as `CategoryValidatorFactory`, can be used to create instances of the validator.

#### Example

```typescript
import { CategoryValidatorFactory } from '@nodelib/services/ddd/admin-videos-catalog/category/entity';

const validator = CategoryValidatorFactory.create();
```

### FakeBuilder Class

For testing purposes, the library provides a `CategoryFakeBuilder` class that allows you to create fake instances of your entities with customizable properties.

#### Example

```typescript
import { CategoryFakeBuilder } from '@nodelib/services/ddd/admin-videos-catalog/category/entity';

const fakeCategory = CategoryFakeBuilder.aCategory().build();
```

## Unit Tests

You can use `jest` testing framework for this purpose and you can run it with `nx`:

```sh
npx nx test typescript-node-services-ddd-admin-videos-catalog-category-entity
```

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
