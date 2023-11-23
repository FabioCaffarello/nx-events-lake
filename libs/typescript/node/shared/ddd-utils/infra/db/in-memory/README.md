# in-memory

The `in-memory` library provides a simple implementation of an in-memory repository for managing entities. It is particularly useful for testing or scenarios where a persistent data store is not required.


## Usage

### InMemoryRepository

The `InMemoryRepository` class is an abstract base class that provides basic CRUD operations for entities. To use this class, extend it and implement the `getEntity` method.

#### Example:

```typescript
import { InMemoryRepository, Entity, NotFoundError, ValueObject } from '@nodelib/shared/ddd-utils';

export abstract class MyEntity extends Entity {
  // Define your entity properties here
}

export class MyEntityRepository extends InMemoryRepository<MyEntity, ValueObject> {
  getEntity(): new (...args: any[]) => MyEntity {
    return MyEntity;
  }
}
```

### InMemorySearchableRepository

The `InMemorySearchableRepository` class extends the functionality of `InMemoryRepository` by adding support for searching, sorting, and pagination of entities.

#### Example:

```typescript
import { InMemorySearchableRepository, Entity, NotFoundError, ValueObject } from '@nodelib/shared/ddd-utils';

export abstract class MySearchableEntityRepository extends InMemorySearchableRepository<MyEntity, ValueObject, string> {
  sortableFields: string[] = ['name', 'createdAt'];

  protected async applyFilter(items: MyEntity[], filter: string | null): Promise<MyEntity[]> {
    // Implement your filtering logic here
    return items.filter(item => item.name.includes(filter));
  }
}
```

## Example

Here is an example of how to use the `in-memory` library:

```typescript
import { StubEntity, StubInMemoryRepository, Uuid } from '@nodelib/shared/ddd-utils/infra/db/in-memory';

// Create an instance of the repository
const repo = new StubInMemoryRepository();

// Insert a new entity
const entity = new StubEntity({ entity_id: new Uuid(), name: 'Test', price: 100 });
await repo.insert(entity);

// Find all entities
const entities = await repo.findAll();
console.log(entities);

// Search entities with a filter
const searchResult = await repo.search({ filter: 'Test', sort: 'name', sort_dir: 'asc', page: 1, per_page: 10 });
console.log(searchResult);

// Update an entity
const updatedEntity = new StubEntity({ entity_id: entity.entity_id, name: 'UpdatedTest', price: 150 });
await repo.update(updatedEntity);

// Delete an entity
await repo.delete(entity.entity_id);
```

## Testing

You can use `jest` testing framework for this purpose and you can run it with `nx`:

```bash
npx nx test typescript-node-shared-ddd-utils-infra-db-in-memory
```

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
