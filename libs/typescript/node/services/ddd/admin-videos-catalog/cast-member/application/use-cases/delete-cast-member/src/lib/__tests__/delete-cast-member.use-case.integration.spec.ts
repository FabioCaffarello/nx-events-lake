import { DeleteCastMemberUseCase } from '../delete-cast-member.use-case';
import { setupSequelize } from '@nodelib/shared/ddd-utils/infra/testing';
import { NotFoundError } from '@nodelib/shared/errors';
import {
  CastMember,
  CastMemberId,
} from '@nodelib/services/ddd/admin-videos-catalog/cast-member/entity';
import {
  CastMemberModel,
  CastMemberSequelizeRepository,
} from '@nodelib/services/ddd/admin-videos-catalog/cast-member/infra/db/sequelize';

describe('DeleteCastMemberUseCase Integration Tests', () => {
  let useCase: DeleteCastMemberUseCase;
  let repository: CastMemberSequelizeRepository;

  setupSequelize({ models: [CastMemberModel] });

  beforeEach(() => {
    repository = new CastMemberSequelizeRepository(CastMemberModel);
    useCase = new DeleteCastMemberUseCase(repository);
  });

  it('should throws error when entity not found', async () => {
    const castMemberId = new CastMemberId();
    await expect(() =>
      useCase.execute({ id: castMemberId.id }),
    ).rejects.toThrow(new NotFoundError(castMemberId.id, CastMember));
  });

  it('should delete a cast member', async () => {
    const castMember = CastMember.fake().anActor().build();
    await repository.insert(castMember);
    await useCase.execute({
      id: castMember.cast_member_id.id,
    });
    const noHasModel = await CastMemberModel.findByPk(
      castMember.cast_member_id.id,
    );
    expect(noHasModel).toBeNull();
  });
});