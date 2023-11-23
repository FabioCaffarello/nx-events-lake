import { Category } from '@nodelib/services/ddd/admin-videos-catalog/category/entity'
import { CategoryFilter, ICategoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/repository'
import { InMemorySearchableRepository } from '@nodelib/shared/ddd-utils/infra/db/in-memory'
import { SortDirection } from '@nodelib/shared/ddd-utils/repository'
import { Uuid } from '@nodelib/shared/value-objects/uuid'


export class CategoryInMemoryRepository
  extends InMemorySearchableRepository<Category, Uuid>
  implements ICategoryRepository
{
  sortableFields: string[] = ["name", "created_at"];

  protected async applyFilter(
    items: Category[],
    filter: CategoryFilter
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
    sort_dir: SortDirection | null
  ) {
    return sort
      ? super.applySort(items, sort, sort_dir)
      : super.applySort(items, "created_at", "desc");
  }
}
