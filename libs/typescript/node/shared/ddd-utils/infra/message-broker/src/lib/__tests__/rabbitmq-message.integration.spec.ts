import { RabbitMQMessageBroker } from '../rabbitmq-message-broker';
import { IIntegrationEvent } from '@nodelib/shared/ddd-utils/events';
import { Uuid } from '@nodelib/shared/value-objects/uuid';
import { AmqpConnection } from '@golevelup/nestjs-rabbitmq';
import { Config } from '@nodelib/shared/ddd-utils/infra/testing';
import { ConsumeMessage } from 'amqplib';

class TestEvent implements IIntegrationEvent {
  occurred_on: Date = new Date();
  // eslint-disable-next-line @typescript-eslint/no-inferrable-types
  event_version: number = 1;
  event_name: string = TestEvent.name;
  constructor(readonly payload: any) {}
}

describe('RabbitMQMessageBroker Integration tests', () => {
  let service: RabbitMQMessageBroker;
  let connection: AmqpConnection;
  beforeEach(async () => {
    connection = new AmqpConnection({
      uri: Config.rabbitmqUri(),
      connectionInitOptions: { wait: true },
      logger: {
        // eslint-disable-next-line @typescript-eslint/no-empty-function
        debug: () => {},
        // eslint-disable-next-line @typescript-eslint/no-empty-function
        error: () => {},
        // eslint-disable-next-line @typescript-eslint/no-empty-function
        info: () => {},
        // eslint-disable-next-line @typescript-eslint/no-empty-function
        warn: () => {},
        // eslint-disable-next-line @typescript-eslint/no-empty-function
        log: () => {},
      } as any,
    });

    await connection.init();
    const channel = connection.channel;

    await channel.assertExchange('test-exchange', 'direct', {
      durable: false,
    });
    await channel.assertQueue('test-queue', { durable: false });
    await channel.purgeQueue('test-queue');
    await channel.bindQueue('test-queue', 'test-exchange', 'TestEvent');
    service = new RabbitMQMessageBroker(connection);
  });

  afterEach(async () => {
    try {
      await connection.managedConnection.close();
    // eslint-disable-next-line no-empty
    } catch (err) {}
  });

  describe('publish', () => {
    it('should publish events to channel', async () => {
      const event = new TestEvent(new Uuid());

      await service.publishEvent(event);
      const msg: ConsumeMessage = await new Promise((resolve) => {
        connection.channel.consume('test-queue', (msg) => {
          resolve(msg as any);
        });
      });
      const msgObj = JSON.parse(msg.content.toString());
      expect(msgObj).toEqual({
        event_name: TestEvent.name,
        event_version: 1,
        occurred_on: event.occurred_on.toISOString(),
        payload: event.payload,
      });
    });
  });
});