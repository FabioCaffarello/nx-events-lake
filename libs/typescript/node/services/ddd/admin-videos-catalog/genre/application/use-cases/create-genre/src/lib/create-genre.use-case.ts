import { CategoriesIdExistsInDatabaseValidator } from '@nodelib/services/ddd/admin-videos-catalog/category/application/validations';
import { ICategoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/repository';
import { IUseCase } from '@nodelib/shared/ddd-utils/use-case';
import { IUnitOfWork } from '@nodelib/shared/ddd-utils/repository';
import { EntityValidationError } from '@nodelib/shared/validators';
import { Genre } from '@nodelib/services/ddd/admin-videos-catalog/genre/entity';
import { IGenreRepository } from '@nodelib/services/ddd/admin-videos-catalog/genre/repository';
import { GenreOutput, GenreOutputMapper } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/use-cases/common';
import { CreateGenreInput } from './create-genre.input';

export class CreateGenreUseCase
  implements IUseCase<CreateGenreInput, CreateGenreOutput>
{
  constructor(
    private uow: IUnitOfWork,
    private genreRepo: IGenreRepository,
    private categoryRepo: ICategoryRepository,
    private categoriesIdExistsInStorage: CategoriesIdExistsInDatabaseValidator,
  ) {}

  async execute(input: CreateGenreInput): Promise<CreateGenreOutput> {
    const [categoriesId, errorsCategoriesIds] = (
      await this.categoriesIdExistsInStorage.validate(input.categories_id)
    ).asArray();

    const { name, is_active } = input;

    const entity = Genre.create({
      name,
      categories_id: errorsCategoriesIds ? [] : categoriesId,
      is_active,
    });

    const notification = entity.notification;

    if (errorsCategoriesIds) {
      notification.setError(
        errorsCategoriesIds.map((e) => e.message),
        'categories_id',
      );
    }

    if (notification.hasErrors()) {
      throw new EntityValidationError(notification.toJSON());
    }

    await this.uow.do(async () => {
      return this.genreRepo.insert(entity);
    });

    const categories = await this.categoryRepo.findByIds(
      Array.from(entity.categories_id.values()),
    );

    return GenreOutputMapper.toOutput(entity, categories);
  }
}

export type CreateGenreOutput = GenreOutput;