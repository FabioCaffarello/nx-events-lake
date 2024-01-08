import { OmitType } from '@nestjs/mapped-types';
import { UpdateCastMemberInput } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/use-cases/update-cast-member';

export class UpdateCastMemberInputWithoutId extends OmitType(
  UpdateCastMemberInput,
  ['id'] as any,
) {}

export class UpdateCastMemberDto extends UpdateCastMemberInputWithoutId {}