package dnsapi

import (
	"context"
	v1 "github.com/Open-Taxi-Community/open-taxi-dns/pkg/dnsapi/v1"

	connect_go "connectrpc.com/connect"
)

type DnsApiServer struct {
}

func NewDnsApiServer() *DnsApiServer {
	return &DnsApiServer{}
}

func (s *DnsApiServer) AddDomain(ctx context.Context, req *connect_go.Request[v1.AddDomainRequest]) (*connect_go.Response[v1.AddDomainResponse], error) {

	return &connect_go.Response[v1.AddDomainResponse]{Msg: &v1.AddDomainResponse{}}, nil
}

func (s *DnsApiServer) GetDomain(ctx context.Context, req *connect_go.Request[v1.GetDomainRequest]) (*connect_go.Response[v1.GetDomainResponse], error) {

	return &connect_go.Response[v1.GetDomainResponse]{Msg: &v1.GetDomainResponse{}}, nil
}

func (s *DnsApiServer) RemoveDomain(ctx context.Context, req *connect_go.Request[v1.RemoveDomainRequest]) (*connect_go.Response[v1.RemoveDomainResponse], error) {

	return &connect_go.Response[v1.RemoveDomainResponse]{Msg: &v1.RemoveDomainResponse{}}, nil
}
