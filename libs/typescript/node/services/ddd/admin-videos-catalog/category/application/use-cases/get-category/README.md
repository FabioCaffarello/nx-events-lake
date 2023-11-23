# get-category

The `get-category` library is a TypeScript utility library that provides a use case for retrieving category details in the context of a video catalog application. It includes integration tests for both in-memory and Sequelize repository implementations.

## Usage

### Use Case

The primary use case, `GetCategoryUseCase`, is responsible for retrieving category details by ID from the specified repository.

#### Example

```typescript
import { CategorySequelizeRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/infra/db/sequelize';
import { GetCategoryUseCase } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/get-category';

const repository = new CategorySequelizeRepository(CategoryModel);
const useCase = new GetCategoryUseCase(repository);

const categoryId = '123'; // Replace with the actual category ID
const categoryDetails = await useCase.execute({ id: categoryId });
console.log(categoryDetails);
```

## Unit Tests

You can use `jest` testing framework for this purpose and you can run it with `nx`:

```sh
npx nx test typescript-node-services-ddd-admin-videos-catalog-category-application-use-cases-get-category
```

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.


