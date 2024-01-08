import { IUseCase } from '@nodelib/shared/application';
import { CastMemberId } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/entity';
import { ICastMemberRepository } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/repository';

export class DeleteCastMemberUseCase
  implements IUseCase<DeleteCastMemberInput, DeleteCastMemberOutput>
{
  constructor(private castMemberRepository: ICastMemberRepository) {}

  async execute(input: DeleteCastMemberInput): Promise<DeleteCastMemberOutput> {
    const castMemberId = new CastMemberId(input.id);
    await this.castMemberRepository.delete(castMemberId);
  }
}

export type DeleteCastMemberInput = {
  id: string;
};

type DeleteCastMemberOutput = void;