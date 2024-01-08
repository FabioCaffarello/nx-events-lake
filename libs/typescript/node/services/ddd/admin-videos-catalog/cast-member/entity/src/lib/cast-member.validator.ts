import { MaxLength } from 'class-validator';
import { ClassValidatorFields, Notification } from '@nodelib/shared/validators';
import { CastMember } from './cast-member.aggregate';

export class CastMemberRules {
  @MaxLength(255, { groups: ['name'] })
  name: string;

  constructor(entity: CastMember) {
    Object.assign(this, entity);
  }
}

export class CastMemberValidator extends ClassValidatorFields {
  validate(
    notification: Notification,
    data: CastMember,
    fields?: string[],
  ): boolean {
    const newFields = fields?.length ? fields : ['name'];
    return super.validate(notification, new CastMemberRules(data), newFields);
  }
}

export class CastMemberValidatorFactory {
  static create() {
    return new CastMemberValidator();
  }
}

export default CastMemberValidatorFactory;