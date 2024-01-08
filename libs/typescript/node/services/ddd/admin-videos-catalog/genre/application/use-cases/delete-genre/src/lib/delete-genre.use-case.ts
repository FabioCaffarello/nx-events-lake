import { IUseCase } from '@nodelib/shared/ddd-utils/use-case';
import { IUnitOfWork } from '@nodelib/shared/ddd-utils/repository';
import { GenreId } from '@nodelib/services/ddd/admin-videos-catalog/genre/entity';
import { IGenreRepository } from '@nodelib/services/ddd/admin-videos-catalog/genre/repository';

export class DeleteGenreUseCase
  implements IUseCase<DeleteGenreInput, DeleteGenreOutput>
{
  constructor(
    private uow: IUnitOfWork,
    private genreRepo: IGenreRepository,
  ) {}

  async execute(input: DeleteGenreInput): Promise<DeleteGenreOutput> {
    const genreId = new GenreId(input.id);
    return this.uow.do(async () => {
      return this.genreRepo.delete(genreId);
    });
  }
}

export type DeleteGenreInput = {
  id: string;
};

type DeleteGenreOutput = void;