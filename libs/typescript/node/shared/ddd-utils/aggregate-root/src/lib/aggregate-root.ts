import { Entity } from '@nodelib/shared/ddd-utils/entity';
// import { IDomainEvent } from '@nodelib/shared/ddd-utils/events';
// import EventEmitter2 from 'eventemitter2';

// export abstract class AggregateRoot extends Entity {
//   events: Set<IDomainEvent> = new Set<IDomainEvent>();
//   localMediator = new EventEmitter2();
//   applyEvent(event: IDomainEvent) {
//     this.events.add(event);
//     this.localMediator.emit(event.constructor.name, event);
//   }

//   registerHandler(event: string, handler: (event: IDomainEvent) => void) {
//     this.localMediator.on(event, handler);
//   }
// }


export abstract class AggregateRoot extends Entity {}