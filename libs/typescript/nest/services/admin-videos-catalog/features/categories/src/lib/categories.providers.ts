
import { getModelToken } from '@nestjs/sequelize';
import { CategoryInMemoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/infra/db/in-memory';
import { CreateCategoryUseCase } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/create-category';
import { UpdateCategoryUseCase } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/update-category';
import { ListCategoriesUseCase } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/list-categories';
import { GetCategoryUseCase } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/get-category';
import { DeleteCategoryUseCase } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/delete-category';
import { CategorySequelizeRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/infra/db/sequelize';
import { CategoryModel } from '@nodelib/services/ddd/admin-videos-catalog/category/infra/db/sequelize';
import { ICategoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/repository';
import { CategoriesIdExistsInDatabaseValidator } from '@nodelib/services/ddd/admin-videos-catalog/category/application/validations';

export const REPOSITORIES = {
  CATEGORY_REPOSITORY: {
    provide: 'CategoryRepository',
    useExisting: CategorySequelizeRepository,
  },
  CATEGORY_IN_MEMORY_REPOSITORY: {
    provide: CategoryInMemoryRepository,
    useClass: CategoryInMemoryRepository,
  },
  CATEGORY_SEQUELIZE_REPOSITORY: {
    provide: CategorySequelizeRepository,
    useFactory: (categoryModel: typeof CategoryModel) => {
      return new CategorySequelizeRepository(categoryModel);
    },
    inject: [getModelToken(CategoryModel)],
  },
};

export const USE_CASES = {
  CREATE_CATEGORY_USE_CASE: {
    provide: CreateCategoryUseCase,
    useFactory: (categoryRepo: ICategoryRepository) => {
      return new CreateCategoryUseCase(categoryRepo);
    },
    inject: [REPOSITORIES.CATEGORY_REPOSITORY.provide],
  },
  UPDATE_CATEGORY_USE_CASE: {
    provide: UpdateCategoryUseCase,
    useFactory: (categoryRepo: ICategoryRepository) => {
      return new UpdateCategoryUseCase(categoryRepo);
    },
    inject: [REPOSITORIES.CATEGORY_REPOSITORY.provide],
  },
  LIST_CATEGORIES_USE_CASE: {
    provide: ListCategoriesUseCase,
    useFactory: (categoryRepo: ICategoryRepository) => {
      return new ListCategoriesUseCase(categoryRepo);
    },
    inject: [REPOSITORIES.CATEGORY_REPOSITORY.provide],
  },
  GET_CATEGORY_USE_CASE: {
    provide: GetCategoryUseCase,
    useFactory: (categoryRepo: ICategoryRepository) => {
      return new GetCategoryUseCase(categoryRepo);
    },
    inject: [REPOSITORIES.CATEGORY_REPOSITORY.provide],
  },
  DELETE_CATEGORY_USE_CASE: {
    provide: DeleteCategoryUseCase,
    useFactory: (categoryRepo: ICategoryRepository) => {
      return new DeleteCategoryUseCase(categoryRepo);
    },
    inject: [REPOSITORIES.CATEGORY_REPOSITORY.provide],
  },
};

export const VALIDATIONS = {
  CATEGORIES_IDS_EXISTS_IN_DATABASE_VALIDATOR: {
    provide: CategoriesIdExistsInDatabaseValidator,
    useFactory: (categoryRepo: ICategoryRepository) => {
      return new CategoriesIdExistsInDatabaseValidator(categoryRepo);
    },
    inject: [REPOSITORIES.CATEGORY_REPOSITORY.provide],
  },
};

export const CATEGORY_PROVIDERS = {
  REPOSITORIES,
  USE_CASES,
  VALIDATIONS,
};