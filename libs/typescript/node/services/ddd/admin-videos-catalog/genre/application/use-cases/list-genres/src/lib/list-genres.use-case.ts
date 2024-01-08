import { GenreOutput, GenreOutputMapper } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/use-cases/common';
import {
  PaginationOutput,
  PaginationOutputMapper,
} from '@nodelib/shared/application';
import {
  IGenreRepository,
  GenreSearchParams,
  GenreSearchResult,
} from '@nodelib/services/ddd/admin-videos-catalog/genre/repository';
import { ICategoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/repository';
import { ListGenresInput } from './list-genres.input';
import { IUseCase } from '@nodelib/shared/ddd-utils/use-case';
import { CategoryId } from '@nodelib/services/ddd/admin-videos-catalog/category/entity';

export class ListGenresUseCase
  implements IUseCase<ListGenresInput, ListGenresOutput>
{
  constructor(
    private genreRepo: IGenreRepository,
    private categoryRepo: ICategoryRepository,
  ) {}

  async execute(input: ListGenresInput): Promise<ListGenresOutput> {
    const params = GenreSearchParams.create(input);
    const searchResult = await this.genreRepo.search(params);
    return this.toOutput(searchResult);
  }

  private async toOutput(
    searchResult: GenreSearchResult,
  ): Promise<ListGenresOutput> {
    const { items: _items } = searchResult;

    const categoriesIdRelated = searchResult.items.reduce<CategoryId[]>(
      (acc, item) => {
        return acc.concat([...item.categories_id.values()]);
      },
      [],
    );
    //TODO - retirar duplicados
    const categoriesRelated =
      await this.categoryRepo.findByIds(categoriesIdRelated);

    const items = _items.map((i) => {
      const categoriesOfGenre = categoriesRelated.filter((c) =>
        i.categories_id.has(c.category_id.id),
      );
      return GenreOutputMapper.toOutput(i, categoriesOfGenre);
    });
    return PaginationOutputMapper.toOutput(items, searchResult);
  }
}

export type ListGenresOutput = PaginationOutput<GenreOutput>;