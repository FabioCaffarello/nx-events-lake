import { Either } from '@nodelib/shared/ddd-utils/either';
import {
  InvalidMediaFileMimeTypeError,
  InvalidMediaFileSizeError,
  MediaFileValidator,
} from '@nodelib/shared/validators';
import { ImageMedia } from '@nodelib/shared/value-objects/image-media';
import { VideoId } from '@nodelib/services/ddd/admin-videos-catalog/video/entity';

export class ThumbnailHalf extends ImageMedia {
  static max_size = 1024 * 1024 * 2;
  static mime_types = ['image/jpeg', 'image/png'];

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
      ThumbnailHalf.max_size,
      ThumbnailHalf.mime_types,
    );
    return Either.safe<
      ThumbnailHalf,
      InvalidMediaFileSizeError | InvalidMediaFileMimeTypeError
    >(() => {
      const { name } = mediaFileValidator.validate({
        raw_name,
        mime_type,
        size,
      });
      return new ThumbnailHalf({
        name: `${video_id.id}-${name}`,
        location: `videos/${video_id.id}/images`,
      });
    });
  }
}