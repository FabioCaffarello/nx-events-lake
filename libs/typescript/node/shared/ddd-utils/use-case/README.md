# use-case

The `use-case` library provides a base interface and utility class for defining and executing use cases in TypeScript. Additionally, it includes a utility for mapping search results to a standardized pagination output.


## IUseCase Interface

The `IUseCase` interface defines a generic structure for implementing use cases. Use cases encapsulate the business logic of an application and can be executed with specific input parameters to produce output.

```typescript
export interface IUseCase<Input, Output> {
  execute(input: Input): Promise<Output>;
}
```

## PaginationOutputMapper Class

The `PaginationOutputMapper` class provides a utility method for mapping search results to a standardized pagination output format.

```typescript
import { SearchResult } from '@nodelib/shared/ddd-utils/repository';

export type PaginationOutput<Item = any> = {
  items: Item[];
  total: number;
  current_page: number;
  last_page: number;
  per_page: number;
};

export class PaginationOutputMapper {
  static toOutput<Item = any>(
    items: Item[],
    props: Omit<SearchResult, 'items'>
  ): PaginationOutput<Item> {
    return {
      items,
      total: props.total,
      current_page: props.current_page,
      last_page: props.last_page,
      per_page: props.per_page,
    };
  }
}
```

## Usage

### Creating a Use Case

To create a use case, implement the `IUseCase` interface with specific input and output types.

```typescript
import { IUseCase } from '@nodelib/shared/ddd-utils/use-case';

interface MyUseCaseInput {
  // Define your input properties
}

interface MyUseCaseOutput {
  // Define your output properties
}

class MyUseCase implements IUseCase<MyUseCaseInput, MyUseCaseOutput> {
  async execute(input: MyUseCaseInput): Promise<MyUseCaseOutput> {
    // Implement your business logic here
    return Promise.resolve(/* your output */);
  }
}
```

### Using Pagination Output Mapper

The `PaginationOutputMapper` class can be used to transform search results into a standardized pagination output format.

```typescript
import { PaginationOutputMapper, PaginationOutput } from '@nodelib/shared/ddd-utils/use-case';
import { SearchResult } from '@nodelib/shared/ddd-utils/repository';

// Example usage with search results
const searchResults: SearchResult<MyEntityType> = /* your search results */;
const paginationOutput: PaginationOutput<MyEntityType> = PaginationOutputMapper.toOutput(
  searchResults.items,
  {
    total: searchResults.total,
    current_page: searchResults.current_page,
    last_page: searchResults.last_page,
    per_page: searchResults.per_page,
  }
);

console.log(paginationOutput);
```

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.