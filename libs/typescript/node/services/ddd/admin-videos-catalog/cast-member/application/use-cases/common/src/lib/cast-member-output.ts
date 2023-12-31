import { CastMember } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/entity';

export type CastMemberOutput = {
  id: string;
  name: string;
  type: number;
  created_at: Date;
};

export class CastMemberOutputMapper {
  static toOutput(entity: CastMember): CastMemberOutput {
    const { cast_member_id, ...other_props } = entity.toJSON();
    return {
      id: cast_member_id,
      ...other_props,
    };
  }
}