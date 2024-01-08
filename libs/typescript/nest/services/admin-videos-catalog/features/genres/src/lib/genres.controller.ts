import {
    Controller,
    Get,
    Post,
    Body,
    Patch,
    Param,
    Delete,
    Inject,
    ParseUUIDPipe,
    HttpCode,
    Query,
  } from '@nestjs/common';
  import { CreateGenreDto } from './dto/create-genre.dto';
  import { UpdateGenreDto } from './dto/update-genre.dto';
  import { SearchGenreDto } from './dto/search-genres.dto';
  import { GenreCollectionPresenter, GenrePresenter } from './genres.presenter';
  import { CreateGenreUseCase } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/use-cases/create-genre';
  import { UpdateGenreUseCase } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/use-cases/update-genre';
  import { DeleteGenreUseCase } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/use-cases/delete-genre';
  import { GetGenreUseCase } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/use-cases/get-genre';
  import { ListGenresUseCase } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/use-cases/list-genres';
  import { UpdateGenreInput } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/use-cases/update-genre';
  import { GenreOutput } from '@nodelib/services/ddd/admin-videos-catalog/genre/application/use-cases/common';
  
  @Controller('genres')
  export class GenresController {
    @Inject(CreateGenreUseCase)
    private createUseCase!: CreateGenreUseCase;
  
    @Inject(UpdateGenreUseCase)
    private updateUseCase!: UpdateGenreUseCase;
  
    @Inject(DeleteGenreUseCase)
    private deleteUseCase!: DeleteGenreUseCase;
  
    @Inject(GetGenreUseCase)
    private getUseCase!: GetGenreUseCase;
  
    @Inject(ListGenresUseCase)
    private listUseCase!: ListGenresUseCase;
  
    @Post()
    async create(@Body() createGenreDto: CreateGenreDto) {
      const output = await this.createUseCase.execute(createGenreDto);
      return GenresController.serialize(output);
    }
  
    @Get()
    async search(@Query() searchParams: SearchGenreDto) {
      const output = await this.listUseCase.execute(searchParams);
      return new GenreCollectionPresenter(output);
    }
  
    @Get(':id')
    async findOne(
      @Param('id', new ParseUUIDPipe({ errorHttpStatusCode: 422 })) id: string,
    ) {
      const output = await this.getUseCase.execute({ id });
      return GenresController.serialize(output);
    }
  
    @Patch(':id')
    async update(
      @Param('id', new ParseUUIDPipe({ errorHttpStatusCode: 422 })) id: string,
      @Body() updateGenreDto: UpdateGenreDto,
    ) {
      const input = new UpdateGenreInput({ id, ...updateGenreDto });
      const output = await this.updateUseCase.execute(input);
      return GenresController.serialize(output);
    }
  
    @HttpCode(204)
    @Delete(':id')
    remove(
      @Param('id', new ParseUUIDPipe({ errorHttpStatusCode: 422 })) id: string,
    ) {
      return this.deleteUseCase.execute({ id });
    }
  
    static serialize(output: GenreOutput) {
      return new GenrePresenter(output);
    }
  }
  