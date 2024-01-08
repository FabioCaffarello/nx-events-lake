import { CategoryModel, CategorySequelizeRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/infra/db/sequelize';
import { setupSequelize } from '@nodelib/shared/ddd-utils/infra/testing';
import { CategoryId } from '@nodelib/services/ddd/admin-videos-catalog/category/entity'
import { CreateCategoryUseCase } from '../create-category.use-case';


describe('CreateCategoryUseCase Integration Tests', () => {
  let useCase: CreateCategoryUseCase;
  let repository: CategorySequelizeRepository;

  setupSequelize({ models: [CategoryModel] });

  beforeEach(() => {
    repository = new CategorySequelizeRepository(CategoryModel);
    useCase = new CreateCategoryUseCase(repository);
  });

  it('should create a category', async () => {
    let output = await useCase.execute({ name: 'test' });
    let entity = await repository.findById(new CategoryId(output.id));
    expect(output).toStrictEqual({
      id: entity!.category_id.id,
      name: 'test',
      description: null,
      is_active: true,
      created_at: entity!.created_at,
    });

    output = await useCase.execute({
      name: 'test',
      description: 'some description',
    });
    entity = await repository.findById(new CategoryId(output.id));
    expect(output).toStrictEqual({
      id: entity!.category_id.id,
      name: 'test',
      description: 'some description',
      is_active: true,
      created_at: entity!.created_at,
    });

    output = await useCase.execute({
      name: 'test',
      description: 'some description',
      is_active: true,
    });
    entity = await repository.findById(new CategoryId(output.id));
    expect(output).toStrictEqual({
      id: entity!.category_id.id,
      name: 'test',
      description: 'some description',
      is_active: true,
      created_at: entity!.created_at,
    });

    output = await useCase.execute({
      name: 'test',
      description: 'some description',
      is_active: false,
    });
    entity = await repository.findById(new CategoryId(output.id));
    expect(output).toStrictEqual({
      id: entity!.category_id.id,
      name: 'test',
      description: 'some description',
      is_active: false,
      created_at: entity!.created_at,
    });
  });
});
