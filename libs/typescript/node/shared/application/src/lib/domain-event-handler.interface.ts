import {
    IDomainEvent,
    IIntegrationEvent,
  } from '@nodelib/shared/ddd-utils/events';
  
  export interface IDomainEventHandler {
    handle(event: IDomainEvent): Promise<void>;
  }
  
  export interface IIntegrationEventHandler {
    handle(event: IIntegrationEvent): Promise<void>;
  }