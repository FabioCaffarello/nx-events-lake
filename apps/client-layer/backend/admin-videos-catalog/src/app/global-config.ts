import { EntityValidationErrorFilter, NotFoundErrorFilter } from '@nestlib/shared/filters';
import { WrapperDataInterceptor } from '@nestlib/shared/interceptors/wrapper-data';
import {
  ClassSerializerInterceptor,
  INestApplication,
  ValidationPipe,
} from '@nestjs/common';
import { Reflector } from '@nestjs/core';


export function applyGlobalConfig(app: INestApplication) {
  app.useGlobalPipes(
    new ValidationPipe({
      errorHttpStatusCode: 422,
      transform: true,
    }),
  );
  app.useGlobalInterceptors(
    new WrapperDataInterceptor(),
    new ClassSerializerInterceptor(app.get(Reflector)),
  );
  app.useGlobalFilters(
    new EntityValidationErrorFilter(),
    new NotFoundErrorFilter(),
  );
}
