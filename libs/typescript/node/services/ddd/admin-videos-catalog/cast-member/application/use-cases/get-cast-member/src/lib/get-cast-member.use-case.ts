import { IUseCase } from '@nodelib/shared/application';
import { NotFoundError } from '@nodelib/shared/errors';
import {
  CastMember,
  CastMemberId,
} from '@nodelib/services/ddd/admin-videos-catalog/cast-member/entity';
import { ICastMemberRepository } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/repository';
import {
  CastMemberOutput,
  CastMemberOutputMapper,
} from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/use-cases/common';

export class GetCastMemberUseCase
  implements IUseCase<GetCastMemberInput, GetCastMemberOutput>
{
  constructor(private castMemberRepo: ICastMemberRepository) {}

  async execute(input: GetCastMemberInput): Promise<GetCastMemberOutput> {
    const castMemberId = new CastMemberId(input.id);
    const castMember = await this.castMemberRepo.findById(castMemberId);
    if (!castMember) {
      throw new NotFoundError(input.id, CastMember);
    }

    return CastMemberOutputMapper.toOutput(castMember);
  }
}

export type GetCastMemberInput = {
  id: string;
};

export type GetCastMemberOutput = CastMemberOutput;