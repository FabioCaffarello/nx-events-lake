import { getModelToken } from '@nestjs/sequelize';
import { GenreInMemoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/genre/infra/db/in-memory';
import { CreateGenreUseCase } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/use-cases/create-genre';
import { UpdateGenreUseCase } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/use-cases/update-genre';
import { ListGenresUseCase } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/use-cases/list-genres';
import { GetGenreUseCase } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/use-cases/get-genre';
import { DeleteGenreUseCase } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/use-cases/delete-genre';
import { IGenreRepository } from '@nodelib/services/ddd/admin-videos-catalog/genre/repository';
import { UnitOfWorkSequelize } from '@nodelib/shared/ddd-utils/infra/db/sequelize';
import { IUnitOfWork } from '@nodelib/shared/ddd-utils/repository';
import { ICategoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/repository';
import { CATEGORY_PROVIDERS } from '@nestlib/services/admin-videos-catalog/features/categories';
import { 
    GenreModel,
    GenreSequelizeRepository
} from '@nodelib/services/ddd/admin-videos-catalog/genre/infra/db/sequelize';
import { CategoriesIdExistsInDatabaseValidator } from '@nodelib/services/ddd/admin-videos-catalog/category/application/validations';

export const REPOSITORIES = {
  GENRE_REPOSITORY: {
    provide: 'GenreRepository',
    useExisting: GenreSequelizeRepository,
  },
  GENRE_IN_MEMORY_REPOSITORY: {
    provide: GenreInMemoryRepository,
    useClass: GenreInMemoryRepository,
  },
  GENRE_SEQUELIZE_REPOSITORY: {
    provide: GenreSequelizeRepository,
    useFactory: (genreModel: typeof GenreModel, uow: UnitOfWorkSequelize) => {
      return new GenreSequelizeRepository(genreModel, uow);
    },
    inject: [getModelToken(GenreModel), 'UnitOfWork'],
  },
};

export const USE_CASES = {
  CREATE_GENRE_USE_CASE: {
    provide: CreateGenreUseCase,
    useFactory: (
      uow: IUnitOfWork,
      genreRepo: IGenreRepository,
      categoryRepo: ICategoryRepository,
      categoriesIdValidator: CategoriesIdExistsInDatabaseValidator,
    ) => {
      return new CreateGenreUseCase(
        uow,
        genreRepo,
        categoryRepo,
        categoriesIdValidator,
      );
    },
    inject: [
      'UnitOfWork',
      REPOSITORIES.GENRE_REPOSITORY.provide,
      CATEGORY_PROVIDERS.REPOSITORIES.CATEGORY_REPOSITORY.provide,
      CATEGORY_PROVIDERS.VALIDATIONS.CATEGORIES_IDS_EXISTS_IN_DATABASE_VALIDATOR
        .provide,
    ],
  },
  UPDATE_GENRE_USE_CASE: {
    provide: UpdateGenreUseCase,
    useFactory: (
      uow: IUnitOfWork,
      genreRepo: IGenreRepository,
      categoryRepo: ICategoryRepository,
      categoriesIdExistsInStorageValidator: CategoriesIdExistsInDatabaseValidator,
    ) => {
      return new UpdateGenreUseCase(
        uow,
        genreRepo,
        categoryRepo,
        categoriesIdExistsInStorageValidator,
      );
    },
    inject: [
      'UnitOfWork',
      REPOSITORIES.GENRE_REPOSITORY.provide,
      CATEGORY_PROVIDERS.REPOSITORIES.CATEGORY_REPOSITORY.provide,
      CATEGORY_PROVIDERS.VALIDATIONS.CATEGORIES_IDS_EXISTS_IN_DATABASE_VALIDATOR
        .provide,
    ],
  },
  LIST_GENRES_USE_CASE: {
    provide: ListGenresUseCase,
    useFactory: (
      genreRepo: IGenreRepository,
      categoryRepo: ICategoryRepository,
    ) => {
      return new ListGenresUseCase(genreRepo, categoryRepo);
    },
    inject: [
      REPOSITORIES.GENRE_REPOSITORY.provide,
      CATEGORY_PROVIDERS.REPOSITORIES.CATEGORY_REPOSITORY.provide,
    ],
  },
  GET_GENRE_USE_CASE: {
    provide: GetGenreUseCase,
    useFactory: (
      genreRepo: IGenreRepository,
      categoryRepo: ICategoryRepository,
    ) => {
      return new GetGenreUseCase(genreRepo, categoryRepo);
    },
    inject: [
      REPOSITORIES.GENRE_REPOSITORY.provide,
      CATEGORY_PROVIDERS.REPOSITORIES.CATEGORY_REPOSITORY.provide,
    ],
  },
  DELETE_GENRE_USE_CASE: {
    provide: DeleteGenreUseCase,
    useFactory: (uow: IUnitOfWork, genreRepo: IGenreRepository) => {
      return new DeleteGenreUseCase(uow, genreRepo);
    },
    inject: ['UnitOfWork', REPOSITORIES.GENRE_REPOSITORY.provide],
  },
};

export const GENRES_PROVIDERS = {
  REPOSITORIES,
  USE_CASES,
};