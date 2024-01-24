import { ICastMemberRepository } from '@nodelib/services/ddd/admin-videos-catalog/cast-member/repository';
import { ICategoryRepository } from '@nodelib/services/ddd/admin-videos-catalog/category/repository';
import { IGenreRepository } from '@nodelib/services/ddd/admin-videos-catalog/genre/repository';
import { IUseCase } from '@nodelib/shared/application';
import { NotFoundError } from '@nodelib/shared/errors';
import { Video, VideoId } from '@nodelib/services/ddd/admin-videos-catalog/video/entity';
import { IVideoRepository } from '@nodelib/services/ddd/admin-videos-catalog/video/repository';
import { VideoOutput, VideoOutputMapper } from '@nodelib/services/ddd/admin-videos-catalog/video/application/use-cases/common';

export class GetVideoUseCase
  implements IUseCase<GetVideoInput, GetVideoOutput>
{
  constructor(
    private videoRepo: IVideoRepository,
    private categoryRepo: ICategoryRepository,
    private genreRepo: IGenreRepository,
    private castMemberRepo: ICastMemberRepository,
  ) {}

  async execute(input: GetVideoInput): Promise<GetVideoOutput> {
    const videoId = new VideoId(input.id);
    const video = await this.videoRepo.findById(videoId);
    if (!video) {
      throw new NotFoundError(input.id, Video);
    }
    const genres = await this.genreRepo.findByIds(
      Array.from(video.genres_id.values()),
    );

    const categories = await this.categoryRepo.findByIds(
      Array.from(video.categories_id.values()).concat(
        genres.flatMap((g) => Array.from(g.categories_id.values())),
      ),
    );

    const castMembers = await this.castMemberRepo.findByIds(
      Array.from(video.cast_members_id.values()),
    );

    return VideoOutputMapper.toOutput({
      video,
      genres,
      cast_members: castMembers,
      allCategoriesOfVideoAndGenre: categories,
    });
  }
}

export type GetVideoInput = {
  id: string;
};

export type GetVideoOutput = VideoOutput;