# update-category

The `update-category` library is a TypeScript utility library that provides a use case for updating categories in the context of a video catalog application. It includes unit and integration tests for both in-memory and Sequelize repository implementations.

## Usage

### Use Case

The primary use case, `UpdateCategoryUseCase`, is responsible for updating categories based on the provided input parameters such as ID, name, description, and is_active.

#### Example

```typescript
import { CategorySequelizeRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/infra/db/sequelize';
import { UpdateCategoryUseCase } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/update-category';

const repository = new CategorySequelizeRepository(CategoryModel);
const useCase = new UpdateCategoryUseCase(repository);

const updateParams = {
  id: 'category-id',
  name: 'Updated Category',
  description: 'Updated category description',
  is_active: true,
};

const updatedCategory = await useCase.execute(updateParams);
console.log(updatedCategory);
```

## Unit Tests

You can use `jest` testing framework for this purpose and you can run it with `nx`:

```sh
npx nx test typescript-node-services-ddd-admin-videos-catalog-category-application-use-cases-update-category
```

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.