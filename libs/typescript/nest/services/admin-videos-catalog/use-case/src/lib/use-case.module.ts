import { Global, Module, Scope } from '@nestjs/common';
import { ApplicationService } from '@nodelib/shared/application';
import { IUnitOfWork } from '@nodelib/shared/ddd-utils/repository';
import { DomainEventMediator } from '@nodelib/shared/ddd-utils/events';

@Global()
@Module({
  providers: [
    {
      provide: ApplicationService,
      useFactory: (
        uow: IUnitOfWork,
        domainEventMediator: DomainEventMediator,
      ) => {
        return new ApplicationService(uow, domainEventMediator);
      },
      inject: ['UnitOfWork', DomainEventMediator],
      //scope: Scope.REQUEST,
    },
  ],
  exports: [ApplicationService],
})
export class UseCaseModule {}