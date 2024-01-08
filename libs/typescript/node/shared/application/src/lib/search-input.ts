import { SortDirection } from '@nodelib/shared/ddd-utils/repository';

export type SearchInput<Filter = string> = {
  page?: number;
  per_page?: number;
  sort?: string | null;
  sort_dir?: SortDirection | null;
  filter?: Filter | null;
};