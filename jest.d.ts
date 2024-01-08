
import { FieldsErrors } from './libs/typescript/node/shared/validators/src/index';
import { ValueObject } from './libs/typescript/node/shared/value-object/src/index';


declare global {
  namespace jest {
    interface Matchers<R> {
      notificationContainsErrorMessages: (expected: Array<FieldsErrors>) => R;
      toBeValueObject: (expected: ValueObject) => R;
    }
  }
}