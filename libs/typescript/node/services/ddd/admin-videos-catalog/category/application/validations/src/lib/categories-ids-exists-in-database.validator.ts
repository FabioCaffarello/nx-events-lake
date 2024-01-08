import { Either } from '@nodelib/shared/ddd-utils/either';
import { NotFoundError } from '@nodelib/shared/errors';
import { Category, CategoryId } from '@nodelib/services/ddd/admin-videos-catalog/category/entity';
import { ICategoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/repository';

export class CategoriesIdExistsInDatabaseValidator {
  constructor(private categoryRepo: ICategoryRepository) {}

  async validate(
    categories_id: string[],
  ): Promise<Either<CategoryId[], NotFoundError[]>> {
    const categoriesId = categories_id.map((v) => new CategoryId(v));

    const existsResult = await this.categoryRepo.existsById(categoriesId);
    return existsResult.not_exists.length > 0
      ? Either.fail(
          existsResult.not_exists.map((c) => new NotFoundError(c.id, Category)),
        )
      : Either.ok(categoriesId);
  }
}