// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/iam/credentials/v1/iamcredentials.proto

package credentials // import "google.golang.org/genproto/googleapis/iam/credentials/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IAMCredentialsClient is the client API for IAMCredentials service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IAMCredentialsClient interface {
	// Generates an OAuth 2.0 access token for a service account.
	GenerateAccessToken(ctx context.Context, in *GenerateAccessTokenRequest, opts ...grpc.CallOption) (*GenerateAccessTokenResponse, error)
	// Generates an OpenID Connect ID token for a service account.
	GenerateIdToken(ctx context.Context, in *GenerateIdTokenRequest, opts ...grpc.CallOption) (*GenerateIdTokenResponse, error)
	// Signs a blob using a service account's system-managed private key.
	SignBlob(ctx context.Context, in *SignBlobRequest, opts ...grpc.CallOption) (*SignBlobResponse, error)
	// Signs a JWT using a service account's system-managed private key.
	SignJwt(ctx context.Context, in *SignJwtRequest, opts ...grpc.CallOption) (*SignJwtResponse, error)
	// Exchange a JWT signed by third party identity provider to an OAuth 2.0
	// access token
	GenerateIdentityBindingAccessToken(ctx context.Context, in *GenerateIdentityBindingAccessTokenRequest, opts ...grpc.CallOption) (*GenerateIdentityBindingAccessTokenResponse, error)
}

type iAMCredentialsClient struct {
	cc *grpc.ClientConn
}

func NewIAMCredentialsClient(cc *grpc.ClientConn) IAMCredentialsClient {
	return &iAMCredentialsClient{cc}
}

