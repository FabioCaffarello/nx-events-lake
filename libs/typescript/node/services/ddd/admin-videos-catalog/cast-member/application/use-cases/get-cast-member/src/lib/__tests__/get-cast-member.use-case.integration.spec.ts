import { GetCastMemberUseCase } from '../get-cast-member.use-case';
import { setupSequelize } from '@nodelib/shared/ddd-utils/infra/testing';
import {
  CastMember,
  CastMemberId,
} from '@nodelib/services/ddd/admin-videos-catalog/cast-member/entity';
import { NotFoundError } from '@nodelib/shared/errors';
import {
  CastMemberModel,
  CastMemberSequelizeRepository,
} from '@nodelib/services/ddd/admin-videos-catalog/cast-member/infra/db/sequelize';

describe('GetCastMemberUseCase Integration Tests', () => {
  let useCase: GetCastMemberUseCase;
  let repository: CastMemberSequelizeRepository;

  setupSequelize({ models: [CastMemberModel] });

  beforeEach(() => {
    repository = new CastMemberSequelizeRepository(CastMemberModel);
    useCase = new GetCastMemberUseCase(repository);
  });

  it('should throws error when entity not found', async () => {
    const castMemberId = new CastMemberId();
    await expect(() =>
      useCase.execute({ id: castMemberId.id }),
    ).rejects.toThrow(new NotFoundError(castMemberId.id, CastMember));
  });

  it('should returns a cast member', async () => {
    const castMember = CastMember.fake().anActor().build();
    await repository.insert(castMember);
    const output = await useCase.execute({ id: castMember.cast_member_id.id });
    expect(output).toStrictEqual({
      id: castMember.cast_member_id.id,
      name: castMember.name,
      type: castMember.type.type,
      created_at: castMember.created_at,
    });
  });
});