import { getModelToken } from '@nestjs/sequelize';
import { VideoInMemoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/video/infra/db/in-memory';
import { IVideoRepository } from '@nodelib/services/ddd/admin-videos-catalog/video/repository';
import { UnitOfWorkSequelize } from '@nodelib/shared/ddd-utils/infra/db/sequelize';
import { IUnitOfWork } from '@nodelib/shared/ddd-utils/repository';
import { CATEGORY_PROVIDERS } from '@nestlib/services/admin-videos-catalog/features/categories';
import { VideoModel, VideoSequelizeRepository } from '@nodelib/services/ddd/admin-videos-catalog/video/infra/db/sequelize';
import { GENRES_PROVIDERS } from '@nestlib/services/admin-videos-catalog/features/genres';
import { CAST_MEMBERS_PROVIDERS } from '@nestlib/services/admin-videos-catalog/features/cast-member';
import { IGenreRepository } from '@nodelib/services/ddd/admin-videos-catalog/genre/repository';
import { ICastMemberRepository } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/repository';
import { ICategoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/repository';
import { ApplicationService, IStorage, IMessageBroker } from '@nodelib/shared/application';
import { CreateVideoUseCase } from '@nodelib/services/ddd/admin-videos-catalog/video/application/use-cases/create-video';
import { CategoriesIdExistsInDatabaseValidator } from '@nodelib/services/ddd/admin-videos-catalog/category/application/validations';
import { GenresIdExistsInDatabaseValidator } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/validations';
import { CastMembersIdExistsInDatabaseValidator } from'@nodelib/services/ddd/admin-videos-catalog/cast-member/application/validations';
import { UpdateVideoUseCase } from '@nodelib/services/ddd/admin-videos-catalog/video/application/use-cases/update-video';
import { UploadAudioVideoMediasUseCase } from '@nodelib/services/ddd/admin-videos-catalog/video/application/use-cases/upload-audio-video-medias';
import { GetVideoUseCase } from '@nodelib/services/ddd/admin-videos-catalog/video/application/use-cases/get-video';
import { ProcessAudioVideoMediasUseCase } from '@nodelib/services/ddd/admin-videos-catalog/video/application/use-cases/process-audio-video-medias';
import { PublishVideoMediaReplacedInQueueHandler } from '@nodelib/services/ddd/admin-videos-catalog/video/application/handlers';

export const REPOSITORIES = {
  VIDEO_REPOSITORY: {
    provide: 'VideoRepository',
    useExisting: VideoSequelizeRepository,
  },
  VIDEO_IN_MEMORY_REPOSITORY: {
    provide: VideoInMemoryRepository,
    useClass: VideoInMemoryRepository,
  },
  VIDEO_SEQUELIZE_REPOSITORY: {
    provide: VideoSequelizeRepository,
    useFactory: (videoModel: typeof VideoModel, uow: UnitOfWorkSequelize) => {
      return new VideoSequelizeRepository(videoModel, uow);
    },
    inject: [getModelToken(VideoModel), 'UnitOfWork'],
  },
};

export const USE_CASES = {
  CREATE_VIDEO_USE_CASE: {
    provide: CreateVideoUseCase,
    useFactory: (
      uow: IUnitOfWork,
      videoRepo: IVideoRepository,
      categoriesIdValidator: CategoriesIdExistsInDatabaseValidator,
      genresIdValidator: GenresIdExistsInDatabaseValidator,
      castMembersIdValidator: CastMembersIdExistsInDatabaseValidator,
    ) => {
      return new CreateVideoUseCase(
        uow,
        videoRepo,
        categoriesIdValidator,
        genresIdValidator,
        castMembersIdValidator,
      );
    },
    inject: [
      'UnitOfWork',
      REPOSITORIES.VIDEO_REPOSITORY.provide,
      CATEGORY_PROVIDERS.VALIDATIONS.CATEGORIES_IDS_EXISTS_IN_DATABASE_VALIDATOR
        .provide,
      GENRES_PROVIDERS.VALIDATIONS.GENRES_IDS_EXISTS_IN_DATABASE_VALIDATOR
        .provide,
      CAST_MEMBERS_PROVIDERS.VALIDATIONS
        .CAST_MEMBERS_IDS_EXISTS_IN_DATABASE_VALIDATOR.provide,
    ],
  },
  UPDATE_VIDEO_USE_CASE: {
    provide: UpdateVideoUseCase,
    useFactory: (
      uow: IUnitOfWork,
      videoRepo: IVideoRepository,
      categoriesIdValidator: CategoriesIdExistsInDatabaseValidator,
      genresIdValidator: GenresIdExistsInDatabaseValidator,
      castMembersIdValidator: CastMembersIdExistsInDatabaseValidator,
    ) => {
      return new UpdateVideoUseCase(
        uow,
        videoRepo,
        categoriesIdValidator,
        genresIdValidator,
        castMembersIdValidator,
      );
    },
    inject: [
      'UnitOfWork',
      REPOSITORIES.VIDEO_REPOSITORY.provide,
      CATEGORY_PROVIDERS.VALIDATIONS.CATEGORIES_IDS_EXISTS_IN_DATABASE_VALIDATOR
        .provide,
      GENRES_PROVIDERS.VALIDATIONS.GENRES_IDS_EXISTS_IN_DATABASE_VALIDATOR
        .provide,
      CAST_MEMBERS_PROVIDERS.VALIDATIONS
        .CAST_MEMBERS_IDS_EXISTS_IN_DATABASE_VALIDATOR.provide,
    ],
  },
  UPLOAD_AUDIO_VIDEO_MEDIA_USE_CASE: {
    provide: UploadAudioVideoMediasUseCase,
    useFactory: (
      appService: ApplicationService,
      videoRepo: IVideoRepository,
      storage: IStorage,
    ) => {
      return new UploadAudioVideoMediasUseCase(appService, videoRepo, storage);
    },
    inject: [
      ApplicationService,
      REPOSITORIES.VIDEO_REPOSITORY.provide,
      'IStorage',
    ],
  },
  GET_VIDEO_USE_CASE: {
    provide: GetVideoUseCase,
    useFactory: (
      videoRepo: IVideoRepository,
      categoryRepo: ICategoryRepository,
      genreRepo: IGenreRepository,
      castMemberRepo: ICastMemberRepository,
    ) => {
      return new GetVideoUseCase(
        videoRepo,
        categoryRepo,
        genreRepo,
        castMemberRepo,
      );
    },
    inject: [
      REPOSITORIES.VIDEO_REPOSITORY.provide,
      CATEGORY_PROVIDERS.REPOSITORIES.CATEGORY_REPOSITORY.provide,
      GENRES_PROVIDERS.REPOSITORIES.GENRE_REPOSITORY.provide,
      CAST_MEMBERS_PROVIDERS.REPOSITORIES.CAST_MEMBER_REPOSITORY.provide,
    ],
  },
  PROCESS_AUDIO_VIDEO_MEDIA_USE_CASE: {
    provide: ProcessAudioVideoMediasUseCase,
    useFactory: (uow: IUnitOfWork, videoRepo: IVideoRepository) => {
      return new ProcessAudioVideoMediasUseCase(uow, videoRepo);
    },
    inject: ['UnitOfWork', REPOSITORIES.VIDEO_REPOSITORY.provide],
  },
};

export const HANDLERS = {
  PUBLISH_VIDEO_MEDIA_REPLACED_IN_QUEUE_HANDLER: {
    provide: PublishVideoMediaReplacedInQueueHandler,
    useFactory: (messageBroker: IMessageBroker) => {
      return new PublishVideoMediaReplacedInQueueHandler(messageBroker);
    },
    inject: ['IMessageBroker'],
  },
};

export const VIDEOS_PROVIDERS = {
  REPOSITORIES,
  USE_CASES,
  HANDLERS,
};