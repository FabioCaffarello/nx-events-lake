import { Either } from '@nodelib/shared/ddd-utils/either';
import { NotFoundError } from '@nodelib/shared/errors';
import { CastMember, CastMemberId } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/entity';
import { ICastMemberRepository } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/repository';


export class CastMembersIdExistsInDatabaseValidator {
    constructor(private castMemberRepo: ICastMemberRepository) {}
  
    async validate(
      cast_members_id: string[],
    ): Promise<Either<CastMemberId[], NotFoundError[]>> {
      const castMembersId = cast_members_id.map((v) => new CastMemberId(v));
  
      const existsResult = await this.castMemberRepo.existsById(castMembersId);
      return existsResult.not_exists.length > 0
        ? Either.fail(
            existsResult.not_exists.map(
              (c) => new NotFoundError(c.id, CastMember),
            ),
          )
        : Either.ok(castMembersId);
    }
  }