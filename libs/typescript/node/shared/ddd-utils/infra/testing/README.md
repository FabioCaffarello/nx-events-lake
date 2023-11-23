# testing

The `testing` library provides utility functions and configurations for testing applications. It includes functionality for handling environment configurations, Sequelize setup for database testing, and custom Jest matchers for assertions.


## Configuration

### Environment Configuration

The `Config` class provides methods to read environment variables, specifically tailored for testing. The `readEnv` method reads environment variables from a specified file path, allowing you to customize the environment for different testing scenarios.

#### Example:

```typescript
import { Config } from '@nodelib/shared/ddd-utils/infra/testing';

// Read environment variables for the 'admin-videos-catalog' service
Config.readEnv('admin-videos-catalog');

// Access the configured environment variables
console.log(Config.env.DB_HOST);
```

### Sequelize Setup

The `setupSequelize` function simplifies the setup and teardown of a Sequelize instance for testing purposes. It takes Sequelize options as parameters and returns an object containing the Sequelize instance.

#### Example:

```typescript
import { setupSequelize } from '@nodelib/shared/ddd-utils/infra/testing';

// Setup Sequelize for testing
const { sequelize } = setupSequelize();

// Use the sequelize instance for testing database operations
```

### Custom Jest Matchers

The library provides a custom Jest matcher for assertions related to error messages in a `Notification` object. The `notificationContainsErrorMessages` matcher can be used to assert that a `Notification` instance contains specific error messages.

#### Example:

```typescript
import { expect } from '@jest/globals';
import { Notification } from '@nodelib/shared/validators';

const notification = new Notification();
notification.addError('field1', 'Error message for field1');

// Assert that the Notification contains specific error messages
expect(notification).notificationContainsErrorMessages([
  'Error message for field1',
]);

// Assert that the Notification does not contain specific error messages
expect(notification).not.notificationContainsErrorMessages([
  'Error message for field2',
]);
```

## Example

Here is an example of how to use the `testing` library in your project:

```typescript
import { Config, setupSequelize } from '@nodelib/shared/ddd-utils/infra/testing';

// Read environment variables for the 'admin-videos-catalog' service
Config.readEnv('admin-videos-catalog');

// Setup Sequelize for testing
const { sequelize } = setupSequelize();

// Use the sequelize instance for testing database operations

// ... Your test code here ...

// Close the Sequelize connection after testing
await sequelize.close();
```


## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
