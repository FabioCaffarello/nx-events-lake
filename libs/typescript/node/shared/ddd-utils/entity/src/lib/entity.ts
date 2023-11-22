import { Notification } from "@nodelib/shared/validators";
import { ValueObject } from '@nodelib/shared/value-object';

export abstract class Entity {

  notification: Notification = new Notification();

  abstract get entity_id(): ValueObject;
  abstract toJSON(): any;
}
