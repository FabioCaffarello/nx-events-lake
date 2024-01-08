import { CategoryId } from '@nodelib/services/ddd/admin-videos-catalog/category/entity';
import { Notification, LoadEntityError } from '@nodelib/shared/validators';
import { Genre, GenreId } from '@nodelib/services/ddd/admin-videos-catalog/genre/entity';
import { GenreCategoryModel, GenreModel } from './genre-model';

export class GenreModelMapper {
  static toEntity(model: GenreModel) {
    const { genre_id: id, categories_id = [], ...otherData } = model.toJSON();
    const categoriesId = categories_id.map(
      (c) => new CategoryId(c.category_id),
    );

    const notification = new Notification();
    if (!categoriesId.length) {
      notification.addError(
        'categories_id should not be empty',
        'categories_id',
      );
    }

    const genre = new Genre({
      ...otherData,
      genre_id: new GenreId(id),
      categories_id: new Map(categoriesId.map((c) => [c.id, c])),
    });

    genre.validate();

    notification.copyErrors(genre.notification);

    if (notification.hasErrors()) {
      throw new LoadEntityError(notification.toJSON());
    }

    return genre;
  }

  static toModelProps(aggregate: Genre) {
    const { categories_id, ...otherData } = aggregate.toJSON();
    return {
      ...otherData,
      categories_id: categories_id.map(
        (category_id) =>
          new GenreCategoryModel({
            genre_id: aggregate.genre_id.id,
            category_id: category_id,
          }),
      ),
    };
  }
}