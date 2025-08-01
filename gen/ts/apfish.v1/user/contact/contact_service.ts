// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v1.181.2
//   protoc               v4.25.1
// source: apfish.v1/user/contact/contact_service.proto

/* eslint-disable */
import {
  type CallOptions,
  ChannelCredentials,
  Client,
  type ClientOptions,
  type ClientUnaryCall,
  type handleUnaryCall,
  makeGenericClientConstructor,
  Metadata,
  type ServiceError,
  type UntypedServiceImplementation,
} from "@grpc/grpc-js";
import Long from "long";
import _m0 from "protobufjs/minimal";
import { SuccessResponse } from "../../helper/helper";
import { Contact, ContactPatch } from "./contact";
import { ContactSummary } from "./summary/contact_summary";

export const protobufPackage = "apfish.v1.user.contact";

export interface ContactRequest {
  id: string;
}

export interface ContactSummaryResponse {
  contact: ContactSummary | undefined;
}

export interface ListUserContactsRequest {
  userId: string;
}

export interface ListContactsResponse {
  listContacts: ContactSummary[];
}

export interface UpdateContactRequest {
  contact: ContactPatch | undefined;
}

export interface CreateContactRequest {
  userId: string;
  typeId: string;
  value: string;
}

export interface CreateContactResponse {
  contactId: string;
}

export interface DeleteContactRequest {
  contactId: string;
}

function createBaseContactRequest(): ContactRequest {
  return { id: "" };
}

