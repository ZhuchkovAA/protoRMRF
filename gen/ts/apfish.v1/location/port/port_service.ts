// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v1.181.2
//   protoc               v4.25.1
// source: apfish.v1/location/port/port_service.proto

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
import { Port } from "./port";
import { PortSummary } from "./summary/port_summary";

export const protobufPackage = "apfish.v1.location.port";

export interface PortRequest {
  id: string;
}

export interface PortSummaryResponse {
  port: PortSummary | undefined;
}

export interface ListPortsRequest {
  /** Page number (1-based). Default: 1. */
  page: number;
  /** Items per page (default: 20, max: 100). */
  perPage: number;
}

export interface ListPortsResponse {
  listPorts: PortSummary[];
  total: number;
}

function createBasePortRequest(): PortRequest {
  return { id: "" };
}

export const PortRequest = {
  encode(message: PortRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PortRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePortRequest();
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

  fromJSON(object: any): PortRequest {
    return { id: isSet(object.id) ? globalThis.String(object.id) : "" };
  },

  toJSON(message: PortRequest): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<PortRequest>, I>>(base?: I): PortRequest {
    return PortRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<PortRequest>, I>>(object: I): PortRequest {
    const message = createBasePortRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBasePortSummaryResponse(): PortSummaryResponse {
  return { port: undefined };
}

export const PortSummaryResponse = {
  encode(message: PortSummaryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.port !== undefined) {
      PortSummary.encode(message.port, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PortSummaryResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePortSummaryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.port = PortSummary.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PortSummaryResponse {
    return { port: isSet(object.port) ? PortSummary.fromJSON(object.port) : undefined };
  },

  toJSON(message: PortSummaryResponse): unknown {
    const obj: any = {};
    if (message.port !== undefined) {
      obj.port = PortSummary.toJSON(message.port);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<PortSummaryResponse>, I>>(base?: I): PortSummaryResponse {
    return PortSummaryResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<PortSummaryResponse>, I>>(object: I): PortSummaryResponse {
    const message = createBasePortSummaryResponse();
    message.port = (object.port !== undefined && object.port !== null)
      ? PortSummary.fromPartial(object.port)
      : undefined;
    return message;
  },
};

function createBaseListPortsRequest(): ListPortsRequest {
  return { page: 0, perPage: 0 };
}

export const ListPortsRequest = {
  encode(message: ListPortsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.page !== 0) {
      writer.uint32(8).int32(message.page);
    }
    if (message.perPage !== 0) {
      writer.uint32(16).int32(message.perPage);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListPortsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListPortsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.page = reader.int32();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.perPage = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListPortsRequest {
    return {
      page: isSet(object.page) ? globalThis.Number(object.page) : 0,
      perPage: isSet(object.perPage) ? globalThis.Number(object.perPage) : 0,
    };
  },

  toJSON(message: ListPortsRequest): unknown {
    const obj: any = {};
    if (message.page !== 0) {
      obj.page = Math.round(message.page);
    }
    if (message.perPage !== 0) {
      obj.perPage = Math.round(message.perPage);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListPortsRequest>, I>>(base?: I): ListPortsRequest {
    return ListPortsRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ListPortsRequest>, I>>(object: I): ListPortsRequest {
    const message = createBaseListPortsRequest();
    message.page = object.page ?? 0;
    message.perPage = object.perPage ?? 0;
    return message;
  },
};

function createBaseListPortsResponse(): ListPortsResponse {
  return { listPorts: [], total: 0 };
}

export const ListPortsResponse = {
  encode(message: ListPortsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.listPorts) {
      PortSummary.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.total !== 0) {
      writer.uint32(16).int32(message.total);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListPortsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListPortsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.listPorts.push(PortSummary.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.total = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListPortsResponse {
    return {
      listPorts: globalThis.Array.isArray(object?.listPorts)
        ? object.listPorts.map((e: any) => PortSummary.fromJSON(e))
        : [],
      total: isSet(object.total) ? globalThis.Number(object.total) : 0,
    };
  },

  toJSON(message: ListPortsResponse): unknown {
    const obj: any = {};
    if (message.listPorts?.length) {
      obj.listPorts = message.listPorts.map((e) => PortSummary.toJSON(e));
    }
    if (message.total !== 0) {
      obj.total = Math.round(message.total);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListPortsResponse>, I>>(base?: I): ListPortsResponse {
    return ListPortsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ListPortsResponse>, I>>(object: I): ListPortsResponse {
    const message = createBaseListPortsResponse();
    message.listPorts = object.listPorts?.map((e) => PortSummary.fromPartial(e)) || [];
    message.total = object.total ?? 0;
    return message;
  },
};

export type PortServiceService = typeof PortServiceService;
export const PortServiceService = {
  getPort: {
    path: "/apfish.v1.location.port.PortService/GetPort",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: PortRequest) => Buffer.from(PortRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => PortRequest.decode(value),
    responseSerialize: (value: Port) => Buffer.from(Port.encode(value).finish()),
    responseDeserialize: (value: Buffer) => Port.decode(value),
  },
  getPortSummary: {
    path: "/apfish.v1.location.port.PortService/GetPortSummary",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: PortRequest) => Buffer.from(PortRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => PortRequest.decode(value),
    responseSerialize: (value: PortSummaryResponse) => Buffer.from(PortSummaryResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => PortSummaryResponse.decode(value),
  },
  listPorts: {
    path: "/apfish.v1.location.port.PortService/ListPorts",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: ListPortsRequest) => Buffer.from(ListPortsRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => ListPortsRequest.decode(value),
    responseSerialize: (value: ListPortsResponse) => Buffer.from(ListPortsResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => ListPortsResponse.decode(value),
  },
} as const;

export interface PortServiceServer extends UntypedServiceImplementation {
  getPort: handleUnaryCall<PortRequest, Port>;
  getPortSummary: handleUnaryCall<PortRequest, PortSummaryResponse>;
  listPorts: handleUnaryCall<ListPortsRequest, ListPortsResponse>;
}

export interface PortServiceClient extends Client {
  getPort(request: PortRequest, callback: (error: ServiceError | null, response: Port) => void): ClientUnaryCall;
  getPort(
    request: PortRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: Port) => void,
  ): ClientUnaryCall;
  getPort(
    request: PortRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: Port) => void,
  ): ClientUnaryCall;
  getPortSummary(
    request: PortRequest,
    callback: (error: ServiceError | null, response: PortSummaryResponse) => void,
  ): ClientUnaryCall;
  getPortSummary(
    request: PortRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: PortSummaryResponse) => void,
  ): ClientUnaryCall;
  getPortSummary(
    request: PortRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: PortSummaryResponse) => void,
  ): ClientUnaryCall;
  listPorts(
    request: ListPortsRequest,
    callback: (error: ServiceError | null, response: ListPortsResponse) => void,
  ): ClientUnaryCall;
  listPorts(
    request: ListPortsRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: ListPortsResponse) => void,
  ): ClientUnaryCall;
  listPorts(
    request: ListPortsRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: ListPortsResponse) => void,
  ): ClientUnaryCall;
}

export const PortServiceClient = makeGenericClientConstructor(
  PortServiceService,
  "apfish.v1.location.port.PortService",
) as unknown as {
  new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): PortServiceClient;
  service: typeof PortServiceService;
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
