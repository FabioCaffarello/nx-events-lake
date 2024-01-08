import { SearchInput } from '@nodelib/shared/application';
import { SortDirection } from '@nodelib/shared/ddd-utils/repository';
import { CastMemberTypes } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/entity';
import { IsInt, ValidateNested, validateSync } from 'class-validator';

export class ListCastMembersFilter {
  name?: string | null;
  @IsInt()
  type?: CastMemberTypes | null;
}

export class ListCastMembersInput
  implements SearchInput<ListCastMembersFilter>
{
  page?: number;
  per_page?: number;
  sort?: string;
  sort_dir?: SortDirection;
  @ValidateNested()
  filter?: ListCastMembersFilter;
}

export class ValidateListCastMembersInput {
  static validate(input: ListCastMembersInput) {
    return validateSync(input);
  }
}