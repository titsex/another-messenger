// @generated by protoc-gen-es v2.2.5 with parameter "target=ts"
// @generated from file user/v1/user.proto (package user.v1, syntax proto3)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { enumDesc, fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file user/v1/user.proto.
 */
export const file_user_v1_user: GenFile = /*@__PURE__*/
  fileDesc("ChJ1c2VyL3YxL3VzZXIucHJvdG8SB3VzZXIudjEiogMKBFVzZXISCgoCaWQYASABKAkSDAoEbmFtZRgCIAEoCRIQCgh1c2VybmFtZRgDIAEoCRINCgVlbWFpbBgEIAEoCRIVCg1wYXNzd29yZF9oYXNoGAUgASgJEhIKCmNyZWF0ZWRfYXQYBiABKAMSEgoKdXBkYXRlZF9hdBgHIAEoAxIRCglpc19iYW5uZWQYECABKAgSGwoEcm9sZRgJIAEoDjINLnVzZXIudjEuUm9sZRIbCgR0eXBlGAogASgOMg0udXNlci52MS5UeXBlEh8KBnN0YXR1cxgLIAEoDjIPLnVzZXIudjEuU3RhdHVzEhcKCmRlbGV0ZWRfYXQYCCABKANIAIgBARIQCgNiaW8YDCABKAlIAYgBARISCgVwaG9uZRgNIAEoCUgCiAEBEhcKCmF2YXRhcl91cmwYDiABKAlIA4gBARIZCgxsYXN0X3NlZW5fYXQYDyABKANIBIgBAUINCgtfZGVsZXRlZF9hdEIGCgRfYmlvQggKBl9waG9uZUINCgtfYXZhdGFyX3VybEIPCg1fbGFzdF9zZWVuX2F0KjcKBFJvbGUSCAoEVVNFUhAAEgcKA0JPVBABEg0KCU1PREVSQVRPUhACEg0KCURFVkVMT1BFUhADKiwKBFR5cGUSCgoGTk9STUFMEAASCwoHU1BBTU1FUhABEgsKB1NDQU1NRVIQAiotCgZTdGF0dXMSCgoGSElEREVOEAASCgoGT05MSU5FEAESCwoHT0ZGTElORRACQk5aTGdpdGh1Yi5jb20vdGl0c2V4L2Fub3RoZXItbWVzc2VuZ2VyL3NoYXJlZC9wcm90by9nZW5lcmF0ZWQvZ28vdXNlci92MTt1c2VycGJiBnByb3RvMw");

/**
 * @generated from message user.v1.User
 */
export type User = Message<"user.v1.User"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string name = 2;
   */
  name: string;

  /**
   * @generated from field: string username = 3;
   */
  username: string;

  /**
   * @generated from field: string email = 4;
   */
  email: string;

  /**
   * @generated from field: string password_hash = 5;
   */
  passwordHash: string;

  /**
   * @generated from field: int64 created_at = 6;
   */
  createdAt: bigint;

  /**
   * @generated from field: int64 updated_at = 7;
   */
  updatedAt: bigint;

  /**
   * @generated from field: bool is_banned = 16;
   */
  isBanned: boolean;

  /**
   * @generated from field: user.v1.Role role = 9;
   */
  role: Role;

  /**
   * @generated from field: user.v1.Type type = 10;
   */
  type: Type;

  /**
   * @generated from field: user.v1.Status status = 11;
   */
  status: Status;

  /**
   * @generated from field: optional int64 deleted_at = 8;
   */
  deletedAt?: bigint;

  /**
   * @generated from field: optional string bio = 12;
   */
  bio?: string;

  /**
   * @generated from field: optional string phone = 13;
   */
  phone?: string;

  /**
   * @generated from field: optional string avatar_url = 14;
   */
  avatarUrl?: string;

  /**
   * @generated from field: optional int64 last_seen_at = 15;
   */
  lastSeenAt?: bigint;
};

/**
 * Describes the message user.v1.User.
 * Use `create(UserSchema)` to create a new message.
 */
export const UserSchema: GenMessage<User> = /*@__PURE__*/
  messageDesc(file_user_v1_user, 0);

/**
 * @generated from enum user.v1.Role
 */
export enum Role {
  /**
   * @generated from enum value: USER = 0;
   */
  USER = 0,

  /**
   * @generated from enum value: BOT = 1;
   */
  BOT = 1,

  /**
   * @generated from enum value: MODERATOR = 2;
   */
  MODERATOR = 2,

  /**
   * @generated from enum value: DEVELOPER = 3;
   */
  DEVELOPER = 3,
}

/**
 * Describes the enum user.v1.Role.
 */
export const RoleSchema: GenEnum<Role> = /*@__PURE__*/
  enumDesc(file_user_v1_user, 0);

/**
 * @generated from enum user.v1.Type
 */
export enum Type {
  /**
   * @generated from enum value: NORMAL = 0;
   */
  NORMAL = 0,

  /**
   * @generated from enum value: SPAMMER = 1;
   */
  SPAMMER = 1,

  /**
   * @generated from enum value: SCAMMER = 2;
   */
  SCAMMER = 2,
}

/**
 * Describes the enum user.v1.Type.
 */
export const TypeSchema: GenEnum<Type> = /*@__PURE__*/
  enumDesc(file_user_v1_user, 1);

/**
 * @generated from enum user.v1.Status
 */
export enum Status {
  /**
   * @generated from enum value: HIDDEN = 0;
   */
  HIDDEN = 0,

  /**
   * @generated from enum value: ONLINE = 1;
   */
  ONLINE = 1,

  /**
   * @generated from enum value: OFFLINE = 2;
   */
  OFFLINE = 2,
}

/**
 * Describes the enum user.v1.Status.
 */
export const StatusSchema: GenEnum<Status> = /*@__PURE__*/
  enumDesc(file_user_v1_user, 2);

