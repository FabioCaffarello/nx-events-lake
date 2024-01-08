import { CategoryId } from '@nodelib/services/ddd/admin-videos-catalog/category/entity';
import {
  SearchParams,
  SearchParamsConstructorProps,
  ISearchableRepository,
  SearchResult,
} from '@nodelib/shared/ddd-utils/repository';
import { Genre, GenreId } from '@nodelib/services/ddd/admin-videos-catalog/genre/entity';

export type GenreFilter = {
  name?: string;
  categories_id?: CategoryId[];
};

export class GenreSearchParams extends SearchParams<GenreFilter> {
  private constructor(props: SearchParamsConstructorProps<GenreFilter> = {}) {
    super(props);
  }

  static create(
    props: Omit<SearchParamsConstructorProps<GenreFilter>, 'filter'> & {
      filter?: {
        name?: string;
        categories_id?: CategoryId[] | string[];
      };
    } = {},
  ) {
    const categories_id = props.filter?.categories_id?.map((c) => {
      return c instanceof CategoryId ? c : new CategoryId(c);
    });

    return new GenreSearchParams({
      ...props,
      filter: {
        name: props.filter?.name,
        categories_id,
      },
    });
  }

  get filter(): GenreFilter | null {
    return this._filter;
  }

  protected set filter(value: GenreFilter | null) {
    const _value =
      !value || (value as unknown) === '' || typeof value !== 'object'
        ? null
        : value;

    const filter = {
      ...(_value?.name && { name: `${_value.name}` }),
      ...(_value?.categories_id &&
        _value?.categories_id.length && {
          categories_id: _value.categories_id,
        }),
    };

    this._filter = Object.keys(filter).length === 0 ? null : filter;
  }
}

export class GenreSearchResult extends SearchResult<Genre> {}

export interface IGenreRepository
  extends ISearchableRepository<
    Genre,
    GenreId,
    GenreFilter,
    GenreSearchParams,
    GenreSearchResult
  > {}