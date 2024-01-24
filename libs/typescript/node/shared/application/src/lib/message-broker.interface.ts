import { IIntegrationEvent } from '@nodelib/shared/ddd-utils/events';

export interface IMessageBroker {
  publishEvent(event: IIntegrationEvent): Promise<void>;
}