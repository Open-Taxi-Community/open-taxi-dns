// @generated by protoc-gen-es v1.7.2 with parameter "target=ts,import_extension=none"
// @generated from file dnsapi/v1/dns.proto (package dnsapi.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * Requests
 *
 * @generated from message dnsapi.v1.AddDomainRequest
 */
export class AddDomainRequest extends Message<AddDomainRequest> {
  /**
   * @generated from field: string serverDomain = 1;
   */
  serverDomain = "";

  constructor(data?: PartialMessage<AddDomainRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dnsapi.v1.AddDomainRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "serverDomain", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddDomainRequest {
    return new AddDomainRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddDomainRequest {
    return new AddDomainRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddDomainRequest {
    return new AddDomainRequest().fromJsonString(jsonString, options);
  }

  static equals(a: AddDomainRequest | PlainMessage<AddDomainRequest> | undefined, b: AddDomainRequest | PlainMessage<AddDomainRequest> | undefined): boolean {
    return proto3.util.equals(AddDomainRequest, a, b);
  }
}

/**
 * @generated from message dnsapi.v1.GetDomainRequest
 */
export class GetDomainRequest extends Message<GetDomainRequest> {
  /**
   * @generated from field: string serverDomain = 1;
   */
  serverDomain = "";

  constructor(data?: PartialMessage<GetDomainRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dnsapi.v1.GetDomainRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "serverDomain", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetDomainRequest {
    return new GetDomainRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetDomainRequest {
    return new GetDomainRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetDomainRequest {
    return new GetDomainRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetDomainRequest | PlainMessage<GetDomainRequest> | undefined, b: GetDomainRequest | PlainMessage<GetDomainRequest> | undefined): boolean {
    return proto3.util.equals(GetDomainRequest, a, b);
  }
}

/**
 * @generated from message dnsapi.v1.RemoveDomainRequest
 */
export class RemoveDomainRequest extends Message<RemoveDomainRequest> {
  /**
   * @generated from field: string serverDomain = 1;
   */
  serverDomain = "";

  constructor(data?: PartialMessage<RemoveDomainRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dnsapi.v1.RemoveDomainRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "serverDomain", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RemoveDomainRequest {
    return new RemoveDomainRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RemoveDomainRequest {
    return new RemoveDomainRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RemoveDomainRequest {
    return new RemoveDomainRequest().fromJsonString(jsonString, options);
  }

  static equals(a: RemoveDomainRequest | PlainMessage<RemoveDomainRequest> | undefined, b: RemoveDomainRequest | PlainMessage<RemoveDomainRequest> | undefined): boolean {
    return proto3.util.equals(RemoveDomainRequest, a, b);
  }
}

/**
 * Responses
 *
 * @generated from message dnsapi.v1.AddDomainResponse
 */
export class AddDomainResponse extends Message<AddDomainResponse> {
  constructor(data?: PartialMessage<AddDomainResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dnsapi.v1.AddDomainResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddDomainResponse {
    return new AddDomainResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddDomainResponse {
    return new AddDomainResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddDomainResponse {
    return new AddDomainResponse().fromJsonString(jsonString, options);
  }

  static equals(a: AddDomainResponse | PlainMessage<AddDomainResponse> | undefined, b: AddDomainResponse | PlainMessage<AddDomainResponse> | undefined): boolean {
    return proto3.util.equals(AddDomainResponse, a, b);
  }
}

/**
 * @generated from message dnsapi.v1.GetDomainResponse
 */
export class GetDomainResponse extends Message<GetDomainResponse> {
  constructor(data?: PartialMessage<GetDomainResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dnsapi.v1.GetDomainResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetDomainResponse {
    return new GetDomainResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetDomainResponse {
    return new GetDomainResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetDomainResponse {
    return new GetDomainResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetDomainResponse | PlainMessage<GetDomainResponse> | undefined, b: GetDomainResponse | PlainMessage<GetDomainResponse> | undefined): boolean {
    return proto3.util.equals(GetDomainResponse, a, b);
  }
}

/**
 * @generated from message dnsapi.v1.RemoveDomainResponse
 */
export class RemoveDomainResponse extends Message<RemoveDomainResponse> {
  constructor(data?: PartialMessage<RemoveDomainResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dnsapi.v1.RemoveDomainResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RemoveDomainResponse {
    return new RemoveDomainResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RemoveDomainResponse {
    return new RemoveDomainResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RemoveDomainResponse {
    return new RemoveDomainResponse().fromJsonString(jsonString, options);
  }

  static equals(a: RemoveDomainResponse | PlainMessage<RemoveDomainResponse> | undefined, b: RemoveDomainResponse | PlainMessage<RemoveDomainResponse> | undefined): boolean {
    return proto3.util.equals(RemoveDomainResponse, a, b);
  }
}
