# list-categories

The `list-categories` library is a TypeScript utility library that provides a use case for listing categories with pagination, sorting, and filtering in the context of a video catalog application. It includes unit and integration tests for both in-memory and Sequelize repository implementations.

## Usage

### Use Case

The primary use case, `ListCategoriesUseCase`, is responsible for listing categories based on the provided input parameters, such as page, per_page, sort, sort_dir, and filter.

#### Example

```typescript
import { CategorySequelizeRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/infra/db/sequelize';
import { ListCategoriesUseCase } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/list-categories';

const repository = new CategorySequelizeRepository(CategoryModel);
const useCase = new ListCategoriesUseCase(repository);

const listParams = {
  page: 1,
  per_page: 10,
  sort: 'name',
  sort_dir: 'asc',
  filter: 'action',
};

const categoryList = await useCase.execute(listParams);
console.log(categoryList);
```

## Unit Tests

You can use `jest` testing framework for this purpose and you can run it with `nx`:

```sh
npx nx test typescript-node-services-ddd-admin-videos-catalog-category-application-use-cases-list-categories
```

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.