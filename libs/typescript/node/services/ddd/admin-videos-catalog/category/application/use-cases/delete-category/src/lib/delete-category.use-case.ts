import { ICategoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/repository';
import { IUseCase } from '@nodelib/shared/ddd-utils/use-case';
import { CategoryId } from '@nodelib/services/ddd/admin-videos-catalog/category/entity'

export class DeleteCategoryUseCase
  implements IUseCase<DeleteCategoryInput, DeleteCategoryOutput>
{
  constructor(private categoryRepo: ICategoryRepository) {}

  async execute(input: DeleteCategoryInput): Promise<DeleteCategoryOutput> {
    const categoryId = new CategoryId(input.id);
    await this.categoryRepo.delete(categoryId);
  }
}

export type DeleteCategoryInput = {
  id: string;
};

type DeleteCategoryOutput = void;
