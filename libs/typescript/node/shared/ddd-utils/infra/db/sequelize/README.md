# sequelize

The `sequelize` library provides a utility for managing database migrations using Sequelize and Umzug. This library is particularly useful for maintaining the schema of your database across different versions of your application.

## Usage

### Migrator

The `migrator` function facilitates the creation of an Umzug instance configured for Sequelize migrations.

#### Example:

```typescript
import { Sequelize } from 'sequelize';
import { migrator } from '@nodelib/shared/ddd-utils/infra/db/sequelize';

// Create a Sequelize instance
const sequelize = new Sequelize('database', 'username', 'password', {
  host: 'localhost',
  dialect: 'postgres',
  // Add other Sequelize configurations as needed
});

// Create a migrator instance
const umzug = migrator(sequelize, {
  // Customize Umzug options if needed
});

// To run migrations
umzug.up().then(() => {
  console.log('Migrations have been executed successfully.');
}).catch((error) => {
  console.error('Error executing migrations:', error);
});
```

## Example

Here is an example of how to use the `sequelize` library in your project:

```typescript
import { Sequelize } from 'sequelize';
import { migrator } from '@nodelib/shared/ddd-utils/infra/db/sequelize';

// Create a Sequelize instance
const sequelize = new Sequelize('database', 'username', 'password', {
  host: 'localhost',
  dialect: 'postgres',
  // Add other Sequelize configurations as needed
});

// Create a migrator instance
const umzug = migrator(sequelize);

// To run migrations
umzug.up().then(() => {
  console.log('Migrations have been executed successfully.');
}).catch((error) => {
  console.error('Error executing migrations:', error);
});
```

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.
