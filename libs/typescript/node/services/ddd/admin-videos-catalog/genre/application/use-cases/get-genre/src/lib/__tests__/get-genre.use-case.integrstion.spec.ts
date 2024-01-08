import { GetGenreUseCase } from '../get-genre.use-case';
import { setupSequelize } from '@nodelib/shared/ddd-utils/infra/testing';
import { Genre, GenreId } from '@nodelib/services/ddd/admin-videos-catalog/genre/entity';
import { NotFoundError } from '@nodelib/shared/errors';
import { UnitOfWorkSequelize } from '@nodelib/shared/ddd-utils/infra/db/sequelize';
import { Category } from '@nodelib/services/ddd/admin-videos-catalog/category/entity';
import { 
    GenreCategoryModel,
    GenreModel,
    GenreSequelizeRepository,
} from '@nodelib/services/ddd/admin-videos-catalog/genre/infra/db/sequelize';
import { 
    CategorySequelizeRepository,
    CategoryModel,
} from '@nodelib/services/ddd/admin-videos-catalog/category/infra/db/sequelize';

describe('GetGenreUseCase Integration Tests', () => {
  let uow: UnitOfWorkSequelize;
  let useCase: GetGenreUseCase;
  let genreRepo: GenreSequelizeRepository;
  let categoryRepo: CategorySequelizeRepository;

  const sequelizeHelper = setupSequelize({
    models: [GenreModel, GenreCategoryModel, CategoryModel],
  });

  beforeEach(() => {
    uow = new UnitOfWorkSequelize(sequelizeHelper.sequelize);
    genreRepo = new GenreSequelizeRepository(GenreModel, uow);
    categoryRepo = new CategorySequelizeRepository(CategoryModel);
    useCase = new GetGenreUseCase(genreRepo, categoryRepo);
  });

  it('should throws error when entity not found', async () => {
    const genreId = new GenreId();
    await expect(() => useCase.execute({ id: genreId.id })).rejects.toThrow(
      new NotFoundError(genreId.id, Genre),
    );
  });

  it('should returns a genre', async () => {
    const categories = Category.fake().theCategories(2).build();
    await categoryRepo.bulkInsert(categories);
    const genre = Genre.fake()
      .aGenre()
      .addCategoryId(categories[0].category_id)
      .addCategoryId(categories[1].category_id)
      .build();
    await genreRepo.insert(genre);
    const output = await useCase.execute({ id: genre.genre_id.id });
    expect(output).toStrictEqual({
      id: genre.genre_id.id,
      name: genre.name,
      categories: expect.arrayContaining([
        expect.objectContaining({
          id: categories[0].category_id.id,
          name: categories[0].name,
          created_at: categories[0].created_at,
        }),
        expect.objectContaining({
          id: categories[1].category_id.id,
          name: categories[1].name,
          created_at: categories[1].created_at,
        }),
      ]),
      categories_id: expect.arrayContaining([
        categories[0].category_id.id,
        categories[1].category_id.id,
      ]),
      is_active: true,
      created_at: genre.created_at,
    });
  });
});