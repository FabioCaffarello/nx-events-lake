import { SequelizeOptions } from 'sequelize-typescript';
import { setupSequelize } from '@nodelib/shared/ddd-utils/infra/testing';
import {
  VideoCastMemberModel,
  VideoCategoryModel,
  VideoGenreModel,
  VideoModel,
  AudioVideoMediaModel,
  ImageMediaModel,
} from '@nodelib/services/ddd/admin-videos-catalog/video/infra/db/sequelize';
import { CastMemberModel } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/infra/db/sequelize';
import { CategoryModel } from '@nodelib/services/ddd/admin-videos-catalog/category/infra/db/sequelize';;
import {
  GenreCategoryModel,
  GenreModel,
} from '@nodelib/services/ddd/admin-videos-catalog/genre/infra/db/sequelize';;

export function setupSequelizeForVideo(options: SequelizeOptions = {}) {
  return setupSequelize({
    models: [
      ImageMediaModel,
      VideoModel,
      AudioVideoMediaModel,
      VideoCategoryModel,
      CategoryModel,
      VideoGenreModel,
      GenreModel,
      GenreCategoryModel,
      VideoCastMemberModel,
      CastMemberModel,
    ],
    ...options,
  });
}