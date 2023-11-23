# delete-category

The `delete-category` library is a TypeScript utility library that provides a use case for deleting categories in the context of a video catalog application. It includes integration tests for both in-memory and Sequelize repository implementations.

## Usage

### Use Case

The primary use case, `DeleteCategoryUseCase`, is responsible for deleting a category entity from the specified repository.

#### Example

```typescript
import { CategorySequelizeRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/infra/db/sequelize';
import { DeleteCategoryUseCase } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/delete-category';

const repository = new CategorySequelizeRepository(CategoryModel);
const useCase = new DeleteCategoryUseCase(repository);

const categoryId = '123'; // Replace with the actual category ID
await useCase.execute({ id: categoryId });
```

## Unit Tests

You can use `jest` testing framework for this purpose and you can run it with `nx`:

```sh
npx nx test typescript-node-services-ddd-admin-videos-catalog-category-application-use-cases-delete-category
```

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.

