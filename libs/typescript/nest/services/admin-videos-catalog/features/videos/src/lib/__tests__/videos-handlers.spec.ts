import { Test, TestingModule } from '@nestjs/testing';
import { DatabaseModule } from '@nestlib/services/admin-videos-catalog/database';
import { VideosModule } from '../videos.module';
import { EventEmitter2 } from '@nestjs/event-emitter';
import { ConfigModule } from '@nestlib/services/admin-videos-catalog/config-setup';
import { SharedModule } from '@nestlib/shared/module';
import { EventModule } from '@nestlib/services/admin-videos-catalog/event';
import { VideoAudioMediaUploadedIntegrationEvent } from '@nodelib/services/ddd/admin-videos-catalog/video/events';
import { UnitOfWorkFakeInMemory } from '@nodelib/shared/ddd-utils/infra/db/in-memory';
import { AmqpConnection } from '@golevelup/nestjs-rabbitmq';
import { UseCaseModule } from '@nestlib/services/admin-videos-catalog/use-case';
import { DynamicModule } from '@nestjs/common';
import { AuthModule } from '@nestlib/services/admin-videos-catalog/auth';

class RabbitmqModuleFake {
  static forRoot(): DynamicModule {
    return {
      module: RabbitmqModuleFake,
      global: true,
      providers: [
        {
          provide: AmqpConnection,
          useValue: {
            publish: jest.fn(),
          },
        },
      ],
      exports: [AmqpConnection],
    };
  }
}

describe('VideosModule Unit Tests', () => {
  let module: TestingModule;
  beforeEach(async () => {
    module = await Test.createTestingModule({
      imports: [
        ConfigModule.forRoot(),
        SharedModule,
        EventModule,
        UseCaseModule,
        DatabaseModule,
        AuthModule,
        RabbitmqModuleFake.forRoot(),
        VideosModule,
      ],
    })
      .overrideProvider('UnitOfWork')
      .useFactory({
        factory: () => {
          return new UnitOfWorkFakeInMemory();
        },
      })
      .compile();
    await module.init();
  });

  afterEach(async () => {
    await module.close();
  });

  it('should register handlers', async () => {
    const eventemitter2 = module.get<EventEmitter2>(EventEmitter2);
    expect(
      eventemitter2.listeners(VideoAudioMediaUploadedIntegrationEvent.name),
    ).toHaveLength(1);
  });
});