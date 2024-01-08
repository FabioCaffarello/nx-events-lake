import { ICategoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/repository';
import { IUseCase } from '@nodelib/shared/ddd-utils/use-case';
import { NotFoundError } from '@nodelib/shared/errors';
import { Genre, GenreId } from '@nodelib/services/ddd/admin-videos-catalog/genre/entity';
import { IGenreRepository } from '@nodelib/services/ddd/admin-videos-catalog/genre/repository';
import { GenreOutput, GenreOutputMapper } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/use-cases/common';

export class GetGenreUseCase
  implements IUseCase<GetGenreInput, GetGenreOutput>
{
  constructor(
    private genreRepo: IGenreRepository,
    private categoryRepo: ICategoryRepository,
  ) {}

  async execute(input: GetGenreInput): Promise<GetGenreOutput> {
    const genreId = new GenreId(input.id);
    const genre = await this.genreRepo.findById(genreId);
    if (!genre) {
      throw new NotFoundError(input.id, Genre);
    }
    const categories = await this.categoryRepo.findByIds([
      ...genre.categories_id.values(),
    ]);
    return GenreOutputMapper.toOutput(genre, categories);
  }
}

export type GetGenreInput = {
  id: string;
};

export type GetGenreOutput = GenreOutput;