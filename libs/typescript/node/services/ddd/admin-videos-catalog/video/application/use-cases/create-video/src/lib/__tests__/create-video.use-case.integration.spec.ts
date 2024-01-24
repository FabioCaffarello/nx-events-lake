import { CastMembersIdExistsInDatabaseValidator } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/validations';
import { CastMember } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/entity';
import {
  CastMemberModel,
  CastMemberSequelizeRepository,
} from '@nodelib/services/ddd/admin-videos-catalog/cast-member/infra/db/sequelize';
import { CategoriesIdExistsInDatabaseValidator } from  '@nodelib/services/ddd/admin-videos-catalog/category/application/validations';
import { Category } from '@nodelib/services/ddd/admin-videos-catalog/category/entity';
import { CategoryModel, CategorySequelizeRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/infra/db/sequelize';
import { GenresIdExistsInDatabaseValidator } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/validations';
import { Genre } from '@nodelib/services/ddd/admin-videos-catalog/genre/entity';
import { GenreModel, GenreSequelizeRepository } from '@nodelib/services/ddd/admin-videos-catalog/genre/infra/db/sequelize';
import { UnitOfWorkSequelize } from '@nodelib/shared/ddd-utils/infra/db/sequelize';
import { RatingValues } from '@nodelib/shared/value-objects/rating';
import { Video, VideoId } from '@nodelib/services/ddd/admin-videos-catalog/video/entity';
import { setupSequelizeForVideo } from '@nodelib/services/ddd/admin-videos-catalog/video/infra/db/testing';
import { VideoModel, VideoSequelizeRepository } from '@nodelib/services/ddd/admin-videos-catalog/video/infra/db/sequelize';
import { CreateVideoUseCase } from '../create-video.use-case';

import { DatabaseError } from 'sequelize';
describe('CreateVideoUseCase Integration Tests', () => {
  let uow: UnitOfWorkSequelize;
  let useCase: CreateVideoUseCase;

  let videoRepo: VideoSequelizeRepository;
  let genreRepo: GenreSequelizeRepository;
  let castMemberRepo: CastMemberSequelizeRepository;

  let categoryRepo: CategorySequelizeRepository;
  let categoriesIdsValidator: CategoriesIdExistsInDatabaseValidator;
  let genresIdsValidator: GenresIdExistsInDatabaseValidator;
  let castMembersIdsValidator: CastMembersIdExistsInDatabaseValidator;

  const sequelizeHelper = setupSequelizeForVideo();

  beforeEach(() => {
    uow = new UnitOfWorkSequelize(sequelizeHelper.sequelize);
    videoRepo = new VideoSequelizeRepository(VideoModel, uow);
    genreRepo = new GenreSequelizeRepository(GenreModel, uow);
    categoryRepo = new CategorySequelizeRepository(CategoryModel);
    castMemberRepo = new CastMemberSequelizeRepository(CastMemberModel);
    categoriesIdsValidator = new CategoriesIdExistsInDatabaseValidator(
      categoryRepo,
    );
    genresIdsValidator = new GenresIdExistsInDatabaseValidator(genreRepo);
    castMembersIdsValidator = new CastMembersIdExistsInDatabaseValidator(
      castMemberRepo,
    );
    useCase = new CreateVideoUseCase(
      uow,
      videoRepo,
      categoriesIdsValidator,
      genresIdsValidator,
      castMembersIdsValidator,
    );
  });

  it('should create a video', async () => {
    const categories = Category.fake().theCategories(2).build();
    await categoryRepo.bulkInsert(categories);
    const categoriesId = categories.map((c) => c.category_id.id);

    const genres = Genre.fake().theGenres(2).build();
    genres[0].syncCategoriesId([categories[0].category_id]);
    genres[1].syncCategoriesId([categories[1].category_id]);
    await genreRepo.bulkInsert(genres);
    const genresId = genres.map((c) => c.genre_id.id);

    const castMembers = CastMember.fake().theCastMembers(2).build();
    await castMemberRepo.bulkInsert(castMembers);
    const castMembersId = castMembers.map((c) => c.cast_member_id.id);

    const output = await useCase.execute({
      title: 'test video',
      description: 'test description',
      year_launched: 2021,
      duration: 90,
      rating: RatingValues.R10,
      is_opened: true,
      categories_id: categoriesId,
      genres_id: genresId,
      cast_members_id: castMembersId,
    });
    expect(output).toStrictEqual({
      id: expect.any(String),
    });
    const video = await videoRepo.findById(new VideoId(output.id));
    expect(video!.toJSON()).toStrictEqual({
      video_id: expect.any(String),
      title: 'test video',
      description: 'test description',
      year_launched: 2021,
      duration: 90,
      rating: RatingValues.R10,
      is_opened: true,
      is_published: false,
      banner: null,
      thumbnail: null,
      thumbnail_half: null,
      trailer: null,
      video: null,
      categories_id: expect.arrayContaining(categoriesId),
      genres_id: expect.arrayContaining(genresId),
      cast_members_id: expect.arrayContaining(castMembersId),
      created_at: expect.any(Date),
    });
  });

  it('rollback transaction', async () => {
    const categories = Category.fake().theCategories(2).build();
    await categoryRepo.bulkInsert(categories);
    const categoriesId = categories.map((c) => c.category_id.id);

    const genres = Genre.fake().theGenres(2).build();
    genres[0].syncCategoriesId([categories[0].category_id]);
    genres[1].syncCategoriesId([categories[1].category_id]);
    await genreRepo.bulkInsert(genres);
    const genresId = genres.map((c) => c.genre_id.id);

    const castMembers = CastMember.fake().theCastMembers(2).build();
    await castMemberRepo.bulkInsert(castMembers);
    const castMembersId = castMembers.map((c) => c.cast_member_id.id);

    const video = Video.fake().aVideoWithoutMedias().build();
    video.title = 't'.repeat(256);

    const mockCreate = jest
      .spyOn(Video, 'create')
      .mockImplementation(() => video);

    await expect(
      useCase.execute({
        title: 'test video',
        rating: RatingValues.R10,
        categories_id: categoriesId,
        genres_id: genresId,
        cast_members_id: castMembersId,
      } as any),
    ).rejects.toThrowError(DatabaseError);

    const videos = await videoRepo.findAll();
    expect(videos.length).toEqual(0);

    mockCreate.mockRestore();
  });
});