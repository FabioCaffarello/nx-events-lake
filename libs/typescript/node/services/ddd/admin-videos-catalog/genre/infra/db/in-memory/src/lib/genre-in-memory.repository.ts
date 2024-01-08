import { Genre, GenreId } from '@nodelib/services/ddd/admin-videos-catalog/genre/entity';
import {
  IGenreRepository,
  GenreFilter,
} from '@nodelib/services/ddd/admin-videos-catalog/genre/repository';
import { SortDirection } from '@nodelib/shared/ddd-utils/repository';
import { InMemorySearchableRepository } from '@nodelib/shared/ddd-utils/infra/db/in-memory';

export class GenreInMemoryRepository
  extends InMemorySearchableRepository<Genre, GenreId, GenreFilter>
  implements IGenreRepository
{
  sortableFields: string[] = ['name', 'created_at'];

  getEntity(): new (...args: any[]) => Genre {
    return Genre;
  }

  protected async applyFilter(
    items: Genre[],
    filter: GenreFilter,
  ): Promise<Genre[]> {
    if (!filter) {
      return items;
    }

    return items.filter((genre) => {
      const containsName =
        filter.name &&
        genre.name.toLowerCase().includes(filter.name.toLowerCase());
      const containsCategoriesId =
        filter.categories_id &&
        filter.categories_id.some((c) => genre.categories_id.has(c.id));
      return filter.name && filter.categories_id
        ? containsName && containsCategoriesId
        : filter.name
        ? containsName
        : containsCategoriesId;
    });
  }

  protected applySort(
    items: Genre[],
    sort: string | null,
    sort_dir: SortDirection | null,
  ): Genre[] {
    return !sort
      ? super.applySort(items, 'created_at', 'desc')
      : super.applySort(items, sort, sort_dir);
  }
}