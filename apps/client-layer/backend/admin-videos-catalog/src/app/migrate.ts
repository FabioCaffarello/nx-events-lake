import { migrator } from '@admin-videos-catalog/core/shared/infra/db/sequelize';
import { MigrationsModule } from '@admin-videos-catalog/nest-modules/database-module';
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
