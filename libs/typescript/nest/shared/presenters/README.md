# presenters

The `presenters` library provides abstract presenter classes to structure and present data in a consistent format. Presenters are particularly useful in transforming data, especially when preparing it for API responses.


## Usage

### CollectionPresenter

The `CollectionPresenter` class serves as an abstract base class for presenters that represent collections of data. It includes a `PaginationPresenter` to handle pagination details.

#### Example

```typescript
import { CollectionPresenter, PaginationPresenterProps } from '@nestlib/shared/presenters';

class MyCollectionPresenter extends CollectionPresenter {
  constructor(props: PaginationPresenterProps) {
    super(props);
  }

  get data() {
    // Implement the logic to retrieve and transform your collection data
    // ...

    return transformedData;
  }
}
```

### PaginationPresenter

The `PaginationPresenter` class transforms pagination properties into the desired format.

#### Example

```typescript
import { PaginationPresenter, PaginationPresenterProps } from '@nestlib/shared/presenters';

const paginationProps: PaginationPresenterProps = {
  current_page: 1,
  per_page: 10,
  last_page: 3,
  total: 30,
};

const paginationPresenter = new PaginationPresenter(paginationProps);
const transformedPagination = paginationPresenter; // Access transformed pagination data
```

## Classes

### CollectionPresenter

```typescript
import { Exclude, Expose } from 'class-transformer';
import { PaginationPresenter, PaginationPresenterProps } from '@nestlib/shared/presenters';

export abstract class CollectionPresenter {
  @Exclude()
  protected paginationPresenter: PaginationPresenter;

  constructor(props: PaginationPresenterProps) {
    this.paginationPresenter = new PaginationPresenter(props);
  }

  @Expose({ name: 'meta' })
  get meta() {
    return this.paginationPresenter;
  }

  abstract get data(): any;
}
```

### PaginationPresenter

```typescript
import { Transform } from 'class-transformer';

export type PaginationPresenterProps = {
  current_page: number;
  per_page: number;
  last_page: number;
  total: number;
};

export class PaginationPresenter {
  @Transform(({ value }) => parseInt(value))
  current_page: number;
  @Transform(({ value }) => parseInt(value))
  per_page: number;
  @Transform(({ value }) => parseInt(value))
  last_page: number;
  @Transform(({ value }) => parseInt(value))
  total: number;

  constructor(props: PaginationPresenterProps) {
    this.current_page = props.current_page;
    this.per_page = props.per_page;
    this.last_page = props.last_page;
    this.total = props.total;
  }
}
```


## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.