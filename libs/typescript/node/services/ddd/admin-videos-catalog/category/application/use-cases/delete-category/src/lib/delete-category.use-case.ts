import { ICategoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/repository';
import { IUseCase } from '@nodelib/shared/ddd-utils/use-case';
import { Uuid } from '@nodelib/shared/value-objects/uuid';

export class DeleteCategoryUseCase
  implements IUseCase<DeleteCategoryInput, DeleteCategoryOutput>
{
  constructor(private categoryRepo: ICategoryRepository) {}

  async execute(input: DeleteCategoryInput): Promise<DeleteCategoryOutput> {
    const uuid = new Uuid(input.id);
    await this.categoryRepo.delete(uuid);
  }
}

export type DeleteCategoryInput = {
  id: string;
};

type DeleteCategoryOutput = void;
