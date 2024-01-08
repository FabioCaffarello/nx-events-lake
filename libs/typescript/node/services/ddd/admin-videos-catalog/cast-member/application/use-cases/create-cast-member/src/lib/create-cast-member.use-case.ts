import { IUseCase } from '@nodelib/shared/application';
import { ICastMemberRepository } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/repository';
import {
  CastMemberOutput,
  CastMemberOutputMapper,
} from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/use-cases/common';
import { CastMember, CastMemberType } from'@nodelib/services/ddd/admin-videos-catalog/cast-member/entity';
import { EntityValidationError } from '@nodelib/shared/validators';
import { CreateCastMemberInput } from './create-cast-member.input';

export class CreateCastMemberUseCase
  implements IUseCase<CreateCastMemberInput, CreateCastMemberOutput>
{
  constructor(private castMemberRepo: ICastMemberRepository) {}

  async execute(input: CreateCastMemberInput): Promise<CastMemberOutput> {
    const [type, errorCastMemberType] = CastMemberType.create(
      input.type,
    ).asArray();
    const entity = CastMember.create({
      ...input,
      type,
    });
    const notification = entity.notification;
    if (errorCastMemberType) {
      notification.setError(errorCastMemberType.message, 'type');
    }

    if (notification.hasErrors()) {
      throw new EntityValidationError(notification.toJSON());
    }

    await this.castMemberRepo.insert(entity);
    return CastMemberOutputMapper.toOutput(entity);
  }
}

export type CreateCastMemberOutput = CastMemberOutput;