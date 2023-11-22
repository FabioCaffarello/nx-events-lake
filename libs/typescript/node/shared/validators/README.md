# Validators

The `validators` library provides utility classes and functions for validation, including a class for integrating with class-validator and a set of validation rules.

## Usage

### ClassValidatorFields

The `ClassValidatorFields` class is designed to facilitate integration with class-validator by validating data against specified fields and populating a `Notification` object with validation errors.

Example:

```typescript
import { ClassValidatorFields, Notification } from '@nodelib/shared/validators';

class MyClassValidator extends ClassValidatorFields {
  // Your class-specific properties and methods go here
}

const notification = new Notification();
const myData = { /* your data */ };
const fieldsToValidate = ['field1', 'field2'];

const isValid = new MyClassValidator().validate(notification, myData, fieldsToValidate);

if (!isValid) {
  console.log(notification.toJSON()); // Outputs validation errors
}
```

### Notification

The `Notification` class manages and organizes validation errors. It supports adding, setting, copying errors, and checking for the presence of errors.

Example:

```typescript
import { Notification } from '@nodelib/shared/validators';

const notification = new Notification();
notification.addError('Validation error', 'fieldName');
console.log(notification.hasErrors()); // Outputs: true
console.log(notification.toJSON()); // Outputs: Array of errors
```

### ValidatorRules

The `ValidatorRules` class provides a set of validation rules for common scenarios like required fields, string length, boolean values, etc.

Example:

```typescript
import { ValidatorRules, ValidationError } from '@nodelib/shared/validators';

const myValue = 'example';
const propertyName = 'exampleProperty';

try {
  ValidatorRules.values(myValue, propertyName).required().string().maxLength(10);
} catch (error) {
  if (error instanceof ValidationError) {
    console.error(error.message); // Outputs: Validation error message
  }
}
```

### ClassValidatorFields

#### `validate(notification: Notification, data: any, fields: string[]): boolean`

Validates the provided `data` against the specified `fields` and populates the `notification` with validation errors. Returns `true` if no errors are found.

### Notification

#### `addError(error: string, field?: string): void`

Adds a validation error to the notification, optionally associating it with a specific field.

#### `setError(error: string | string[], field?: string): void`

Sets a validation error for the notification, optionally associating it with a specific field.

#### `hasErrors(): boolean`

Returns `true` if the notification has any validation errors.

#### `copyErrors(notification: Notification): void`

Copies errors from another `Notification` instance.

#### `toJSON(): Array<string | { [key: string]: string[] }>`

Converts the notification errors to a JSON-friendly format.

### ValidatorRules

#### `required(): Omit<this, 'required'>`

Throws a `ValidationError` if the value is null, undefined, or an empty string.

#### `string(): Omit<this, 'string'>`

Throws a `ValidationError` if the value is not a string.

#### `maxLength(max: number): Omit<this, 'maxLength'>`

Throws a `ValidationError` if the string length exceeds the specified maximum.

#### `boolean(): Omit<this, 'boolean'>`

Throws a `ValidationError` if the value is not a boolean.

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
