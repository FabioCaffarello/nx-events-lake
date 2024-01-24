import { IUseCase, IStorage } from '@nodelib/shared/application';
import { NotFoundError } from '@nodelib/shared/errors';
import { IUnitOfWork } from '@nodelib/shared/ddd-utils/repository';
import { EntityValidationError } from '@nodelib/shared/validators';
import { Banner } from '@nodelib/services/ddd/admin-videos-catalog/video/value-objects/banner';
import { Thumbnail, ThumbnailHalf } from '@nodelib/services/ddd/admin-videos-catalog/video/value-objects/thumbnail';
import { Video, VideoId } from '@nodelib/services/ddd/admin-videos-catalog/video/entity';
import { IVideoRepository } from'@nodelib/services/ddd/admin-videos-catalog/video/repository';
import { UploadImageMediasInput } from './upload-image-medias.input';

export class UploadImageMediasUseCase
  implements IUseCase<UploadImageMediasInput, UploadImageMediasOutput>
{
  constructor(
    private uow: IUnitOfWork,
    private videoRepo: IVideoRepository,
    private storage: IStorage,
  ) {}

  async execute(
    input: UploadImageMediasInput,
  ): Promise<UploadImageMediasOutput> {
    const videoId = new VideoId(input.video_id);
    const video = await this.videoRepo.findById(videoId);

    if (!video) {
      throw new NotFoundError(input.video_id, Video);
    }

    const imagesMap = {
      banner: Banner,
      thumbnail: Thumbnail,
      thumbnail_half: ThumbnailHalf,
    };

    const [image, errorImage] = imagesMap[input.field]
      .createFromFile({
        ...input.file,
        video_id: videoId,
      })
      .asArray();

    if (errorImage) {
      throw new EntityValidationError([
        { [input.field]: [errorImage.message] },
      ]);
    }

    image instanceof Banner && video.replaceBanner(image);
    image instanceof Thumbnail && video.replaceThumbnail(image);
    image instanceof ThumbnailHalf && video.replaceThumbnailHalf(image);

    await this.storage.store({
      data: input.file.data,
      mime_type: input.file.mime_type,
      id: image.url,
    });

    await this.uow.do(async () => {
      await this.videoRepo.update(video);
    });
  }
}

export type UploadImageMediasOutput = void;