import { Test, TestingModule } from '@nestjs/testing';
import { ICastMemberRepository } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/repository';
import { CastMembersController } from '../cast-members.controller';
import { CastMembersModule } from '../cast-members.module';
import { CreateCastMemberUseCase } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/use-cases/create-cast-member';
import { UpdateCastMemberUseCase } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/use-cases/update-cast-member';
import { ListCastMembersUseCase } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/use-cases/list-cast-members';
import { GetCastMemberUseCase } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/use-cases/get-cast-member';
import { DeleteCastMemberUseCase } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/use-cases/delete-cast-member';
import { CastMember } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/entity';
import { Uuid } from '@nodelib/shared/value-objects/uuid';
import { CastMemberCollectionPresenter } from '../cast-members.presenter';
import * as CastMemberProviders from '../cast-members.providers';
import {
  CreateCastMemberFixture,
  ListCastMembersFixture,
  UpdateCastMemberFixture,
} from '../testing/cast-member-fixtures';
import { DatabaseModule } from '@nestlib/services/admin-videos-catalog/database';
import { ConfigModule } from '@nestlib/services/admin-videos-catalog/config-setup';
import { CastMemberOutputMapper } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/use-cases/common';

describe('CastMembersController Integration Tests', () => {
  let controller: CastMembersController;
  let repository: ICastMemberRepository;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      imports: [ConfigModule.forRoot(), DatabaseModule, CastMembersModule],
    }).compile();

    controller = module.get(CastMembersController);
    repository = module.get(
      CastMemberProviders.REPOSITORIES.CAST_MEMBER_REPOSITORY.provide,
    );
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
    expect(controller['createUseCase']).toBeInstanceOf(CreateCastMemberUseCase);
    expect(controller['updateUseCase']).toBeInstanceOf(UpdateCastMemberUseCase);
    expect(controller['listUseCase']).toBeInstanceOf(ListCastMembersUseCase);
    expect(controller['getUseCase']).toBeInstanceOf(GetCastMemberUseCase);
    expect(controller['deleteUseCase']).toBeInstanceOf(DeleteCastMemberUseCase);
  });

  describe('should create a cast member', () => {
    const arrange = CreateCastMemberFixture.arrangeForCreate();

    test.each(arrange)(
      'when body is $send_data',
      async ({ send_data, expected }) => {
        const presenter = await controller.create(send_data);
        const entity = await repository.findById(new Uuid(presenter.id));

        expect(entity!.toJSON()).toStrictEqual({
          cast_member_id: presenter.id,
          created_at: presenter.created_at,
          ...expected,
        });

        expect(presenter).toEqual(
          CastMembersController.serialize(
            CastMemberOutputMapper.toOutput(entity!),
          ),
        );
      },
    );
  });

  describe('should update a cast member', () => {
    const arrange = UpdateCastMemberFixture.arrangeForUpdate();

    const castMember = CastMember.fake().anActor().build();
    beforeEach(async () => {
      await repository.insert(castMember);
    });

    test.each(arrange)(
      'with request $send_data',
      async ({ send_data, expected }) => {
        const presenter = await controller.update(
          castMember.cast_member_id.id,
          send_data,
        );
        const entity = await repository.findById(new Uuid(presenter.id));

        expect(entity!.toJSON()).toStrictEqual({
          cast_member_id: presenter.id,
          created_at: presenter.created_at,
          name: expected.name ?? castMember.name,
          type: expected.type ?? castMember.type.type,
        });
        expect(presenter).toEqual(
          CastMembersController.serialize(
            CastMemberOutputMapper.toOutput(entity!),
          ),
        );
      },
    );
  });

  it('should delete a cast member', async () => {
    const castMember = CastMember.fake().anActor().build();
    await repository.insert(castMember);
    const response = await controller.remove(castMember.entity_id.id);
    expect(response).not.toBeDefined();
    await expect(
      repository.findById(castMember.cast_member_id),
    ).resolves.toBeNull();
  });

  it('should get a cast member', async () => {
    const castMember = CastMember.fake().anActor().build();
    await repository.insert(castMember);
    const presenter = await controller.findOne(castMember.cast_member_id.id);
    expect(presenter.id).toBe(castMember.cast_member_id.id);
    expect(presenter.name).toBe(castMember.name);
    expect(presenter.type).toBe(castMember.type.type);
    expect(presenter.created_at).toEqual(castMember.created_at);
  });

  describe('search method', () => {
    describe('should returns cast members using query empty ordered by created_at', () => {
      const { entitiesMap, arrange } =
        ListCastMembersFixture.arrangeIncrementedWithCreatedAt();

      beforeEach(async () => {
        await repository.bulkInsert(Object.values(entitiesMap));
      });

      test.each(arrange)(
        'when send_data is $send_data',
        async ({ send_data, expected }) => {
          const presenter = await controller.search(send_data);
          const { entities, ...paginationProps } = expected;
          expect(presenter).toEqual(
            new CastMemberCollectionPresenter({
              items: entities.map(CastMemberOutputMapper.toOutput),
              ...paginationProps.meta,
            }),
          );
        },
      );
    });

    describe('should returns output using pagination, sort and filter', () => {
      const { entitiesMap, arrange } = ListCastMembersFixture.arrangeUnsorted();

      beforeEach(async () => {
        await repository.bulkInsert(Object.values(entitiesMap));
      });

      test.each(arrange)(
        'when send_data is {"filter": $send_data.filter, "page": $send_data.page, "per_page": $send_data.per_page, "sort": $send_data.sort, "sort_dir": $send_data.sort_dir}',
        async ({ send_data, expected }) => {
          const presenter = await controller.search(send_data);
          const { entities, ...paginationProps } = expected;
          expect(presenter).toEqual(
            new CastMemberCollectionPresenter({
              items: entities.map(CastMemberOutputMapper.toOutput),
              ...paginationProps.meta,
            }),
          );
        },
      );
    });
  });
});