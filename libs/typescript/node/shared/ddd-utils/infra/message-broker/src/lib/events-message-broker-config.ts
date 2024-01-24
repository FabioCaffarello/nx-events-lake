import { VideoAudioMediaUploadedIntegrationEvent } from '@nodelib/services/ddd/admin-videos-catalog/video/events';

VideoAudioMediaUploadedIntegrationEvent;
export const EVENTS_MESSAGE_BROKER_CONFIG = {
  [VideoAudioMediaUploadedIntegrationEvent.name]: {
    exchange: 'amq.direct',
    routing_key: VideoAudioMediaUploadedIntegrationEvent.name,
  },
  TestEvent: {
    exchange: 'test-exchange',
    routing_key: 'TestEvent',
  },
};