import { Transform } from 'class-transformer';
import { ListCastMembersOutput } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/use-cases/list-cast-members';
import { CastMemberTypes } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/entity';
import { CollectionPresenter } from '@nestlib/shared/presenters';
import { CastMemberOutput } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/application/use-cases/common';

export class CastMemberPresenter {
  id: string;
  name: string;
  type: CastMemberTypes;
  @Transform(({ value }: { value: Date }) => {
    return value.toISOString();
  })
  created_at: Date;

  constructor(output: CastMemberOutput) {
    this.id = output.id;
    this.name = output.name;
    this.type = output.type;
    this.created_at = output.created_at;
  }
}

export class CastMemberCollectionPresenter extends CollectionPresenter {
  data: CastMemberPresenter[];

  constructor(output: ListCastMembersOutput) {
    const { items, ...paginationProps } = output;
    super(paginationProps);
    this.data = items.map((item) => new CastMemberPresenter(item));
  }
}