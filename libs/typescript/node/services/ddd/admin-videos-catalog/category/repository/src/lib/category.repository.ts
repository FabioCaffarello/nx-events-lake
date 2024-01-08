/* eslint-disable @typescript-eslint/no-empty-interface */
import { Category, CategoryId } from '@nodelib/services/ddd/admin-videos-catalog/category/entity';
import {
  ISearchableRepository,
  SearchParams, SearchResult
} from '@nodelib/shared/ddd-utils/repository';

export type CategoryFilter = string;

export class CategorySearchParams extends SearchParams<CategoryFilter> {}

export class CategorySearchResult extends SearchResult<Category> {}

export interface ICategoryRepository
  extends ISearchableRepository<
    Category,
    CategoryId,
    CategoryFilter,
    CategorySearchParams,
    CategorySearchResult
  > {}
