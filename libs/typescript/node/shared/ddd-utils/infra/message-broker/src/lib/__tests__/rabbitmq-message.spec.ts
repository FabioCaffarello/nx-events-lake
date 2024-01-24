import { ChannelWrapper } from 'amqp-connection-manager';
import { RabbitMQMessageBroker } from '../rabbitmq-message-broker';
import { IIntegrationEvent } from '@nodelib/shared/ddd-utils/events';
import { Uuid } from '@nodelib/shared/value-objects/uuid';
import { EVENTS_MESSAGE_BROKER_CONFIG } from '../events-message-broker-config';

class TestEvent implements IIntegrationEvent {
  occurred_on: Date = new Date();
  // eslint-disable-next-line @typescript-eslint/no-inferrable-types
  event_version: number = 1;
  event_name: string = TestEvent.name;
  constructor(readonly payload: any) {}
}

describe('RabbitMQMessageBroker Unit tests', () => {
  let service: RabbitMQMessageBroker;
  let connection: ChannelWrapper;
  beforeEach(async () => {
    connection = {
      publish: jest.fn(),
    } as any;
    service = new RabbitMQMessageBroker(connection as any);
  });

  describe('publish', () => {
    it('should publish events to channel', async () => {
      const event = new TestEvent(new Uuid());

      await service.publishEvent(event);

      expect(connection.publish).toBeCalledWith(
        EVENTS_MESSAGE_BROKER_CONFIG[TestEvent.name].exchange,
        EVENTS_MESSAGE_BROKER_CONFIG[TestEvent.name].routing_key,
        event,
      );
    });
  });
});