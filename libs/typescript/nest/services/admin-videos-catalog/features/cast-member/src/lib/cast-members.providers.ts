import { getModelToken } from '@nestjs/sequelize';
import { CastMemberInMemoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/infra/db/in-memory';
import { CreateCastMemberUseCase } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/use-cases/create-cast-member';
import { UpdateCastMemberUseCase } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/use-cases/update-cast-member';
import { ListCastMembersUseCase } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/use-cases/list-cast-members';
import { GetCastMemberUseCase } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/use-cases/get-cast-member';
import { DeleteCastMemberUseCase } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/use-cases/delete-cast-member';
import {
  CastMemberModel,
  CastMemberSequelizeRepository,
} from '@nodelib/services/ddd/admin-videos-catalog/cast-member/infra/db/sequelize';
import { ICastMemberRepository } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/repository';
import { CastMembersIdExistsInDatabaseValidator } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/validations';

export const REPOSITORIES = {
  CAST_MEMBER_REPOSITORY: {
    provide: 'CastMemberRepository',
    useExisting: CastMemberSequelizeRepository,
  },
  CAST_MEMBER_IN_MEMORY_REPOSITORY: {
    provide: CastMemberInMemoryRepository,
    useClass: CastMemberInMemoryRepository,
  },
  CAST_MEMBER_SEQUELIZE_REPOSITORY: {
    provide: CastMemberSequelizeRepository,
    useFactory: (castMemberModel: typeof CastMemberModel) => {
      return new CastMemberSequelizeRepository(castMemberModel);
    },
    inject: [getModelToken(CastMemberModel)],
  },
};

export const USE_CASES = {
  CREATE_CAST_MEMBER_USE_CASE: {
    provide: CreateCastMemberUseCase,
    useFactory: (castMemberRepo: ICastMemberRepository) => {
      return new CreateCastMemberUseCase(castMemberRepo);
    },
    inject: [REPOSITORIES.CAST_MEMBER_REPOSITORY.provide],
  },
  UPDATE_CAST_MEMBER_USE_CASE: {
    provide: UpdateCastMemberUseCase,
    useFactory: (castMemberRepo: ICastMemberRepository) => {
      return new UpdateCastMemberUseCase(castMemberRepo);
    },
    inject: [REPOSITORIES.CAST_MEMBER_REPOSITORY.provide],
  },
  LIST_CAST_MEMBERS_USE_CASE: {
    provide: ListCastMembersUseCase,
    useFactory: (castMemberRepo: ICastMemberRepository) => {
      return new ListCastMembersUseCase(castMemberRepo);
    },
    inject: [REPOSITORIES.CAST_MEMBER_REPOSITORY.provide],
  },
  GET_CAST_MEMBER_USE_CASE: {
    provide: GetCastMemberUseCase,
    useFactory: (castMemberRepo: ICastMemberRepository) => {
      return new GetCastMemberUseCase(castMemberRepo);
    },
    inject: [REPOSITORIES.CAST_MEMBER_REPOSITORY.provide],
  },
  DELETE_CAST_MEMBER_USE_CASE: {
    provide: DeleteCastMemberUseCase,
    useFactory: (castMemberRepo: ICastMemberRepository) => {
      return new DeleteCastMemberUseCase(castMemberRepo);
    },
    inject: [REPOSITORIES.CAST_MEMBER_REPOSITORY.provide],
  },
};

export const VALIDATIONS = {
  CAST_MEMBERS_IDS_EXISTS_IN_DATABASE_VALIDATOR: {
    provide: CastMembersIdExistsInDatabaseValidator,
    useFactory: (castMemberRepo: ICastMemberRepository) => {
      return new CastMembersIdExistsInDatabaseValidator(castMemberRepo);
    },
    inject: [REPOSITORIES.CAST_MEMBER_REPOSITORY.provide],
  },
};

export const CAST_MEMBERS_PROVIDERS = {
  REPOSITORIES,
  USE_CASES,
  VALIDATIONS,
};