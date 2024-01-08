import { Either } from '@nodelib/shared/ddd-utils/either';
import { MediaFileValidator } from '@nodelib/shared/validators';
import {
  AudioVideoMedia,
  AudioVideoMediaStatus,
} from '@nodelib/shared/value-objects/audio-video-media';
import { VideoId } from '@nodelib/services/ddd/admin-videos-catalog/video/entity';

export class VideoMedia extends AudioVideoMedia {
  static max_size = 1024 * 1024 * 1024 * 50; // 50GB
  static mime_types = ['video/mp4'];

  static createFromFile({
    raw_name,
    mime_type,
    size,
    video_id,
  }: {
    raw_name: string;
    mime_type: string;
    size: number;
    video_id: VideoId;
  }) {
    const mediaFileValidator = new MediaFileValidator(
      VideoMedia.max_size,
      VideoMedia.mime_types,
    );

    return Either.safe(() => {
      const { name: newName } = mediaFileValidator.validate({
        raw_name,
        mime_type,
        size,
      });
      return VideoMedia.create({
        name: `${video_id.id}-${newName}`,
        raw_location: `videos/${video_id.id}/videos`,
      });
    });
  }

  static create({ name, raw_location }) {
    return new VideoMedia({
      name,
      raw_location,
      status: AudioVideoMediaStatus.PENDING,
    });
  }

  process() {
    return new VideoMedia({
      name: this.name,
      raw_location: this.raw_location,
      encoded_location: this.encoded_location!,
      status: AudioVideoMediaStatus.PROCESSING,
    });
  }

  complete(encoded_location: string) {
    return new VideoMedia({
      name: this.name,
      raw_location: this.raw_location,
      encoded_location,
      status: AudioVideoMediaStatus.COMPLETED,
    });
  }

  fail() {
    return new VideoMedia({
      name: this.name,
      raw_location: this.raw_location,
      encoded_location: this.encoded_location!,
      status: AudioVideoMediaStatus.FAILED,
    });
  }
}