
import { FieldsErrors } from '@nodelib/shared/validators';
import { ValueObject } from '@nodelib/shared/value-object';


declare global {
  namespace jest {
    interface Matchers<R> {
      //containsErrorMessages: (expected: FieldsErrors) => R;
      notificationContainsErrorMessages: (expected: Array<FieldsErrors>) => R;
      toBeValueObject: (expected: ValueObject) => R;
    }
  }
}