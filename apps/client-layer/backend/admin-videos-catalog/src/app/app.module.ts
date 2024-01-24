import { ConfigModule } from '@nestlib/services/admin-videos-catalog/config-setup';
import { DatabaseModule } from '@nestlib/services/admin-videos-catalog/database';
import { CategoriesModule } from '@nestlib/services/admin-videos-catalog/features/categories';
import { SharedModule } from '@nestlib/shared/module';
import { Module } from '@nestjs/common';
import { CastMembersModule } from '@nestlib/services/admin-videos-catalog/features/cast-member';
import { GenresModule } from '@nestlib/services/admin-videos-catalog/features/genres';
import { VideosModule } from '@nestlib/services/admin-videos-catalog/features/videos';
import { EventModule } from '@nestlib/services/admin-videos-catalog/event';
import { UseCaseModule } from '@nestlib/services/admin-videos-catalog/use-case';
import { RabbitmqFakeController, RabbitMQFakeConsumer } from '@nestlib/services/admin-videos-catalog/rabbitmq-fake';
import { RabbitmqModule } from '@nestlib/services/admin-videos-catalog/rabbitmq';
import { AuthModule } from '@nestlib/services/admin-videos-catalog/auth';

@Module({
  imports: [
    ConfigModule.forRoot(),
    SharedModule,
    DatabaseModule,
    EventModule,
    UseCaseModule,
    RabbitmqModule.forRoot(),
    AuthModule,
    CategoriesModule,
    CastMembersModule,
    GenresModule,
    VideosModule,
  ],
  providers: [RabbitMQFakeConsumer],
  controllers: [RabbitmqFakeController],
})
export class AppModule {}