func (c *iAMCredentialsClient) GenerateAccessToken(ctx context.Context, in *GenerateAccessTokenRequest, opts ...grpc.CallOption) (*GenerateAccessTokenResponse, error) {
	out := new(GenerateAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/google.iam.credentials.v1.IAMCredentials/GenerateAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMCredentialsClient) GenerateIdToken(ctx context.Context, in *GenerateIdTokenRequest, opts ...grpc.CallOption) (*GenerateIdTokenResponse, error) {
	out := new(GenerateIdTokenResponse)
	err := c.cc.Invoke(ctx, "/google.iam.credentials.v1.IAMCredentials/GenerateIdToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMCredentialsClient) SignBlob(ctx context.Context, in *SignBlobRequest, opts ...grpc.CallOption) (*SignBlobResponse, error) {
	out := new(SignBlobResponse)
	err := c.cc.Invoke(ctx, "/google.iam.credentials.v1.IAMCredentials/SignBlob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMCredentialsClient) SignJwt(ctx context.Context, in *SignJwtRequest, opts ...grpc.CallOption) (*SignJwtResponse, error) {
	out := new(SignJwtResponse)
	err := c.cc.Invoke(ctx, "/google.iam.credentials.v1.IAMCredentials/SignJwt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMCredentialsClient) GenerateIdentityBindingAccessToken(ctx context.Context, in *GenerateIdentityBindingAccessTokenRequest, opts ...grpc.CallOption) (*GenerateIdentityBindingAccessTokenResponse, error) {
	out := new(GenerateIdentityBindingAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/google.iam.credentials.v1.IAMCredentials/GenerateIdentityBindingAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IAMCredentialsServer is the server API for IAMCredentials service.
type IAMCredentialsServer interface {
	// Generates an OAuth 2.0 access token for a service account.
	GenerateAccessToken(context.Context, *GenerateAccessTokenRequest) (*GenerateAccessTokenResponse, error)
	// Generates an OpenID Connect ID token for a service account.
	GenerateIdToken(context.Context, *GenerateIdTokenRequest) (*GenerateIdTokenResponse, error)
	// Signs a blob using a service account's system-managed private key.
	SignBlob(context.Context, *SignBlobRequest) (*SignBlobResponse, error)
	// Signs a JWT using a service account's system-managed private key.
	SignJwt(context.Context, *SignJwtRequest) (*SignJwtResponse, error)
	// Exchange a JWT signed by third party identity provider to an OAuth 2.0
	// access token
	GenerateIdentityBindingAccessToken(context.Context, *GenerateIdentityBindingAccessTokenRequest) (*GenerateIdentityBindingAccessTokenResponse, error)
}

func RegisterIAMCredentialsServer(s *grpc.Server, srv IAMCredentialsServer) {
	s.RegisterService(&_IAMCredentials_serviceDesc, srv)
}

func _IAMCredentials_GenerateAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMCredentialsServer).GenerateAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.credentials.v1.IAMCredentials/GenerateAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMCredentialsServer).GenerateAccessToken(ctx, req.(*GenerateAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMCredentials_GenerateIdToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateIdTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMCredentialsServer).GenerateIdToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.credentials.v1.IAMCredentials/GenerateIdToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMCredentialsServer).GenerateIdToken(ctx, req.(*GenerateIdTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMCredentials_SignBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignBlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMCredentialsServer).SignBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.credentials.v1.IAMCredentials/SignBlob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMCredentialsServer).SignBlob(ctx, req.(*SignBlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMCredentials_SignJwt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignJwtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMCredentialsServer).SignJwt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.credentials.v1.IAMCredentials/SignJwt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMCredentialsServer).SignJwt(ctx, req.(*SignJwtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMCredentials_GenerateIdentityBindingAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateIdentityBindingAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMCredentialsServer).GenerateIdentityBindingAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.credentials.v1.IAMCredentials/GenerateIdentityBindingAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMCredentialsServer).GenerateIdentityBindingAccessToken(ctx, req.(*GenerateIdentityBindingAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IAMCredentials_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.iam.credentials.v1.IAMCredentials",
	HandlerType: (*IAMCredentialsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateAccessToken",
			Handler:    _IAMCredentials_GenerateAccessToken_Handler,
		},
		{
			MethodName: "GenerateIdToken",
			Handler:    _IAMCredentials_GenerateIdToken_Handler,
		},
		{
			MethodName: "SignBlob",
			Handler:    _IAMCredentials_SignBlob_Handler,
		},
		{
			MethodName: "SignJwt",
			Handler:    _IAMCredentials_SignJwt_Handler,
		},
		{
			MethodName: "GenerateIdentityBindingAccessToken",
			Handler:    _IAMCredentials_GenerateIdentityBindingAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/iam/credentials/v1/iamcredentials.proto",
}

func init() {
	proto.RegisterFile("google/iam/credentials/v1/iamcredentials.proto", fileDescriptor_iamcredentials_549bd6a612073d41)
}

var fileDescriptor_iamcredentials_549bd6a612073d41 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0xaa, 0xd4, 0x30,
	0x18, 0xc5, 0xc9, 0x5d, 0x78, 0x25, 0x0b, 0x85, 0xde, 0x95, 0x83, 0xab, 0x0a, 0x82, 0x15, 0x1a,
	0x67, 0x74, 0x14, 0x3a, 0x0a, 0x4e, 0x1d, 0x95, 0x29, 0x88, 0x83, 0x7f, 0x36, 0xee, 0x32, 0x69,
	0x08, 0xd1, 0x36, 0x5f, 0x6d, 0x32, 0x33, 0x88, 0xb8, 0x11, 0x04, 0xf7, 0x6e, 0x5d, 0xf8, 0x20,
	0x3e, 0x82, 0x3b, 0x5f, 0xc1, 0x9d, 0x2f, 0xe0, 0x52, 0xd2, 0xa6, 0x4c, 0xc5, 0xa9, 0xb6, 0xdc,
	0x65, 0xd3, 0x73, 0xce, 0xf7, 0x3b, 0x24, 0x7c, 0x38, 0x14, 0x00, 0x22, 0xe3, 0x44, 0xd2, 0x9c,
	0xb0, 0x92, 0xa7, 0x5c, 0x19, 0x49, 0x33, 0x4d, 0xb6, 0x63, 0x7b, 0xd4, 0x3a, 0x09, 0x8b, 0x12,
	0x0c, 0x78, 0x17, 0x6a, 0x7d, 0x28, 0x69, 0x1e, 0xb6, 0xff, 0x6e, 0xc7, 0xa3, 0x8b, 0x2e, 0x8a,
	0x16, 0x92, 0x50, 0xa5, 0xc0, 0x50, 0x23, 0x41, 0x39, 0xe3, 0xe8, 0x72, 0xf7, 0x20, 0x06, 0x79,
	0x0e, 0xaa, 0xd6, 0x4d, 0x7e, 0x1e, 0xe3, 0x73, 0xcb, 0xf9, 0xa3, 0x7b, 0x7b, 0x89, 0xf7, 0x0d,
	0xe1, 0x93, 0x87, 0x5c, 0xf1, 0x92, 0x1a, 0x3e, 0x67, 0x8c, 0x6b, 0xfd, 0x0c, 0x5e, 0x71, 0xe5,
	0x4d, 0xc3, 0x4e, 0x98, 0xf0, 0x80, 0xfe, 0x09, 0x7f, 0xbd, 0xe1, 0xda, 0x8c, 0x6e, 0x0e, 0xb5,
	0xe9, 0x02, 0x94, 0xe6, 0xfe, 0x83, 0xf7, 0xdf, 0x7f, 0x7c, 0x3a, 0xba, 0xeb, 0xcf, 0x2c, 0xf3,
	0x5b, 0x45, 0x73, 0x7e, 0xa7, 0x28, 0xe1, 0x25, 0x67, 0x46, 0x93, 0x80, 0x68, 0x5e, 0x6e, 0x25,
	0xb3, 0x46, 0xd8, 0x28, 0x7b, 0xf2, 0x2e, 0x12, 0x7f, 0x87, 0x45, 0x28, 0xf0, 0xbe, 0x22, 0x7c,
	0xbe, 0x99, 0xb3, 0x4c, 0xeb, 0x2a, 0xe3, 0x1e, 0x4c, 0x4e, 0xdb, 0xd4, 0x98, 0x0c, 0xb1, 0xb8,
	0x0a, 0x71, 0x55, 0xe1, 0xb6, 0x7f, 0x6b, 0x68, 0x05, 0x17, 0x64, 0xf1, 0xbf, 0x20, 0x7c, 0xf6,
	0xa9, 0x14, 0x2a, 0xce, 0x60, 0xed, 0x05, 0xff, 0x80, 0x68, 0x44, 0x0d, 0xf0, 0xd5, 0x5e, 0x5a,
	0x47, 0x3a, 0xab, 0x48, 0xa7, 0xfe, 0xb5, 0xbe, 0xa4, 0xda, 0x25, 0x58, 0xc4, 0xcf, 0x08, 0x1f,
	0xdb, 0xc4, 0x64, 0x67, 0xbc, 0x2b, 0xff, 0x99, 0x9a, 0xec, 0x4c, 0x03, 0x18, 0xf4, 0x91, 0x3a,
	0xbe, 0xa8, 0xe2, 0xbb, 0xe1, 0x93, 0x21, 0x7c, 0xc9, 0xce, 0x58, 0xbc, 0x8f, 0x47, 0xd8, 0xdf,
	0xdf, 0x90, 0x1d, 0x62, 0xde, 0xc4, 0x52, 0xa5, 0x52, 0x89, 0xf6, 0xf3, 0x5e, 0xf4, 0xba, 0xe0,
	0x2e, 0x7b, 0x53, 0xea, 0xfe, 0x29, 0x53, 0x5c, 0xdf, 0xe7, 0x55, 0xdf, 0xc7, 0x7e, 0x32, 0xfc,
	0xe5, 0x74, 0x65, 0x47, 0x28, 0x88, 0x3f, 0x20, 0x7c, 0x89, 0x41, 0xde, 0x30, 0xb2, 0x0c, 0x36,
	0xe9, 0x01, 0xd2, 0xf8, 0xe4, 0xcf, 0x95, 0xb0, 0xb2, 0xab, 0x62, 0x85, 0x5e, 0x2c, 0x9c, 0x4f,
	0x40, 0x46, 0x95, 0x08, 0xa1, 0x14, 0x44, 0x70, 0x55, 0x2d, 0x12, 0x52, 0xff, 0xa2, 0x85, 0xd4,
	0x07, 0x76, 0xce, 0xac, 0xf5, 0xf9, 0x0b, 0xa1, 0xf5, 0x99, 0xca, 0x73, 0xfd, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x2e, 0x2e, 0xa2, 0x0d, 0x0f, 0x05, 0x00, 0x00,
}
