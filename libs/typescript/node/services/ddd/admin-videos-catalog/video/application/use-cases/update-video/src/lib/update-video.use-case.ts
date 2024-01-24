import { CastMembersIdExistsInDatabaseValidator } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/validations';
import { CategoriesIdExistsInDatabaseValidator } from '@nodelib/services/ddd/admin-videos-catalog/category/application/validations';
import { GenresIdExistsInDatabaseValidator } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/validations';
import { IUseCase } from '@nodelib/shared/application';
import { NotFoundError } from '@nodelib/shared/errors';
import { IUnitOfWork } from '@nodelib/shared/ddd-utils/repository';
import { EntityValidationError } from '@nodelib/shared/validators';
import { Rating } from '@nodelib/shared/value-objects/rating';
import { Video, VideoId } from '@nodelib/services/ddd/admin-videos-catalog/video/entity';
import { IVideoRepository } from '@nodelib/services/ddd/admin-videos-catalog/video/repository';
import { UpdateVideoInput } from './update-video.input';

export class UpdateVideoUseCase
  implements IUseCase<UpdateVideoInput, UpdateVideoOutput>
{
  constructor(
    private uow: IUnitOfWork,
    private videoRepo: IVideoRepository,
    private categoriesIdValidator: CategoriesIdExistsInDatabaseValidator,
    private genresIdValidator: GenresIdExistsInDatabaseValidator,
    private castMembersIdValidator: CastMembersIdExistsInDatabaseValidator,
  ) {}

  async execute(input: UpdateVideoInput): Promise<UpdateVideoOutput> {
    const videoId = new VideoId(input.id);
    const video = await this.videoRepo.findById(videoId);

    if (!video) {
      throw new NotFoundError(input.id, Video);
    }

    input.title && video.changeTitle(input.title);
    input.description && video.changeDescription(input.description);
    input.year_launched && video.changeYearLaunched(input.year_launched);
    input.duration && video.changeDuration(input.duration);
    if (input.rating) {
      const [type, errorRating] = Rating.create(input.rating).asArray();

      video.changeRating(type);

      errorRating && video.notification.setError(errorRating.message, 'type');
    }

    if (input.is_opened === true) {
      video.markAsOpened();
    }

    if (input.is_opened === false) {
      video.markAsNotOpened();
    }

    const notification = video.notification;

    if (input.categories_id) {
      const [categoriesId, errorsCategoriesId] = (
        await this.categoriesIdValidator.validate(input.categories_id)
      ).asArray();

      categoriesId && video.syncCategoriesId(categoriesId);

      errorsCategoriesId &&
        notification.setError(
          errorsCategoriesId.map((e) => e.message),
          'categories_id',
        );
    }

    if (input.genres_id) {
      const [genresId, errorsGenresId] = (
        await this.genresIdValidator.validate(input.genres_id)
      ).asArray();

      genresId && video.syncGenresId(genresId);

      errorsGenresId &&
        notification.setError(
          errorsGenresId.map((e) => e.message),
          'genres_id',
        );
    }

    if (input.cast_members_id) {
      const [castMembersId, errorsCastMembersId] = (
        await this.castMembersIdValidator.validate(input.cast_members_id)
      ).asArray();

      castMembersId && video.syncCastMembersId(castMembersId);

      errorsCastMembersId &&
        notification.setError(
          errorsCastMembersId.map((e) => e.message),
          'cast_members_id',
        );
    }

    if (video.notification.hasErrors()) {
      throw new EntityValidationError(video.notification.toJSON());
    }

    await this.uow.do(async () => {
      return this.videoRepo.update(video);
    });

    return { id: video.video_id.id };
  }
}

export type UpdateVideoOutput = { id: string };