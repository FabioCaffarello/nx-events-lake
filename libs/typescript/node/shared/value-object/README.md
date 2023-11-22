# value-object

The `value-object` library provides a base class for creating value objects in TypeScript. Value objects are immutable objects that represent a descriptive aspect of the system. This library aims to simplify the implementation of value objects by providing a common base class with equality checking.


## Usage

To create a value object, extend the `ValueObject` class provided by this library:

```typescript
import isEqual from "lodash/isEqual";
import { ValueObject } from "@nodelib/shared/value-object";

class MyValueObject extends ValueObject {
  // Your value object properties and methods go here
}

// Example usage
const obj1 = new MyValueObject(/* pass values for initialization */);
const obj2 = new MyValueObject(/* pass values for initialization */);

console.log(obj1.equals(obj2)); // Outputs: true or false
```

The `equals` method is provided by the `ValueObject` class for comparing instances of your value objects. It uses deep equality checking provided by the lodash library.

### Note

- Ensure that your value object class has appropriate properties and methods based on your application requirements.
- Equality checking is based on the lodash `isEqual` function, so make sure your class properties are suitable for deep equality comparison.

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
