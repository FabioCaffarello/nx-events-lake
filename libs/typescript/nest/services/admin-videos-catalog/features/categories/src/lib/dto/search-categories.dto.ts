import { ListCategoriesInput } from '@nodelib/services/ddd/admin-videos-catalog/category/application/use-cases/list-categories';
import { SortDirection } from '@nodelib/shared/ddd-utils/repository';

export class SearchCategoriesDto implements ListCategoriesInput {
  page?: number;
  per_page?: number;
  sort?: string;
  sort_dir?: SortDirection;
  filter?: string;
}
