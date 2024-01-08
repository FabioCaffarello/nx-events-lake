
import { NotFoundError } from '@nodelib/shared/errors';
import { UnitOfWorkFakeInMemory } from '@nodelib/shared/ddd-utils/infra/db/in-memory';
import { Genre, GenreId } from '@nodelib/services/ddd/admin-videos-catalog/genre/entity';
import { GenreInMemoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/genre/infra/db/in-memory';
import { DeleteGenreUseCase } from '../delete-genre.use-case';

describe('DeleteGenreUseCase Unit Tests', () => {
  let useCase: DeleteGenreUseCase;
  let repository: GenreInMemoryRepository;
  let uow: UnitOfWorkFakeInMemory;

  beforeEach(() => {
    uow = new UnitOfWorkFakeInMemory();
    repository = new GenreInMemoryRepository();
    useCase = new DeleteGenreUseCase(uow, repository);
  });

  it('should throws error when entity not found', async () => {
    const genreId = new GenreId();

    await expect(() => useCase.execute({ id: genreId.id })).rejects.toThrow(
      new NotFoundError(genreId.id, Genre),
    );
  });

  it('should delete a genre', async () => {
    const items = [Genre.fake().aGenre().build()];
    repository.items = items;
    const spyOnDo = jest.spyOn(uow, 'do');
    await useCase.execute({
      id: items[0].genre_id.id,
    });
    expect(spyOnDo).toBeCalledTimes(1);
    expect(repository.items).toHaveLength(0);
  });
});