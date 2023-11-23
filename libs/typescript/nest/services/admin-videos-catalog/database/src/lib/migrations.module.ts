import { ConfigModule } from '@nestlib/services/admin-videos-catalog/config-setup';
import { Module } from '@nestjs/common';
import { DatabaseModule } from './database.module';

@Module({
  imports: [ConfigModule.forRoot(), DatabaseModule],
})
export class MigrationsModule {}
