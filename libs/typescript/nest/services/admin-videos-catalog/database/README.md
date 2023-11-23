# database


The `database` library is designed to simplify database setup in a NestJS application. It leverages the `@nestjs/sequelize` module and provides a flexible way to configure and connect to different databases, such as SQLite and MySQL. This library is particularly useful for managing database connections and models.

## Usage

### DatabaseModule

The `DatabaseModule` simplifies the configuration and connection to databases. It supports both SQLite and MySQL configurations based on the `DB_VENDOR` environment variable.

#### Example

```typescript
import { Module } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { SequelizeModule } from '@nestjs/sequelize';
import { CONFIG_SCHEMA_TYPE } from '@nestlib/services/admin-videos-catalog/config-setup';

const models = [/* ... Add your Sequelize models here */];

@Module({
  imports: [
    SequelizeModule.forRootAsync({
      useFactory: (configService: ConfigService<CONFIG_SCHEMA_TYPE>) => {
        const dbVendor = configService.get('DB_VENDOR');
        // Configure for SQLite
        if (dbVendor === 'sqlite') {
          return {
            dialect: 'sqlite',
            host: configService.get('DB_HOST'),
            models,
            logging: configService.get('DB_LOGGING'),
            autoLoadModels: configService.get('DB_AUTO_LOAD_MODELS'),
          };
        }
        // Configure for MySQL
        if (dbVendor === 'mysql') {
          return {
            dialect: 'mysql',
            host: configService.get('DB_HOST'),
            port: configService.get('DB_PORT'),
            database: configService.get('DB_DATABASE'),
            username: configService.get('DB_USERNAME'),
            password: configService.get('DB_PASSWORD'),
            models,
            logging: configService.get('DB_LOGGING'),
            autoLoadModels: configService.get('DB_AUTO_LOAD_MODELS'),
          };
        }

        throw new Error(`Unsupported database configuration: ${dbVendor}`);
      },
      inject: [ConfigService],
    }),
  ],
})
export class DatabaseModule {}
```

### MigrationsModule

The `MigrationsModule` integrates the `ConfigModule` and `DatabaseModule` for easy migration setup.

#### Example

```typescript
import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestlib/services/admin-videos-catalog/config-setup';
import { DatabaseModule } from '@nestlib/services/admin-videos-catalog/database';

@Module({
  imports: [ConfigModule.forRoot(), DatabaseModule],
})
export class MigrationsModule {}
```

## Unit Tests

You can use `jest` testing framework for this purpose and you can run it with `nx`:

```sh
npx nx test typescript-nest-services-admin-videos-catalog-database
```

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.