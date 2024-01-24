import { OmitType } from '@nestjs/mapped-types';
import { UpdateVideoInput } from '@nodelib/services/ddd/admin-videos-catalog/video/application/use-cases/update-video';

export class UpdateVideoInputWithoutId extends OmitType(UpdateVideoInput, [
  'id',
] as any) {}

export class UpdateVideoDto extends UpdateVideoInputWithoutId {}