export const ContactRequest = {
  encode(message: ContactRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ContactRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseContactRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ContactRequest {
    return { id: isSet(object.id) ? globalThis.String(object.id) : "" };
  },

  toJSON(message: ContactRequest): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ContactRequest>, I>>(base?: I): ContactRequest {
    return ContactRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ContactRequest>, I>>(object: I): ContactRequest {
    const message = createBaseContactRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseContactSummaryResponse(): ContactSummaryResponse {
  return { contact: undefined };
}

export const ContactSummaryResponse = {
  encode(message: ContactSummaryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.contact !== undefined) {
      ContactSummary.encode(message.contact, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ContactSummaryResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseContactSummaryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.contact = ContactSummary.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ContactSummaryResponse {
    return { contact: isSet(object.contact) ? ContactSummary.fromJSON(object.contact) : undefined };
  },

  toJSON(message: ContactSummaryResponse): unknown {
    const obj: any = {};
    if (message.contact !== undefined) {
      obj.contact = ContactSummary.toJSON(message.contact);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ContactSummaryResponse>, I>>(base?: I): ContactSummaryResponse {
    return ContactSummaryResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ContactSummaryResponse>, I>>(object: I): ContactSummaryResponse {
    const message = createBaseContactSummaryResponse();
    message.contact = (object.contact !== undefined && object.contact !== null)
      ? ContactSummary.fromPartial(object.contact)
      : undefined;
    return message;
  },
};

function createBaseListUserContactsRequest(): ListUserContactsRequest {
  return { userId: "" };
}

export const ListUserContactsRequest = {
  encode(message: ListUserContactsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListUserContactsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListUserContactsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.userId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListUserContactsRequest {
    return { userId: isSet(object.userId) ? globalThis.String(object.userId) : "" };
  },

  toJSON(message: ListUserContactsRequest): unknown {
    const obj: any = {};
    if (message.userId !== "") {
      obj.userId = message.userId;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListUserContactsRequest>, I>>(base?: I): ListUserContactsRequest {
    return ListUserContactsRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ListUserContactsRequest>, I>>(object: I): ListUserContactsRequest {
    const message = createBaseListUserContactsRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseListContactsResponse(): ListContactsResponse {
  return { listContacts: [] };
}

export const ListContactsResponse = {
  encode(message: ListContactsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.listContacts) {
      ContactSummary.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListContactsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListContactsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.listContacts.push(ContactSummary.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListContactsResponse {
    return {
      listContacts: globalThis.Array.isArray(object?.listContacts)
        ? object.listContacts.map((e: any) => ContactSummary.fromJSON(e))
        : [],
    };
  },

  toJSON(message: ListContactsResponse): unknown {
    const obj: any = {};
    if (message.listContacts?.length) {
      obj.listContacts = message.listContacts.map((e) => ContactSummary.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListContactsResponse>, I>>(base?: I): ListContactsResponse {
    return ListContactsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ListContactsResponse>, I>>(object: I): ListContactsResponse {
    const message = createBaseListContactsResponse();
    message.listContacts = object.listContacts?.map((e) => ContactSummary.fromPartial(e)) || [];
    return message;
  },
};

function createBaseUpdateContactRequest(): UpdateContactRequest {
  return { contact: undefined };
}

export const UpdateContactRequest = {
  encode(message: UpdateContactRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.contact !== undefined) {
      ContactPatch.encode(message.contact, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateContactRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateContactRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.contact = ContactPatch.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateContactRequest {
    return { contact: isSet(object.contact) ? ContactPatch.fromJSON(object.contact) : undefined };
  },

  toJSON(message: UpdateContactRequest): unknown {
    const obj: any = {};
    if (message.contact !== undefined) {
      obj.contact = ContactPatch.toJSON(message.contact);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateContactRequest>, I>>(base?: I): UpdateContactRequest {
    return UpdateContactRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<UpdateContactRequest>, I>>(object: I): UpdateContactRequest {
    const message = createBaseUpdateContactRequest();
    message.contact = (object.contact !== undefined && object.contact !== null)
      ? ContactPatch.fromPartial(object.contact)
      : undefined;
    return message;
  },
};

function createBaseCreateContactRequest(): CreateContactRequest {
  return { userId: "", typeId: "", value: "" };
}

export const CreateContactRequest = {
  encode(message: CreateContactRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    if (message.typeId !== "") {
      writer.uint32(18).string(message.typeId);
    }
    if (message.value !== "") {
      writer.uint32(26).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateContactRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateContactRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.userId = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.typeId = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.value = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateContactRequest {
    return {
      userId: isSet(object.userId) ? globalThis.String(object.userId) : "",
      typeId: isSet(object.typeId) ? globalThis.String(object.typeId) : "",
      value: isSet(object.value) ? globalThis.String(object.value) : "",
    };
  },

  toJSON(message: CreateContactRequest): unknown {
    const obj: any = {};
    if (message.userId !== "") {
      obj.userId = message.userId;
    }
    if (message.typeId !== "") {
      obj.typeId = message.typeId;
    }
    if (message.value !== "") {
      obj.value = message.value;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateContactRequest>, I>>(base?: I): CreateContactRequest {
    return CreateContactRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateContactRequest>, I>>(object: I): CreateContactRequest {
    const message = createBaseCreateContactRequest();
    message.userId = object.userId ?? "";
    message.typeId = object.typeId ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseCreateContactResponse(): CreateContactResponse {
  return { contactId: "" };
}

export const CreateContactResponse = {
  encode(message: CreateContactResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.contactId !== "") {
      writer.uint32(10).string(message.contactId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateContactResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateContactResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.contactId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateContactResponse {
    return { contactId: isSet(object.contactId) ? globalThis.String(object.contactId) : "" };
  },

  toJSON(message: CreateContactResponse): unknown {
    const obj: any = {};
    if (message.contactId !== "") {
      obj.contactId = message.contactId;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateContactResponse>, I>>(base?: I): CreateContactResponse {
    return CreateContactResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateContactResponse>, I>>(object: I): CreateContactResponse {
    const message = createBaseCreateContactResponse();
    message.contactId = object.contactId ?? "";
    return message;
  },
};

function createBaseDeleteContactRequest(): DeleteContactRequest {
  return { contactId: "" };
}

export const DeleteContactRequest = {
  encode(message: DeleteContactRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.contactId !== "") {
      writer.uint32(10).string(message.contactId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteContactRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteContactRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.contactId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteContactRequest {
    return { contactId: isSet(object.contactId) ? globalThis.String(object.contactId) : "" };
  },

  toJSON(message: DeleteContactRequest): unknown {
    const obj: any = {};
    if (message.contactId !== "") {
      obj.contactId = message.contactId;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteContactRequest>, I>>(base?: I): DeleteContactRequest {
    return DeleteContactRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DeleteContactRequest>, I>>(object: I): DeleteContactRequest {
    const message = createBaseDeleteContactRequest();
    message.contactId = object.contactId ?? "";
    return message;
  },
};

export type ContactServiceService = typeof ContactServiceService;
export const ContactServiceService = {
  getContact: {
    path: "/apfish.v1.user.contact.ContactService/GetContact",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: ContactRequest) => Buffer.from(ContactRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => ContactRequest.decode(value),
    responseSerialize: (value: Contact) => Buffer.from(Contact.encode(value).finish()),
    responseDeserialize: (value: Buffer) => Contact.decode(value),
  },
  getContactSummary: {
    path: "/apfish.v1.user.contact.ContactService/GetContactSummary",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: ContactRequest) => Buffer.from(ContactRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => ContactRequest.decode(value),
    responseSerialize: (value: ContactSummaryResponse) => Buffer.from(ContactSummaryResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => ContactSummaryResponse.decode(value),
  },
  listUserContacts: {
    path: "/apfish.v1.user.contact.ContactService/ListUserContacts",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: ListUserContactsRequest) => Buffer.from(ListUserContactsRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => ListUserContactsRequest.decode(value),
    responseSerialize: (value: ListContactsResponse) => Buffer.from(ListContactsResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => ListContactsResponse.decode(value),
  },
  updateContact: {
    path: "/apfish.v1.user.contact.ContactService/UpdateContact",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: UpdateContactRequest) => Buffer.from(UpdateContactRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => UpdateContactRequest.decode(value),
    responseSerialize: (value: SuccessResponse) => Buffer.from(SuccessResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => SuccessResponse.decode(value),
  },
  createContact: {
    path: "/apfish.v1.user.contact.ContactService/CreateContact",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: CreateContactRequest) => Buffer.from(CreateContactRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => CreateContactRequest.decode(value),
    responseSerialize: (value: CreateContactResponse) => Buffer.from(CreateContactResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => CreateContactResponse.decode(value),
  },
  deleteContact: {
    path: "/apfish.v1.user.contact.ContactService/DeleteContact",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: DeleteContactRequest) => Buffer.from(DeleteContactRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => DeleteContactRequest.decode(value),
    responseSerialize: (value: SuccessResponse) => Buffer.from(SuccessResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => SuccessResponse.decode(value),
  },
} as const;

export interface ContactServiceServer extends UntypedServiceImplementation {
  getContact: handleUnaryCall<ContactRequest, Contact>;
  getContactSummary: handleUnaryCall<ContactRequest, ContactSummaryResponse>;
  listUserContacts: handleUnaryCall<ListUserContactsRequest, ListContactsResponse>;
  updateContact: handleUnaryCall<UpdateContactRequest, SuccessResponse>;
  createContact: handleUnaryCall<CreateContactRequest, CreateContactResponse>;
  deleteContact: handleUnaryCall<DeleteContactRequest, SuccessResponse>;
}

export interface ContactServiceClient extends Client {
  getContact(
    request: ContactRequest,
    callback: (error: ServiceError | null, response: Contact) => void,
  ): ClientUnaryCall;
  getContact(
    request: ContactRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: Contact) => void,
  ): ClientUnaryCall;
  getContact(
    request: ContactRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: Contact) => void,
  ): ClientUnaryCall;
  getContactSummary(
    request: ContactRequest,
    callback: (error: ServiceError | null, response: ContactSummaryResponse) => void,
  ): ClientUnaryCall;
  getContactSummary(
    request: ContactRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: ContactSummaryResponse) => void,
  ): ClientUnaryCall;
  getContactSummary(
    request: ContactRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: ContactSummaryResponse) => void,
  ): ClientUnaryCall;
  listUserContacts(
    request: ListUserContactsRequest,
    callback: (error: ServiceError | null, response: ListContactsResponse) => void,
  ): ClientUnaryCall;
  listUserContacts(
    request: ListUserContactsRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: ListContactsResponse) => void,
  ): ClientUnaryCall;
  listUserContacts(
    request: ListUserContactsRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: ListContactsResponse) => void,
  ): ClientUnaryCall;
  updateContact(
    request: UpdateContactRequest,
    callback: (error: ServiceError | null, response: SuccessResponse) => void,
  ): ClientUnaryCall;
  updateContact(
    request: UpdateContactRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: SuccessResponse) => void,
  ): ClientUnaryCall;
  updateContact(
    request: UpdateContactRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: SuccessResponse) => void,
  ): ClientUnaryCall;
  createContact(
    request: CreateContactRequest,
    callback: (error: ServiceError | null, response: CreateContactResponse) => void,
  ): ClientUnaryCall;
  createContact(
    request: CreateContactRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: CreateContactResponse) => void,
  ): ClientUnaryCall;
  createContact(
    request: CreateContactRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: CreateContactResponse) => void,
  ): ClientUnaryCall;
  deleteContact(
    request: DeleteContactRequest,
    callback: (error: ServiceError | null, response: SuccessResponse) => void,
  ): ClientUnaryCall;
  deleteContact(
    request: DeleteContactRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: SuccessResponse) => void,
  ): ClientUnaryCall;
  deleteContact(
    request: DeleteContactRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: SuccessResponse) => void,
  ): ClientUnaryCall;
}

export const ContactServiceClient = makeGenericClientConstructor(
  ContactServiceService,
  "apfish.v1.user.contact.ContactService",
) as unknown as {
  new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): ContactServiceClient;
  service: typeof ContactServiceService;
  serviceName: string;
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
