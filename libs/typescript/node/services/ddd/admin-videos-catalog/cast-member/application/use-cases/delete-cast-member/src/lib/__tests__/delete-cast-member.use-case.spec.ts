import { NotFoundError } from '@nodelib/shared/errors';
import {
    CastMember,
    CastMemberId,
  } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/entity';
import { CastMemberInMemoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/infra/db/in-memory';
import { DeleteCastMemberUseCase } from '../delete-cast-member.use-case';

describe('DeleteCastMemberUseCase Unit Tests', () => {
  let useCase: DeleteCastMemberUseCase;
  let repository: CastMemberInMemoryRepository;

  beforeEach(() => {
    repository = new CastMemberInMemoryRepository();
    useCase = new DeleteCastMemberUseCase(repository);
  });

  it('should throws error when entity not found', async () => {
    const castMemberId = new CastMemberId();

    await expect(() =>
      useCase.execute({ id: castMemberId.id }),
    ).rejects.toThrow(new NotFoundError(castMemberId.id, CastMember));
  });

  it('should delete a cast member', async () => {
    const items = [CastMember.fake().anActor().build()];
    repository.items = items;
    await useCase.execute({
      id: items[0].cast_member_id.id,
    });
    expect(repository.items).toHaveLength(0);
  });
});