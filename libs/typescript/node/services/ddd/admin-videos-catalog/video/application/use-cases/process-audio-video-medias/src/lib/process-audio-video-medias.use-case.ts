import { IUseCase } from '@nodelib/shared/application';
import { NotFoundError } from '@nodelib/shared/errors';
import { IUnitOfWork } from '@nodelib/shared/ddd-utils/repository';
import { AudioVideoMediaStatus } from '@nodelib/shared/value-objects/audio-video-media';
import { Video, VideoId } from '@nodelib/services/ddd/admin-videos-catalog/video/entity';
import { IVideoRepository } from '@nodelib/services/ddd/admin-videos-catalog/video/repository';
import { ProcessAudioVideoMediasInput } from './process-audio-video-medias.input';

export class ProcessAudioVideoMediasUseCase
  implements
    IUseCase<ProcessAudioVideoMediasInput, ProcessAudioVideoMediasOutput>
{
  constructor(
    private uow: IUnitOfWork,
    private videoRepo: IVideoRepository,
  ) {}

  async execute(input: ProcessAudioVideoMediasInput) {
    const videoId = new VideoId(input.video_id);
    const video = await this.videoRepo.findById(videoId);

    if (!video) {
      throw new NotFoundError(input.video_id, Video);
    }

    if (input.field === 'trailer') {
      if (!video.trailer) {
        throw new Error('Trailer not found');
      }

      video.trailer =
        input.status === AudioVideoMediaStatus.COMPLETED
          ? video.trailer.complete(input.encoded_location)
          : video.trailer.fail();
    }

    if (input.field === 'video') {
      if (!video.video) {
        throw new Error('Video not found');
      }

      video.video =
        input.status === AudioVideoMediaStatus.COMPLETED
          ? video.video.complete(input.encoded_location)
          : video.video.fail();
    }

    this.uow.do(async () => {
      await this.videoRepo.update(video);
    });
  }
}

type ProcessAudioVideoMediasOutput = void;