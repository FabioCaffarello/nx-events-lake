import { CategoriesIdExistsInDatabaseValidator } from '@nodelib/services/ddd/admin-videos-catalog/category/application/validations';
import { ICategoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/repository';
import { IUseCase } from '@nodelib/shared/ddd-utils/use-case';
import { NotFoundError } from '@nodelib/shared/errors';
import { IUnitOfWork } from '@nodelib/shared/ddd-utils/repository';
import { EntityValidationError } from '@nodelib/shared/validators';
import { Genre, GenreId } from '@nodelib/services/ddd/admin-videos-catalog/genre/entity';
import { IGenreRepository } from '@nodelib/services/ddd/admin-videos-catalog/genre/repository';
import { GenreOutput, GenreOutputMapper } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/use-cases/common';
import { UpdateGenreInput } from './update-genre.input';

export class UpdateGenreUseCase
  implements IUseCase<UpdateGenreInput, UpdateGenreOutput>
{
  constructor(
    private uow: IUnitOfWork,
    private genreRepo: IGenreRepository,
    private categoryRepo: ICategoryRepository,
    private categoriesIdExistsInStorageValidator: CategoriesIdExistsInDatabaseValidator,
  ) {}

  async execute(input: UpdateGenreInput): Promise<UpdateGenreOutput> {
    const genreId = new GenreId(input.id);
    const genre = await this.genreRepo.findById(genreId);

    if (!genre) {
      throw new NotFoundError(input.id, Genre);
    }

    input.name && genre.changeName(input.name);

    if (input.is_active === true) {
      genre.activate();
    }

    if (input.is_active === false) {
      genre.deactivate();
    }

    const notification = genre.notification;

    if (input.categories_id) {
      const [categoriesId, errorsCategoriesId] = (
        await this.categoriesIdExistsInStorageValidator.validate(
          input.categories_id,
        )
      ).asArray();

      categoriesId && genre.syncCategoriesId(categoriesId);

      errorsCategoriesId &&
        notification.setError(
          errorsCategoriesId.map((e) => e.message),
          'categories_id',
        );
    }

    if (genre.notification.hasErrors()) {
      throw new EntityValidationError(genre.notification.toJSON());
    }

    await this.uow.do(async () => {
      return this.genreRepo.update(genre);
    });

    const categories = await this.categoryRepo.findByIds(
      Array.from(genre.categories_id.values()),
    );

    return GenreOutputMapper.toOutput(genre, categories);
  }
}

export type UpdateGenreOutput = GenreOutput;