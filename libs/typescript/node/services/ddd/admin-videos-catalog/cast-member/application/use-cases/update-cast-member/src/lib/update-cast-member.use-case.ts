import { IUseCase } from '@nodelib/shared/application';
import { NotFoundError } from '@nodelib/shared/errors';
import { EntityValidationError } from '@nodelib/shared/validators';
import {
  CastMember,
  CastMemberId,
  CastMemberType,
} from '@nodelib/services/ddd/admin-videos-catalog/cast-member/entity';
import { ICastMemberRepository } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/repository';
import {
  CastMemberOutput,
  CastMemberOutputMapper,
} from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/use-cases/common';
import { UpdateCastMemberInput } from './update-cast-member.input';

export class UpdateCastMemberUseCase
  implements IUseCase<UpdateCastMemberInput, UpdateCastMemberOutput>
{
  constructor(private castMemberRepo: ICastMemberRepository) {}

  async execute(input: UpdateCastMemberInput): Promise<UpdateCastMemberOutput> {
    const castMemberId = new CastMemberId(input.id);
    const castMember = await this.castMemberRepo.findById(castMemberId);

    if (!castMember) {
      throw new NotFoundError(input.id, CastMember);
    }

    input.name && castMember.changeName(input.name);

    if (input.type) {
      const [type, errorCastMemberType] = CastMemberType.create(
        input.type,
      ).asArray();

      castMember.changeType(type);

      errorCastMemberType &&
        castMember.notification.setError(errorCastMemberType.message, 'type');
    }

    if (castMember.notification.hasErrors()) {
      throw new EntityValidationError(castMember.notification.toJSON());
    }

    await this.castMemberRepo.update(castMember);

    return CastMemberOutputMapper.toOutput(castMember);
  }
}

export type UpdateCastMemberOutput = CastMemberOutput;