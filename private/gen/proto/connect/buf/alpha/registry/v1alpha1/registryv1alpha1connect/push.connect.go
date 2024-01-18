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
// Source: buf/alpha/registry/v1alpha1/push.proto

package registryv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/registry/v1alpha1"
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
	// PushServiceName is the fully-qualified name of the PushService service.
	PushServiceName = "buf.alpha.registry.v1alpha1.PushService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PushServicePushProcedure is the fully-qualified name of the PushService's Push RPC.
	PushServicePushProcedure = "/buf.alpha.registry.v1alpha1.PushService/Push"
	// PushServicePushManifestAndBlobsProcedure is the fully-qualified name of the PushService's
	// PushManifestAndBlobs RPC.
	PushServicePushManifestAndBlobsProcedure = "/buf.alpha.registry.v1alpha1.PushService/PushManifestAndBlobs"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	pushServiceServiceDescriptor                    = v1alpha1.File_buf_alpha_registry_v1alpha1_push_proto.Services().ByName("PushService")
	pushServicePushMethodDescriptor                 = pushServiceServiceDescriptor.Methods().ByName("Push")
	pushServicePushManifestAndBlobsMethodDescriptor = pushServiceServiceDescriptor.Methods().ByName("PushManifestAndBlobs")
)

// PushServiceClient is a client for the buf.alpha.registry.v1alpha1.PushService service.
type PushServiceClient interface {
	// Push pushes.
	// NOTE: Newer clients should use PushManifestAndBlobs.
	Push(context.Context, *connect.Request[v1alpha1.PushRequest]) (*connect.Response[v1alpha1.PushResponse], error)
	// PushManifestAndBlobs pushes a module by encoding it in a manifest and blobs format.
	PushManifestAndBlobs(context.Context, *connect.Request[v1alpha1.PushManifestAndBlobsRequest]) (*connect.Response[v1alpha1.PushManifestAndBlobsResponse], error)
}

// NewPushServiceClient constructs a client for the buf.alpha.registry.v1alpha1.PushService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPushServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) PushServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &pushServiceClient{
		push: connect.NewClient[v1alpha1.PushRequest, v1alpha1.PushResponse](
			httpClient,
			baseURL+PushServicePushProcedure,
			connect.WithSchema(pushServicePushMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyIdempotent),
			connect.WithClientOptions(opts...),
		),
		pushManifestAndBlobs: connect.NewClient[v1alpha1.PushManifestAndBlobsRequest, v1alpha1.PushManifestAndBlobsResponse](
			httpClient,
			baseURL+PushServicePushManifestAndBlobsProcedure,
			connect.WithSchema(pushServicePushManifestAndBlobsMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyIdempotent),
			connect.WithClientOptions(opts...),
		),
	}
}

// pushServiceClient implements PushServiceClient.
type pushServiceClient struct {
	push                 *connect.Client[v1alpha1.PushRequest, v1alpha1.PushResponse]
	pushManifestAndBlobs *connect.Client[v1alpha1.PushManifestAndBlobsRequest, v1alpha1.PushManifestAndBlobsResponse]
}

// Push calls buf.alpha.registry.v1alpha1.PushService.Push.
func (c *pushServiceClient) Push(ctx context.Context, req *connect.Request[v1alpha1.PushRequest]) (*connect.Response[v1alpha1.PushResponse], error) {
	return c.push.CallUnary(ctx, req)
}

// PushManifestAndBlobs calls buf.alpha.registry.v1alpha1.PushService.PushManifestAndBlobs.
func (c *pushServiceClient) PushManifestAndBlobs(ctx context.Context, req *connect.Request[v1alpha1.PushManifestAndBlobsRequest]) (*connect.Response[v1alpha1.PushManifestAndBlobsResponse], error) {
	return c.pushManifestAndBlobs.CallUnary(ctx, req)
}

// PushServiceHandler is an implementation of the buf.alpha.registry.v1alpha1.PushService service.
type PushServiceHandler interface {
	// Push pushes.
	// NOTE: Newer clients should use PushManifestAndBlobs.
	Push(context.Context, *connect.Request[v1alpha1.PushRequest]) (*connect.Response[v1alpha1.PushResponse], error)
	// PushManifestAndBlobs pushes a module by encoding it in a manifest and blobs format.
	PushManifestAndBlobs(context.Context, *connect.Request[v1alpha1.PushManifestAndBlobsRequest]) (*connect.Response[v1alpha1.PushManifestAndBlobsResponse], error)
}

// NewPushServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPushServiceHandler(svc PushServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	pushServicePushHandler := connect.NewUnaryHandler(
		PushServicePushProcedure,
		svc.Push,
		connect.WithSchema(pushServicePushMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
		connect.WithHandlerOptions(opts...),
	)
	pushServicePushManifestAndBlobsHandler := connect.NewUnaryHandler(
		PushServicePushManifestAndBlobsProcedure,
		svc.PushManifestAndBlobs,
		connect.WithSchema(pushServicePushManifestAndBlobsMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
		connect.WithHandlerOptions(opts...),
	)
	return "/buf.alpha.registry.v1alpha1.PushService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PushServicePushProcedure:
			pushServicePushHandler.ServeHTTP(w, r)
		case PushServicePushManifestAndBlobsProcedure:
			pushServicePushManifestAndBlobsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPushServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPushServiceHandler struct{}

func (UnimplementedPushServiceHandler) Push(context.Context, *connect.Request[v1alpha1.PushRequest]) (*connect.Response[v1alpha1.PushResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.PushService.Push is not implemented"))
}

func (UnimplementedPushServiceHandler) PushManifestAndBlobs(context.Context, *connect.Request[v1alpha1.PushManifestAndBlobsRequest]) (*connect.Response[v1alpha1.PushManifestAndBlobsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.PushService.PushManifestAndBlobs is not implemented"))
}
