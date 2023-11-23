# in-memory

The `in-memory` library provides an in-memory repository implementation specifically designed for managing `Category` entities in the context of the `admin-videos-catalog` service. This implementation extends the generic `InMemorySearchableRepository` from `ddd-utils` and includes sorting and filtering capabilities.


## Usage

### CategoryInMemoryRepository

The `CategoryInMemoryRepository` class extends the `InMemorySearchableRepository` and is tailored for handling `Category` entities. It includes specific sorting and filtering logic for these entities.

#### Example:

```typescript
import { CategoryInMemoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/infra/db/in-memory';

// Create an instance of the repository
const categoryRepository = new CategoryInMemoryRepository();

// Use repository methods for managing Category entities
```

## Example

Here is an example of how to use the `CategoryInMemoryRepository` in your project:

```typescript
import { CategoryInMemoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/infra/db/in-memory';
import { Category } from '@nodelib/services/ddd/admin-videos-catalog/category/entity';

// Create an instance of the repository
const categoryRepository = new CategoryInMemoryRepository();

// Create and insert Category entities
const category1 = new Category({ /* Category properties */ });
const category2 = new Category({ /* Category properties */ });

await categoryRepository.insert(category1);
await categoryRepository.insert(category2);

// Search and filter Category entities
const filteredCategories = await categoryRepository.search({
  filter: 'SomeFilter',
  // Add other search parameters as needed
});

console.log(filteredCategories);

// Sorting Category entities
const sortedCategories = await categoryRepository.search({
  sort: 'name',
  sort_dir: 'asc',
});

console.log(sortedCategories);

// ... Other repository operations ...
```

## Unit Tests

You can use `jest` testing framework for this purpose and you can run it with `nx`:

```sh
npx nx test typescript-node-services-ddd-admin-videos-catalog-category-infra-db-in-memory
```

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
