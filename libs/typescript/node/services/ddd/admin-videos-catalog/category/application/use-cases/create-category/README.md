# create-category


The `create-category` library is a TypeScript utility library that provides a use case for creating categories in the context of a video catalog application. It includes validation for input properties and supports integration with different repositories.


## Usage

### Input Validation

The library provides an input validation class, `CreateCategoryInput`, which validates the properties of a category before creating it.

#### Example

```typescript
import { CreateCategoryInput, ValidateCreateCategoryInput } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/create-category';

const input = { name: 'Movie', description: 'some description', is_active: true };
const validatedInput = ValidateCreateCategoryInput.validate(new CreateCategoryInput(input));
```

### Use Case

The primary use case, `CreateCategoryUseCase`, is responsible for creating a category entity, validating it, and inserting it into the specified repository.

#### Example

```typescript
import { CategoryInMemoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/infra/db/in-memory';
import { CreateCategoryUseCase } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/create-category';

const repository = new CategoryInMemoryRepository();
const useCase = new CreateCategoryUseCase(repository);

const input = { name: 'Movie', description: 'some description', is_active: true };
const output = await useCase.execute(input);
```

### Output Format

The output format is similar to the `CategoryOutput` format provided by the `common` library.


## Unit Tests

You can use `jest` testing framework for this purpose and you can run it with `nx`:

```sh
npx nx test typescript-node-services-ddd-admin-videos-catalog-category-application-use-cases-create-category
```

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
