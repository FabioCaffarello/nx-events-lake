import { OnEvent } from '@nestjs/event-emitter';
import { IIntegrationEventHandler, IMessageBroker } from '@nodelib/shared/application';
import { VideoAudioMediaUploadedIntegrationEvent } from '@nodelib/services/ddd/admin-videos-catalog/video/events';

export class PublishVideoMediaReplacedInQueueHandler
  implements IIntegrationEventHandler
{
  constructor(private messageBroker: IMessageBroker) {
   // console.log(messageBroker);
  }

  @OnEvent(VideoAudioMediaUploadedIntegrationEvent.name)
  async handle(event: VideoAudioMediaUploadedIntegrationEvent): Promise<void> {
    await this.messageBroker.publishEvent(event);
  }
}