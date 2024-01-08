import { DeleteGenreUseCase } from '../delete-genre.use-case';
import { setupSequelize } from '@nodelib/shared/ddd-utils/infra/testing';
import { Genre, GenreId } from '@nodelib/services/ddd/admin-videos-catalog/genre/entity';
import { NotFoundError } from '@nodelib/shared/errors';
import { UnitOfWorkSequelize } from '@nodelib/shared/ddd-utils/infra/db/sequelize';
import { Category } from '@nodelib/services/ddd/admin-videos-catalog/category/entity';
import { GenreSequelizeRepository } from '@nodelib/services/ddd/admin-videos-catalog/genre/infra/db/sequelize';
import { CategorySequelizeRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/infra/db/sequelize';
import {
  GenreCategoryModel,
  GenreModel,
} from '@nodelib/services/ddd/admin-videos-catalog/genre/infra/db/sequelize';
import { CategoryModel } from '@nodelib/services/ddd/admin-videos-catalog/category/infra/db/sequelize';

describe('DeleteGenreUseCase Integration Tests', () => {
  let uow: UnitOfWorkSequelize;
  let useCase: DeleteGenreUseCase;
  let genreRepo: GenreSequelizeRepository;
  let categoryRepo: CategorySequelizeRepository;

  const sequelizeHelper = setupSequelize({
    models: [GenreModel, GenreCategoryModel, CategoryModel],
  });

  beforeEach(() => {
    uow = new UnitOfWorkSequelize(sequelizeHelper.sequelize);
    categoryRepo = new CategorySequelizeRepository(CategoryModel);
    genreRepo = new GenreSequelizeRepository(GenreModel, uow);
    useCase = new DeleteGenreUseCase(uow, genreRepo);
  });

  it('should throws error when entity not found', async () => {
    const genreId = new GenreId();
    await expect(() => useCase.execute({ id: genreId.id })).rejects.toThrow(
      new NotFoundError(genreId.id, Genre),
    );
  });

  it('should delete a genre', async () => {
    const categories = Category.fake().theCategories(2).build();
    await categoryRepo.bulkInsert(categories);
    const genre = Genre.fake()
      .aGenre()
      .addCategoryId(categories[0].category_id)
      .addCategoryId(categories[1].category_id)
      .build();
    await genreRepo.insert(genre);
    await useCase.execute({
      id: genre.genre_id.id,
    });
    await expect(genreRepo.findById(genre.genre_id)).resolves.toBeNull();
  });

  it('rollback transaction', async () => {
    const categories = Category.fake().theCategories(2).build();
    await categoryRepo.bulkInsert(categories);
    const genre = Genre.fake()
      .aGenre()
      .addCategoryId(categories[0].category_id)
      .addCategoryId(categories[1].category_id)
      .build();
    await genreRepo.insert(genre);

    GenreModel.afterBulkDestroy('hook-test', () => {
      return Promise.reject(new Error('Generic Error'));
    });

    await expect(
      useCase.execute({
        id: genre.genre_id.id,
      }),
    ).rejects.toThrow('Generic Error');

    GenreModel.removeHook('afterBulkDestroy', 'hook-test');

    const genres = await genreRepo.findAll();
    expect(genres.length).toEqual(1);
  });
});