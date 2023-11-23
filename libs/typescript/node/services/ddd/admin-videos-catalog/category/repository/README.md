# repository

The `repository` library provides an interface and related classes for working with repositories, specifically designed for managing `Category` entities in the context of the `admin-videos-catalog` service. It leverages the common repository patterns and interfaces from `ddd-utils`.


## Usage

### ICategoryRepository Interface

The `ICategoryRepository` interface extends the `ISearchableRepository` from `ddd-utils` and provides specific methods for managing `Category` entities. It defines methods for searching, creating, updating, and deleting `Category` entities.

#### Example:

```typescript
import { ICategoryRepository, CategoryFilter, CategorySearchParams, CategorySearchResult } from '@nodelib/services/ddd/admin-videos-catalog/category/repository';
import { Category } from '@nodelib/services/ddd/admin-videos-catalog/category/entity';
import { Uuid } from '@nodelib/shared/value-objects/uuid';

class MyCategoryRepository implements ICategoryRepository {
  // Implement the interface methods here
}
```

### CategorySearchParams and CategorySearchResult

The `CategorySearchParams` and `CategorySearchResult` classes extend the generic `SearchParams` and `SearchResult` classes from `ddd-utils`. They provide specific types for working with search parameters and search results related to `Category` entities.

#### Example:

```typescript
import { CategorySearchParams, CategorySearchResult } from '@nodelib/services/ddd/admin-videos-catalog/category/repository';

// Create an instance of search parameters
const searchParams = new CategorySearchParams({
  filter: 'SomeFilter',
  // Add other parameters as needed
});

// Create an instance of search results
const searchResult = new CategorySearchResult({
  items: [/* Category entities */],
  total: 10,
  current_page: 1,
  per_page: 5,
});
```

## Example

Here is an example of how to use the `repository` library in your project:

```typescript
import { ICategoryRepository, CategoryFilter, CategorySearchParams, CategorySearchResult } from '@nodelib/services/ddd/admin-videos-catalog/category/repository';
import { Category } from '@nodelib/services/ddd/admin-videos-catalog/category/entity';
import { Uuid } from '@nodelib/shared/value-objects/uuid';

class MyCategoryRepository implements ICategoryRepository {
  async search(params: CategorySearchParams): Promise<CategorySearchResult> {
    // Implement search logic here
    // Return a SearchResult object
  }

  async create(entity: Category): Promise<void> {
    // Implement creation logic here
  }

  async update(entity: Category): Promise<void> {
    // Implement update logic here
  }

  async delete(entityId: Uuid): Promise<void> {
    // Implement delete logic here
  }

  async findById(entityId: Uuid): Promise<Category | null> {
    // Implement find by ID logic here
    // Return null if not found
  }

  async findAll(): Promise<Category[]> {
    // Implement find all logic here
    // Return an array of Category entities
  }
}
```

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
