import { IUnitOfWork } from '@nodelib/shared/ddd-utils/repository';

export class UnitOfWorkFakeInMemory implements IUnitOfWork {
  constructor() {}

  async start(): Promise<void> {
    return;
  }

  async commit(): Promise<void> {
    return;
  }

  async rollback(): Promise<void> {
    return;
  }

  do<T>(workFn: (uow: IUnitOfWork) => Promise<T>): Promise<T> {
    return workFn(this);
  }

  getTransaction() {
    return;
  }
}