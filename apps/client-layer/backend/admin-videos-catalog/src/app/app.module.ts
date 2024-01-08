import { ConfigModule } from '@nestlib/services/admin-videos-catalog/config-setup';
import { DatabaseModule } from '@nestlib/services/admin-videos-catalog/database';
import { CategoriesModule } from '@nestlib/services/admin-videos-catalog/features/categories';
import { SharedModule } from '@nestlib/shared/module';
import { Module } from '@nestjs/common';

@Module({
  imports: [
    ConfigModule.forRoot(),
    DatabaseModule,
    CategoriesModule, 
    SharedModule
  ],
})
export class AppModule {}
