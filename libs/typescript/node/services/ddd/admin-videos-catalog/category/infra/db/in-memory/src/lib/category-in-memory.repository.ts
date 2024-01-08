import { Category, CategoryId } from '@nodelib/services/ddd/admin-videos-catalog/category/entity'
import { CategoryFilter, ICategoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/repository'
import { InMemorySearchableRepository } from '@nodelib/shared/ddd-utils/infra/db/in-memory'
import { SortDirection } from '@nodelib/shared/ddd-utils/repository'


export class CategoryInMemoryRepository
  extends InMemorySearchableRepository<Category, CategoryId>
  implements ICategoryRepository
{
  sortableFields: string[] = ['name', 'created_at'];

  protected async applyFilter(
    items: Category[],
    filter: CategoryFilter | null,
  ): Promise<Category[]> {
    if (!filter) {
      return items;
    }

    return items.filter((i) => {
      return i.name.toLowerCase().includes(filter.toLowerCase());
    });
  }
  getEntity(): new (...args: any[]) => Category {
    return Category;
  }

  protected applySort(
    items: Category[],
    sort: string | null,
    sort_dir: SortDirection | null,
  ) {
    return sort
      ? super.applySort(items, sort, sort_dir)
      : super.applySort(items, 'created_at', 'desc');
  }
}