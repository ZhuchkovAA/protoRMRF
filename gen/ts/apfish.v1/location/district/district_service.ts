// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v1.181.2
//   protoc               v4.25.1
// source: apfish.v1/location/district/district_service.proto

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
import { District } from "./district";
import { DistrictSummary } from "./summary/district_summary";

export const protobufPackage = "apfish.v1.location.district";

export interface DistrictRequest {
  id: string;
}

export interface DistrictSummaryResponse {
  district: DistrictSummary | undefined;
}

export interface ListDistrictsRequest {
  /** Page number (1-based). Default: 1. */
  page: number;
  /** Items per page (default: 20, max: 100). */
  perPage: number;
}

export interface ListDistrictsResponse {
  listDistricts: DistrictSummary[];
  total: number;
}

function createBaseDistrictRequest(): DistrictRequest {
  return { id: "" };
}

export const DistrictRequest = {
  encode(message: DistrictRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DistrictRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDistrictRequest();
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

  fromJSON(object: any): DistrictRequest {
    return { id: isSet(object.id) ? globalThis.String(object.id) : "" };
  },

  toJSON(message: DistrictRequest): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DistrictRequest>, I>>(base?: I): DistrictRequest {
    return DistrictRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DistrictRequest>, I>>(object: I): DistrictRequest {
    const message = createBaseDistrictRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseDistrictSummaryResponse(): DistrictSummaryResponse {
  return { district: undefined };
}

export const DistrictSummaryResponse = {
  encode(message: DistrictSummaryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.district !== undefined) {
      DistrictSummary.encode(message.district, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DistrictSummaryResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDistrictSummaryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.district = DistrictSummary.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DistrictSummaryResponse {
    return { district: isSet(object.district) ? DistrictSummary.fromJSON(object.district) : undefined };
  },

  toJSON(message: DistrictSummaryResponse): unknown {
    const obj: any = {};
    if (message.district !== undefined) {
      obj.district = DistrictSummary.toJSON(message.district);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DistrictSummaryResponse>, I>>(base?: I): DistrictSummaryResponse {
    return DistrictSummaryResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DistrictSummaryResponse>, I>>(object: I): DistrictSummaryResponse {
    const message = createBaseDistrictSummaryResponse();
    message.district = (object.district !== undefined && object.district !== null)
      ? DistrictSummary.fromPartial(object.district)
      : undefined;
    return message;
  },
};

function createBaseListDistrictsRequest(): ListDistrictsRequest {
  return { page: 0, perPage: 0 };
}

export const ListDistrictsRequest = {
  encode(message: ListDistrictsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.page !== 0) {
      writer.uint32(8).int32(message.page);
    }
    if (message.perPage !== 0) {
      writer.uint32(16).int32(message.perPage);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListDistrictsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListDistrictsRequest();
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

  fromJSON(object: any): ListDistrictsRequest {
    return {
      page: isSet(object.page) ? globalThis.Number(object.page) : 0,
      perPage: isSet(object.perPage) ? globalThis.Number(object.perPage) : 0,
    };
  },

  toJSON(message: ListDistrictsRequest): unknown {
    const obj: any = {};
    if (message.page !== 0) {
      obj.page = Math.round(message.page);
    }
    if (message.perPage !== 0) {
      obj.perPage = Math.round(message.perPage);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListDistrictsRequest>, I>>(base?: I): ListDistrictsRequest {
    return ListDistrictsRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ListDistrictsRequest>, I>>(object: I): ListDistrictsRequest {
    const message = createBaseListDistrictsRequest();
    message.page = object.page ?? 0;
    message.perPage = object.perPage ?? 0;
    return message;
  },
};

function createBaseListDistrictsResponse(): ListDistrictsResponse {
  return { listDistricts: [], total: 0 };
}

export const ListDistrictsResponse = {
  encode(message: ListDistrictsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.listDistricts) {
      DistrictSummary.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.total !== 0) {
      writer.uint32(16).int32(message.total);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListDistrictsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListDistrictsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.listDistricts.push(DistrictSummary.decode(reader, reader.uint32()));
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

  fromJSON(object: any): ListDistrictsResponse {
    return {
      listDistricts: globalThis.Array.isArray(object?.listDistricts)
        ? object.listDistricts.map((e: any) => DistrictSummary.fromJSON(e))
        : [],
      total: isSet(object.total) ? globalThis.Number(object.total) : 0,
    };
  },

  toJSON(message: ListDistrictsResponse): unknown {
    const obj: any = {};
    if (message.listDistricts?.length) {
      obj.listDistricts = message.listDistricts.map((e) => DistrictSummary.toJSON(e));
    }
    if (message.total !== 0) {
      obj.total = Math.round(message.total);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListDistrictsResponse>, I>>(base?: I): ListDistrictsResponse {
    return ListDistrictsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ListDistrictsResponse>, I>>(object: I): ListDistrictsResponse {
    const message = createBaseListDistrictsResponse();
    message.listDistricts = object.listDistricts?.map((e) => DistrictSummary.fromPartial(e)) || [];
    message.total = object.total ?? 0;
    return message;
  },
};

export type DistrictServiceService = typeof DistrictServiceService;
export const DistrictServiceService = {
  getDistrict: {
    path: "/apfish.v1.location.district.DistrictService/GetDistrict",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: DistrictRequest) => Buffer.from(DistrictRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => DistrictRequest.decode(value),
    responseSerialize: (value: District) => Buffer.from(District.encode(value).finish()),
    responseDeserialize: (value: Buffer) => District.decode(value),
  },
  getDistrictSummary: {
    path: "/apfish.v1.location.district.DistrictService/GetDistrictSummary",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: DistrictRequest) => Buffer.from(DistrictRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => DistrictRequest.decode(value),
    responseSerialize: (value: DistrictSummaryResponse) => Buffer.from(DistrictSummaryResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => DistrictSummaryResponse.decode(value),
  },
  listDistricts: {
    path: "/apfish.v1.location.district.DistrictService/ListDistricts",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: ListDistrictsRequest) => Buffer.from(ListDistrictsRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => ListDistrictsRequest.decode(value),
    responseSerialize: (value: ListDistrictsResponse) => Buffer.from(ListDistrictsResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => ListDistrictsResponse.decode(value),
  },
} as const;

export interface DistrictServiceServer extends UntypedServiceImplementation {
  getDistrict: handleUnaryCall<DistrictRequest, District>;
  getDistrictSummary: handleUnaryCall<DistrictRequest, DistrictSummaryResponse>;
  listDistricts: handleUnaryCall<ListDistrictsRequest, ListDistrictsResponse>;
}

export interface DistrictServiceClient extends Client {
  getDistrict(
    request: DistrictRequest,
    callback: (error: ServiceError | null, response: District) => void,
  ): ClientUnaryCall;
  getDistrict(
    request: DistrictRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: District) => void,
  ): ClientUnaryCall;
  getDistrict(
    request: DistrictRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: District) => void,
  ): ClientUnaryCall;
  getDistrictSummary(
    request: DistrictRequest,
    callback: (error: ServiceError | null, response: DistrictSummaryResponse) => void,
  ): ClientUnaryCall;
  getDistrictSummary(
    request: DistrictRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: DistrictSummaryResponse) => void,
  ): ClientUnaryCall;
  getDistrictSummary(
    request: DistrictRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: DistrictSummaryResponse) => void,
  ): ClientUnaryCall;
  listDistricts(
    request: ListDistrictsRequest,
    callback: (error: ServiceError | null, response: ListDistrictsResponse) => void,
  ): ClientUnaryCall;
  listDistricts(
    request: ListDistrictsRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: ListDistrictsResponse) => void,
  ): ClientUnaryCall;
  listDistricts(
    request: ListDistrictsRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: ListDistrictsResponse) => void,
  ): ClientUnaryCall;
}

export const DistrictServiceClient = makeGenericClientConstructor(
  DistrictServiceService,
  "apfish.v1.location.district.DistrictService",
) as unknown as {
  new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): DistrictServiceClient;
  service: typeof DistrictServiceService;
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
