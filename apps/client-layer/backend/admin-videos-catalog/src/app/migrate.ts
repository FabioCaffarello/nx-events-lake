import { migrator } from '@nodelib/shared/ddd-utils/infra/db/sequelize';
import { MigrationsModule } from '@nestlib/services/admin-videos-catalog/database';
import { NestFactory } from '@nestjs/core';
import { getConnectionToken } from '@nestjs/sequelize';

async function bootstrap() {
  const app = await NestFactory.createApplicationContext(MigrationsModule, {
    logger: ['error'],
  });

  const sequelize = app.get(getConnectionToken());

  migrator(sequelize).runAsCLI();
}
bootstrap();
