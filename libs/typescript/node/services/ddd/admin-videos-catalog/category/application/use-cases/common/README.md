# common

The `common` library is a TypeScript utility library that provides shared functionalities and mappers for entities in the context of a video catalog application. It includes output mappers for converting entities to standardized output formats.

## Usage

### Output Mappers

The library provides an output mapper, such as `CategoryOutputMapper`, which is designed to convert entities to a standardized output format.

#### Example

```typescript
import { Category } from '@nodelib/services/ddd/admin-videos-catalog/category/entity';
import { CategoryOutputMapper } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/common';

const entity = Category.create({
  name: 'Movie',
  description: 'some description',
  is_active: true,
});

const output = CategoryOutputMapper.toOutput(entity);
```

### Output Format

The output format is defined by the `CategoryOutput` type, which includes properties like `id`, `name`, `description`, `is_active`, and `created_at`.

#### Example

```typescript
export type CategoryOutput = {
  id: string;
  name: string;
  description: string | null;
  is_active: boolean;
  created_at: Date;
};
```

## Unit Tests

You can use `jest` testing framework for this purpose and you can run it with `nx`:

```sh
npx nx test typescript-node-services-ddd-admin-videos-catalog-category-application-use-cases-common
```

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.

