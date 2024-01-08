import { IsInt, IsNotEmpty, IsString, validateSync } from 'class-validator';
import { CastMemberTypes } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/entity';

export type CreateCastMemberInputConstructorProps = {
  name: string;
  type: CastMemberTypes;
};

export class CreateCastMemberInput {
  @IsString()
  @IsNotEmpty()
  name: string;

  @IsInt()
  @IsNotEmpty()
  type: CastMemberTypes;

  constructor(props?: CreateCastMemberInputConstructorProps) {
    if (!props) return;
    this.name = props.name;
    this.type = props.type;
  }
}

export class ValidateCreateCastMemberInput {
  static validate(input: CreateCastMemberInput) {
    return validateSync(input);
  }
}