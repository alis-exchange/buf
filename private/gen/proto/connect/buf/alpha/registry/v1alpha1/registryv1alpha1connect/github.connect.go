// Copyright 2020-2024 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: buf/alpha/registry/v1alpha1/github.proto

package registryv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/alis-exchange/buf/private/gen/proto/go/buf/alpha/registry/v1alpha1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// GithubServiceName is the fully-qualified name of the GithubService service.
	GithubServiceName = "buf.alpha.registry.v1alpha1.GithubService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GithubServiceGetGithubAppConfigProcedure is the fully-qualified name of the GithubService's
	// GetGithubAppConfig RPC.
	GithubServiceGetGithubAppConfigProcedure = "/buf.alpha.registry.v1alpha1.GithubService/GetGithubAppConfig"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	githubServiceServiceDescriptor                  = v1alpha1.File_buf_alpha_registry_v1alpha1_github_proto.Services().ByName("GithubService")
	githubServiceGetGithubAppConfigMethodDescriptor = githubServiceServiceDescriptor.Methods().ByName("GetGithubAppConfig")
)

// GithubServiceClient is a client for the buf.alpha.registry.v1alpha1.GithubService service.
type GithubServiceClient interface {
	// GetGithubAppConfig returns a Github Application Configuration.
	GetGithubAppConfig(context.Context, *connect.Request[v1alpha1.GetGithubAppConfigRequest]) (*connect.Response[v1alpha1.GetGithubAppConfigResponse], error)
}

// NewGithubServiceClient constructs a client for the buf.alpha.registry.v1alpha1.GithubService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGithubServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) GithubServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &githubServiceClient{
		getGithubAppConfig: connect.NewClient[v1alpha1.GetGithubAppConfigRequest, v1alpha1.GetGithubAppConfigResponse](
			httpClient,
			baseURL+GithubServiceGetGithubAppConfigProcedure,
			connect.WithSchema(githubServiceGetGithubAppConfigMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// githubServiceClient implements GithubServiceClient.
type githubServiceClient struct {
	getGithubAppConfig *connect.Client[v1alpha1.GetGithubAppConfigRequest, v1alpha1.GetGithubAppConfigResponse]
}

// GetGithubAppConfig calls buf.alpha.registry.v1alpha1.GithubService.GetGithubAppConfig.
func (c *githubServiceClient) GetGithubAppConfig(ctx context.Context, req *connect.Request[v1alpha1.GetGithubAppConfigRequest]) (*connect.Response[v1alpha1.GetGithubAppConfigResponse], error) {
	return c.getGithubAppConfig.CallUnary(ctx, req)
}

// GithubServiceHandler is an implementation of the buf.alpha.registry.v1alpha1.GithubService
// service.
type GithubServiceHandler interface {
	// GetGithubAppConfig returns a Github Application Configuration.
	GetGithubAppConfig(context.Context, *connect.Request[v1alpha1.GetGithubAppConfigRequest]) (*connect.Response[v1alpha1.GetGithubAppConfigResponse], error)
}

// NewGithubServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGithubServiceHandler(svc GithubServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	githubServiceGetGithubAppConfigHandler := connect.NewUnaryHandler(
		GithubServiceGetGithubAppConfigProcedure,
		svc.GetGithubAppConfig,
		connect.WithSchema(githubServiceGetGithubAppConfigMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/buf.alpha.registry.v1alpha1.GithubService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GithubServiceGetGithubAppConfigProcedure:
			githubServiceGetGithubAppConfigHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGithubServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGithubServiceHandler struct{}

func (UnimplementedGithubServiceHandler) GetGithubAppConfig(context.Context, *connect.Request[v1alpha1.GetGithubAppConfigRequest]) (*connect.Response[v1alpha1.GetGithubAppConfigResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.GithubService.GetGithubAppConfig is not implemented"))
}
