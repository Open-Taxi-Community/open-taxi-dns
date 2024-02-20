// @generated by protoc-gen-connect-es v1.3.0 with parameter "target=ts,import_extension=none"
// @generated from file dnsapi/v1/dns.proto (package dnsapi.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AddDomainRequest, AddDomainResponse, GetDomainRequest, GetDomainResponse, RemoveDomainRequest, RemoveDomainResponse } from "./dns_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * DNS service
 *
 * The DNS service is used to interact with the OpenTaxi DNS server.
 *
 * @generated from service dnsapi.v1.DnsService
 */
export const DnsService = {
  typeName: "dnsapi.v1.DnsService",
  methods: {
    /**
     * Add a new domain
     *
     * Adding a new domain will push the domain to the driver server
     * compliance blockchain. Other driver servers will then be able
     * to resolve the domain and make requests to the domain to begin
     * the compliance process.
     *
     * @generated from rpc dnsapi.v1.DnsService.AddDomain
     */
    addDomain: {
      name: "AddDomain",
      I: AddDomainRequest,
      O: AddDomainResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Get a domain
     *
     * This is used to check if a domain is in the OpenTaxi DNS
     *
     * @generated from rpc dnsapi.v1.DnsService.GetDomain
     */
    getDomain: {
      name: "GetDomain",
      I: GetDomainRequest,
      O: GetDomainResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Remove a domain
     *
     * This removes a domain from the OpenTaxi DNS
     *
     * @generated from rpc dnsapi.v1.DnsService.RemoveDomain
     */
    removeDomain: {
      name: "RemoveDomain",
      I: RemoveDomainRequest,
      O: RemoveDomainResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

