import { Module } from '@nestjs/common';
import { GenresController } from './genres.controller';
import { SequelizeModule } from '@nestjs/sequelize';
import { CategoriesModule } from '@nestlib/services/admin-videos-catalog/features/categories';
import { GENRES_PROVIDERS } from './genres.providers';
import {
  GenreCategoryModel,
  GenreModel,
} from '@nodelib/services/ddd/admin-videos-catalog/genre/infra/db/sequelize';

@Module({
  imports: [
    SequelizeModule.forFeature([GenreModel, GenreCategoryModel]),
    CategoriesModule,
  ],
  controllers: [GenresController],
  providers: [
    ...Object.values(GENRES_PROVIDERS.REPOSITORIES),
    ...Object.values(GENRES_PROVIDERS.USE_CASES),
  ],
  exports: [GENRES_PROVIDERS.REPOSITORIES.GENRE_REPOSITORY.provide],
})
export class GenresModule {}