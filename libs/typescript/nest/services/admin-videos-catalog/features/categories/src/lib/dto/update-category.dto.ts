import { UpdateCategoryInput } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/update-category';
import { OmitType } from '@nestjs/mapped-types';

export class UpdateCategoryInputWithoutId extends OmitType(
  UpdateCategoryInput,
  ['id'] as const,
) {}

export class UpdateCategoryDto extends UpdateCategoryInputWithoutId {}
