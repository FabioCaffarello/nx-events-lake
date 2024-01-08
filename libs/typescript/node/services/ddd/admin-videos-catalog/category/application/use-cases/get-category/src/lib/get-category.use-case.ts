import { CategoryOutput, CategoryOutputMapper } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/common';
import { Category, CategoryId } from '@nodelib/services/ddd/admin-videos-catalog/category/entity';
import { ICategoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/repository';
import { IUseCase } from '@nodelib/shared/ddd-utils/use-case';
import { NotFoundError } from '@nodelib/shared/errors';

export class GetCategoryUseCase
  implements IUseCase<GetCategoryInput, GetCategoryOutput>
{
  constructor(private categoryRepo: ICategoryRepository) {}

  async execute(input: GetCategoryInput): Promise<GetCategoryOutput> {
    const categoryId = new CategoryId(input.id);
    const category = await this.categoryRepo.findById(categoryId);
    if (!category) {
      throw new NotFoundError(input.id, Category);
    }

    return CategoryOutputMapper.toOutput(category);
  }
}

export type GetCategoryInput = {
  id: string;
};

export type GetCategoryOutput = CategoryOutput;
