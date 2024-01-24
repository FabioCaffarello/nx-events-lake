import { Test, TestingModule } from '@nestjs/testing';
import { DatabaseModule } from '@nestlib/services/admin-videos-catalog/database';
import { VideosModule } from '../videos.module';
import { UnitOfWorkSequelize } from '@nodelib/shared/ddd-utils/infra/db/sequelize';
import { ConfigModule } from '@nestlib/services/admin-videos-catalog/config-setup';
import { SharedModule } from '@nestlib/shared/module';
import { EventModule } from '@nestlib/services/admin-videos-catalog/event';
import { VideoAudioMediaUploadedIntegrationEvent } from '@nodelib/services/ddd/admin-videos-catalog/video/events';
import { AmqpConnection } from '@golevelup/nestjs-rabbitmq';
import { UseCaseModule } from '@nestlib/services/admin-videos-catalog/use-case';
import { RabbitmqModule } from '@nestlib/services/admin-videos-catalog/rabbitmq';
import { VIDEOS_PROVIDERS } from '../videos.providers';
import { IVideoRepository } from '@nodelib/services/ddd/admin-videos-catalog/video/repository';
import { Video } from '@nodelib/services/ddd/admin-videos-catalog/video/entity';
import { Category } from '@nodelib/services/ddd/admin-videos-catalog/category/entity';
import { Genre } from '@nodelib/services/ddd/admin-videos-catalog/genre/entity';
import { CastMember } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/entity';
import { ICategoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/repository';
import { CATEGORY_PROVIDERS } from '@nestlib/services/admin-videos-catalog/features/categories';
import { IGenreRepository } from '@nodelib/services/ddd/admin-videos-catalog/genre/repository';
import { GENRES_PROVIDERS } from '@nestlib/services/admin-videos-catalog/features/genres';
import { ICastMemberRepository } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/repository';
import { CAST_MEMBERS_PROVIDERS } from '@nestlib/services/admin-videos-catalog/features/cast-member';
import { UploadAudioVideoMediasUseCase } from '@nodelib/services/ddd/admin-videos-catalog/video/application/use-cases/upload-audio-video-medias';
import { getConnectionToken } from '@nestjs/sequelize';
import { Sequelize } from 'sequelize-typescript';
import { EVENTS_MESSAGE_BROKER_CONFIG } from '@nodelib/shared/ddd-utils/infra/message-broker';
import { ChannelWrapper } from 'amqp-connection-manager';
import { ConsumeMessage } from 'amqplib';
import { AuthModule } from '@nestlib/services/admin-videos-catalog/auth';

describe('PublishVideoMediaReplacedInQueueHandler Integration Tests', () => {
  let module: TestingModule;
  let channelWrapper: ChannelWrapper;
  beforeEach(async () => {
    module = await Test.createTestingModule({
      imports: [
        ConfigModule.forRoot(),
        SharedModule,
        DatabaseModule,
        EventModule,
        UseCaseModule,
        RabbitmqModule.forRoot(),
        AuthModule,
        VideosModule,
      ],
    })
      .overrideProvider('UnitOfWork')
      .useFactory({
        factory: (sequelize: Sequelize) => {
          return new UnitOfWorkSequelize(sequelize);
        },
        inject: [getConnectionToken()],
      })
      .compile();
    await module.init();

    const amqpConn = module.get<AmqpConnection>(AmqpConnection);
    channelWrapper = amqpConn.managedConnection.createChannel();
    await channelWrapper.addSetup((channel) => {
      return Promise.all([
        channel.assertQueue('test-queue-video-upload', {
          durable: false,
        }),
        channel.bindQueue(
          'test-queue-video-upload',
          EVENTS_MESSAGE_BROKER_CONFIG[
            VideoAudioMediaUploadedIntegrationEvent.name
          ].exchange,
          EVENTS_MESSAGE_BROKER_CONFIG[
            VideoAudioMediaUploadedIntegrationEvent.name
          ].routing_key,
        ),
      ]).then(() => channel.purgeQueue('test-queue-video-upload'));
    });
  });

  afterEach(async () => {
    await channelWrapper.close();
    await module.close();
  });

  it('should publish video media replaced event in queue', async () => {
    const category = Category.fake().aCategory().build();
    const genre = Genre.fake()
      .aGenre()
      .addCategoryId(category.category_id)
      .build();
    const castMember = CastMember.fake().aDirector().build();
    const video = Video.fake()
      .aVideoWithoutMedias()
      .addCategoryId(category.category_id)
      .addGenreId(genre.genre_id)
      .addCastMemberId(castMember.cast_member_id)
      .build();

    const categoryRepo: ICategoryRepository = module.get(
      CATEGORY_PROVIDERS.REPOSITORIES.CATEGORY_REPOSITORY.provide,
    );
    await categoryRepo.insert(category);

    const genreRepo: IGenreRepository = module.get(
      GENRES_PROVIDERS.REPOSITORIES.GENRE_REPOSITORY.provide,
    );
    await genreRepo.insert(genre);

    const castMemberRepo: ICastMemberRepository = module.get(
      CAST_MEMBERS_PROVIDERS.REPOSITORIES.CAST_MEMBER_REPOSITORY.provide,
    );
    await castMemberRepo.insert(castMember);

    const videoRepo: IVideoRepository = module.get(
      VIDEOS_PROVIDERS.REPOSITORIES.VIDEO_REPOSITORY.provide,
    );
    await videoRepo.insert(video);

    const useCase: UploadAudioVideoMediasUseCase = module.get(
      VIDEOS_PROVIDERS.USE_CASES.UPLOAD_AUDIO_VIDEO_MEDIA_USE_CASE.provide,
    );

    await useCase.execute({
      video_id: video.video_id.id,
      field: 'video',
      file: {
        data: Buffer.from('data'),
        mime_type: 'video/mp4',
        raw_name: 'video.mp4',
        size: 100,
      },
    });

    const msg: ConsumeMessage = await new Promise((resolve) => {
      channelWrapper.consume('test-queue-video-upload', (msg) => {
        resolve(msg);
      });
    });

    const msgObj = JSON.parse(msg.content.toString());
    const updatedVideo = await videoRepo.findById(video.video_id);
    expect(msgObj).toEqual({
      resource_id: `${video.video_id.id}.video`,
      file_path: updatedVideo?.video?.raw_url,
    });
  });
});