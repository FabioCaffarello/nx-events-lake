import { Category } from '@nodelib/services/ddd/admin-videos-catalog/category/entity';
import { CategoryModel, CategorySequelizeRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/infra/db/sequelize';
import { NotFoundError } from '@nodelib/shared/errors';
import { setupSequelize } from '@nodelib/shared/ddd-utils/infra/testing';
import { Uuid } from '@nodelib/shared/value-objects/uuid';
import { DeleteCategoryUseCase } from '../delete-category.use-case';

describe('DeleteCategoryUseCase Integration Tests', () => {
  let useCase: DeleteCategoryUseCase;
  let repository: CategorySequelizeRepository;

  setupSequelize({ models: [CategoryModel] });

  beforeEach(() => {
    repository = new CategorySequelizeRepository(CategoryModel);
    useCase = new DeleteCategoryUseCase(repository);
  });

  it('should throws error when entity not found', async () => {
    const uuid = new Uuid();
    await expect(() => useCase.execute({ id: uuid.id })).rejects.toThrow(
      new NotFoundError(uuid.id, Category)
    );
  });

  it('should delete a category', async () => {
    const category = Category.fake().aCategory().build();
    await repository.insert(category);
    await useCase.execute({
      id: category.category_id.id,
    });
    await expect(repository.findById(category.category_id)).resolves.toBeNull();
  });
});
