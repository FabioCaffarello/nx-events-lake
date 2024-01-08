import { Either } from '@nodelib/shared/ddd-utils/either';
import { NotFoundError } from '@nodelib/shared/errors';
import { Genre, GenreId } from '@nodelib/services/ddd/admin-videos-catalog/genre/entity';
import { IGenreRepository } from '@nodelib/services/ddd/admin-videos-catalog/genre/repository';

export class GenresIdExistsInDatabaseValidator {
  constructor(private genreRepo: IGenreRepository) {}

  async validate(
    genres_id: string[],
  ): Promise<Either<GenreId[], NotFoundError[]>> {
    const genresId = genres_id.map((v) => new GenreId(v));

    const existsResult = await this.genreRepo.existsById(genresId);
    return existsResult.not_exists.length > 0
      ? Either.fail(
          existsResult.not_exists.map((c) => new NotFoundError(c.id, Genre)),
        )
      : Either.ok(genresId);
  }
}