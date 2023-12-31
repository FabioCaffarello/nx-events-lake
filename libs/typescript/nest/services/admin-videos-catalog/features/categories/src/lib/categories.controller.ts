import { CategoryOutput } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/common';
import { CreateCategoryUseCase } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/create-category';
import { DeleteCategoryUseCase } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/delete-category';
import { GetCategoryUseCase } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/get-category';
import { ListCategoriesUseCase } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/list-categories';
import { UpdateCategoryUseCase } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/update-category';
import {
  Body,
  Controller,
  Delete,
  Get,
  HttpCode,
  Inject,
  Param,
  ParseUUIDPipe,
  Patch,
  Post,
  Query,
} from '@nestjs/common';
import {
  CategoryCollectionPresenter,
  CategoryPresenter,
} from './categories.presenter';
import { CreateCategoryDto } from './dto/create-category.dto';
import { SearchCategoriesDto } from './dto/search-categories.dto';
import { UpdateCategoryDto } from './dto/update-category.dto';

@Controller('categories')
export class CategoriesController {
  @Inject(CreateCategoryUseCase)
  private createUseCase!: CreateCategoryUseCase;

  @Inject(UpdateCategoryUseCase)
  private updateUseCase!: UpdateCategoryUseCase;

  @Inject(DeleteCategoryUseCase)
  private deleteUseCase!: DeleteCategoryUseCase;

  @Inject(GetCategoryUseCase)
  private getUseCase!: GetCategoryUseCase;

  @Inject(ListCategoriesUseCase)
  private listUseCase!: ListCategoriesUseCase;

  @Post()
  async create(@Body() createCategoryDto: CreateCategoryDto) {
    const output = await this.createUseCase.execute(createCategoryDto);
    return CategoriesController.serialize(output);
  }

  @Get()
  async search(@Query() searchParamsDto: SearchCategoriesDto) {
    const output = await this.listUseCase.execute(searchParamsDto);
    return new CategoryCollectionPresenter(output);
  }

  @Get(':id')
  async findOne(
    @Param('id', new ParseUUIDPipe({ errorHttpStatusCode: 422 })) id: string,
  ) {
    const output = await this.getUseCase.execute({ id });
    return CategoriesController.serialize(output);
  }

  @Patch(':id')
  async update(
    @Param('id', new ParseUUIDPipe({ errorHttpStatusCode: 422 })) id: string,
    @Body() updateCategoryDto: UpdateCategoryDto,
  ) {
    const output = await this.updateUseCase.execute({
      ...updateCategoryDto,
      id,
    });
    return CategoriesController.serialize(output);
  }

  @HttpCode(204)
  @Delete(':id')
  remove(
    @Param('id', new ParseUUIDPipe({ errorHttpStatusCode: 422 })) id: string,
  ) {
    return this.deleteUseCase.execute({ id });
  }

  static serialize(output: CategoryOutput) {
    return new CategoryPresenter(output);
  }
}
