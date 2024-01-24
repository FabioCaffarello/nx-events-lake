import { ApplicationService, IStorage } from '@nodelib/shared/application';
import { IUseCase } from '@nodelib/shared/application';
import { NotFoundError } from '@nodelib/shared/errors';
import { EntityValidationError } from '@nodelib/shared/validators';
import { Trailer } from '@nodelib/services/ddd/admin-videos-catalog/video/value-objects/trailer';
import { VideoMedia } from '@nodelib/services/ddd/admin-videos-catalog/video/value-objects/video-media';
import { Video, VideoId } from '@nodelib/services/ddd/admin-videos-catalog/video/entity';
import { IVideoRepository } from '@nodelib/services/ddd/admin-videos-catalog/video/repository';
import { UploadAudioVideoMediaInput } from './upload-audio-video-media.input';

export class UploadAudioVideoMediasUseCase
  implements IUseCase<UploadAudioVideoMediaInput, UploadAudioVideoMediaOutput>
{
  constructor(
    private appService: ApplicationService,
    private videoRepo: IVideoRepository,
    private storage: IStorage,
  ) {}

  async execute(
    input: UploadAudioVideoMediaInput,
  ): Promise<UploadAudioVideoMediaOutput> {
    const video = await this.videoRepo.findById(new VideoId(input.video_id));
    if (!video) {
      throw new NotFoundError(input.video_id, Video);
    }

    const audioVideoMediaMap = {
      trailer: Trailer,
      video: VideoMedia,
    };

    const audioMediaClass = audioVideoMediaMap[input.field] as
      | typeof Trailer
      | typeof VideoMedia;
    const [audioVideoMedia, errorAudioMedia] = audioMediaClass
      .createFromFile({
        ...input.file,
        video_id: video.video_id,
      })
      .asArray();

    if (errorAudioMedia) {
      throw new EntityValidationError([
        {
          [input.field]: [errorAudioMedia.message],
        },
      ]);
    }

    audioVideoMedia instanceof Trailer && video.replaceTrailer(audioVideoMedia);
    audioVideoMedia instanceof VideoMedia &&
      video.replaceVideo(audioVideoMedia);

    await this.storage.store({
      data: input.file.data,
      id: audioVideoMedia.raw_url,
      mime_type: input.file.mime_type,
    });

    await this.appService.run(async () => {
      return this.videoRepo.update(video);
    });
  }
}

export type UploadAudioVideoMediaOutput = void;