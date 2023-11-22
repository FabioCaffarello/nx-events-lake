# repository

The `repository` library provides TypeScript interfaces and utility classes for creating repositories that interact with entities. Additionally, it includes a `search-params` class for defining search parameters and a `search-result` class for representing the results of a search.


## IRepository Interface

The `IRepository` interface defines a set of methods for basic CRUD (Create, Read, Update, Delete) operations on entities.

```typescript
import { Entity } from '@nodelib/shared/ddd-utils/entity';
import { ValueObject } from '@nodelib/shared/value-object';

export interface IRepository<E extends Entity, EntityId extends ValueObject> {
  insert(entity: E): Promise<void>;
  bulkInsert(entities: E[]): Promise<void>;
  update(entity: E): Promise<void>;
  delete(entity_id: EntityId): Promise<void>;

  findById(entity_id: EntityId): Promise<E | null>;
  findAll(): Promise<E[]>;

  getEntity(): new (...args: any[]) => E;
}
```

## ISearchableRepository Interface

The `ISearchableRepository` interface extends `IRepository` and adds additional methods for searching entities based on specified parameters.

```typescript
import { Entity } from '@nodelib/shared/ddd-utils/entity';
import { ValueObject } from '@nodelib/shared/value-object';

export interface ISearchableRepository<
  E extends Entity,
  EntityId extends ValueObject,
  Filter = string,
  SearchInput = SearchParams<Filter>,
  SearchOutput = SearchResult
> extends IRepository<E, EntityId> {
  sortableFields: string[];
  search(props: SearchInput): Promise<SearchOutput>;
}
```

## SearchParams Class

The `SearchParams` class provides a way to define search parameters, including pagination, sorting, and filtering.

```typescript
import { ValueObject } from '@nodelib/shared/value-object';

export type SortDirection = 'asc' | 'desc';

export type SearchParamsConstructorProps<Filter = string> = {
  page?: number;
  per_page?: number;
  sort?: string | null;
  sort_dir?: SortDirection | null;
  filter?: Filter | null;
};

export class SearchParams<Filter = string> extends ValueObject {
  // Implementation details...
}
```

## SearchResult Class

The `SearchResult` class represents the result of a search operation, including paginated items and metadata.

```typescript
import { Entity } from '@nodelib/shared/ddd-utils/entity';
import { ValueObject } from '@nodelib/shared/value-object';

type SearchResultConstructorProps<E extends Entity> = {
  items: E[];
  total: number;
  current_page: number;
  per_page: number;
};

export class SearchResult<A extends Entity = Entity> extends ValueObject {
  // Implementation details...
}
```

## Unit Tests

You can use `jest` testing framework for this purpose and you can run it with `nx`:

```sh
npx nx test typescript-node-shared-ddd-utils-repository
```

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
