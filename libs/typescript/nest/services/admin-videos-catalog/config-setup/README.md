# config-setup

The `config-setup` library provides a module for easy configuration setup in a NestJS application. It uses the `@nestjs/config` module and `joi` for environment variable validation. This library is particularly useful for managing configurations related to database connections, logging, and other application-specific settings.

## Usage

### ConfigModule

The `ConfigModule` extends `NestConfigModule` and provides a flexible way to set up environment variables and validation schemas.

#### Example

```typescript
import { Module } from '@nestjs/common';
import { ConfigModuleOptions } from '@nestjs/config';
import { ConfigModule as NestConfigModule } from '@nestlib/services/admin-videos-catalog/config-setup';
import Joi from 'joi';
import { join } from 'path';

type DB_SCHEMA_TYPE = {
  DB_VENDOR: 'mysql' | 'sqlite';
  DB_HOST: string;
  DB_PORT: number;
  DB_USERNAME: string;
  DB_PASSWORD: string;
  DB_DATABASE: string;
  DB_LOGGING: boolean;
  DB_AUTO_LOAD_MODELS: boolean;
};

export const CONFIG_DB_SCHEMA: Joi.StrictSchemaMap<DB_SCHEMA_TYPE> = {
  // ... Define your DB_SCHEMA_TYPE validation here
};

export type CONFIG_SCHEMA_TYPE = DB_SCHEMA_TYPE;

@Module({})
export class ConfigModule extends NestConfigModule {
  static override forRoot(options: ConfigModuleOptions = {}) {

    const { envFilePath, ...otherOptions } = options;

    return super.forRoot({
      isGlobal: true,
      envFilePath: [
        ...(Array.isArray(envFilePath) ? envFilePath : [envFilePath]),
        join(process.cwd(), 'path/to/env/files', `.env.${process.env['NODE_ENV']}`),
        join(process.cwd(), 'path/to/env/files', `.env`),
      ],
      validationSchema: Joi.object({
        ...CONFIG_DB_SCHEMA,
      }),
      ...otherOptions,
    });
  }
}
```

## Unit Tests

You can use `jest` testing framework for this purpose and you can run it with `nx`:

```sh
npx nx test typescript-nest-services-admin-videos-catalog-config-setup
```

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.