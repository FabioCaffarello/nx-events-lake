# filters

The `filters` library provides custom exception filters for handling specific exceptions in a NestJS application. These filters can be used to provide consistent error responses for different types of exceptions.


## Usage

### NotFoundErrorFilter

The `NotFoundErrorFilter` is an exception filter for handling `NotFoundError` exceptions. It returns a JSON response with a 404 status code.

#### Example

```typescript
import { NotFoundErrorFilter } from '@nestlib/shared/filters';

// In your NestJS application module
app.useGlobalFilters(new NotFoundErrorFilter());
```

### EntityValidationErrorFilter

The `EntityValidationErrorFilter` is an exception filter for handling `EntityValidationError` exceptions. It returns a JSON response with a 422 status code, including detailed error messages.

#### Example

```typescript
import { EntityValidationErrorFilter } from '@nestlib/shared/filters';

// In your NestJS application module
app.useGlobalFilters(new EntityValidationErrorFilter());
```

## Filters

### NotFoundErrorFilter

```typescript
import { NotFoundError } from '@nodelib/shared/errors';
import { ArgumentsHost, Catch, ExceptionFilter } from '@nestjs/common';
import { Response } from 'express';

@Catch(NotFoundError)
export class NotFoundErrorFilter implements ExceptionFilter {
  catch(exception: NotFoundError, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response: Response = ctx.getResponse();

    response.status(404).json({
      statusCode: 404,
      error: 'Not Found',
      message: exception.message,
    });
  }
}
```

### EntityValidationErrorFilter

```typescript
import { EntityValidationError } from '@nodelib/shared/validators';
import { ArgumentsHost, Catch, ExceptionFilter } from '@nestjs/common';
import { Response } from 'express';
import { union } from 'lodash';

@Catch(EntityValidationError)
export class EntityValidationErrorFilter implements ExceptionFilter {
  catch(exception: EntityValidationError, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();
    response.status(422).json({
      statusCode: 422,
      error: 'Unprocessable Entity',
      message: union(
        ...exception.error.reduce(
          (acc, error) =>
            acc.concat(
              typeof error === 'string' ? [[error]] : Object.values(error),
            ),
          [],
        ),
      ),
    });
  }
}
```

## Unit Tests

You can use `jest` testing framework for this purpose and you can run it with `nx`:

```sh
npx nx test typescript-nest-shared-filters
```

## Contributing

Contributions to this library are welcome. If you find issues, have suggestions for improvements, or want to add new features, feel free to create an issue or submit a pull request. Your contributions will help improve the library for all users.