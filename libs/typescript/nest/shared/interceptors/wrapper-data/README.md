# wrapper-data

The `wrapper-data` library provides a NestJS interceptor, `WrapperDataInterceptor`, which wraps the response data in a consistent structure. This is particularly useful for creating a standardized API response format.


## Usage

### WrapperDataInterceptor

The `WrapperDataInterceptor` is a NestJS interceptor that wraps the response data in a consistent structure. It checks if the response already contains a "meta" key and, if not, wraps the data in a "data" key.

#### Example

```typescript
import { WrapperDataInterceptor } from '@nestlib/shared/interceptors/wrapper-data';

// In your NestJS application module
app.useGlobalInterceptors(new WrapperDataInterceptor());
```

## Interceptor

### WrapperDataInterceptor

```typescript
import {
  CallHandler,
  ExecutionContext,
  Injectable,
  NestInterceptor,
} from '@nestjs/common';
import { Observable, map } from 'rxjs';

@Injectable()
export class WrapperDataInterceptor implements NestInterceptor {
  intercept(context: ExecutionContext, next: CallHandler): Observable<any> {
    return next
      .handle()
      .pipe(map((body) => (!body || 'meta' in body ? body : { data: body })));
  }
}
```

## Unit Tests

You can use `jest` testing framework for this purpose and you can run it with `nx`:

```sh
npx nx test typescript-nest-shared-interceptors-wrapper-data
```

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.