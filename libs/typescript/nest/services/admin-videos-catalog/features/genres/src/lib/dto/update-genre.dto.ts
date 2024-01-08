import { OmitType } from '@nestjs/mapped-types';
import { UpdateGenreInput } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/use-cases/update-genre';

export class UpdateGenreInputWithoutId extends OmitType(UpdateGenreInput, [
  'id',
] as any) {}

export class UpdateGenreDto extends UpdateGenreInputWithoutId {}