import { MaxLength } from 'class-validator';
import { 
    ClassValidatorFields,
    Notification,
} from '@nodelib/shared/validators';
import { Genre } from './genre.aggregate';

export class GenreRules {
  @MaxLength(255, { groups: ['name'] })
  name: string;

  constructor(entity: Genre) {
    Object.assign(this, entity);
  }
}

export class GenreValidator extends ClassValidatorFields {
  validate(
    notification: Notification,
    data: Genre,
    fields?: string[],
  ): boolean {
    const newFields = fields?.length ? fields : ['name'];
    return super.validate(notification, new GenreRules(data), newFields);
  }
}

export class GenreValidatorFactory {
  static create() {
    return new GenreValidator();
  }
}

export default GenreValidatorFactory;