import {
  IsInt,
  IsNotEmpty,
  IsOptional,
  IsString,
  validateSync,
} from 'class-validator';
import { CastMemberTypes } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/entity';

export type UpdateCastMemberInputConstructorProps = {
  id: string;
  name?: string;
  type?: CastMemberTypes;
};

export class UpdateCastMemberInput {
  @IsString()
  @IsNotEmpty()
  id: string;

  @IsString()
  @IsOptional()
  name: string;

  @IsInt()
  @IsOptional()
  type: CastMemberTypes;

  constructor(props?: UpdateCastMemberInputConstructorProps) {
    if (!props) return;
    this.id = props.id;
    props.name && (this.name = props.name);
    props.type && (this.type = props.type);
  }
}

export class ValidateUpdateCastMemberInput {
  static validate(input: UpdateCastMemberInput) {
    return validateSync(input);
  }
}