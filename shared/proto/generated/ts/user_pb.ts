// @generated by protoc-gen-es v2.2.5 with parameter "target=ts"
// @generated from file user.proto (package user, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from '@bufbuild/protobuf/codegenv1'
import { fileDesc, messageDesc } from '@bufbuild/protobuf/codegenv1'
import type { Message } from '@bufbuild/protobuf'

/**
 * Describes the file user.proto.
 */
export const file_user: GenFile =
	/*@__PURE__*/
	fileDesc(
		'Cgp1c2VyLnByb3RvEgR1c2VyIi8KBFVzZXISCgoCaWQYASABKAkSDAoEbmFtZRgCIAEoCRINCgVlbWFpbBgDIAEoCUJJWkdnaXRodWIuY29tL3RpdHNleC9hbm90aGVyLW1lc3Nlbmdlci9zaGFyZWQvcHJvdG8vZ2VuZXJhdGVkL2dvL3VzZXI7dXNlcmIGcHJvdG8z'
	)

/**
 * @generated from message user.User
 */
export type User = Message<'user.User'> & {
	/**
	 * @generated from field: string id = 1;
	 */
	id: string

	/**
	 * @generated from field: string name = 2;
	 */
	name: string

	/**
	 * @generated from field: string email = 3;
	 */
	email: string
}

/**
 * Describes the message user.User.
 * Use `create(UserSchema)` to create a new message.
 */
export const UserSchema: GenMessage<User> = /*@__PURE__*/ messageDesc(file_user, 0)
