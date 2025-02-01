// Code generated by Kitex v0.9.1. DO NOT EDIT.

package authservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	auth "github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	DeliverTokenByRPC(ctx context.Context, Req *auth.DeliverTokenReq, callOptions ...callopt.Option) (r *auth.DeliveryResp, err error)
	VerifyTokenByRPC(ctx context.Context, Req *auth.VerifyTokenReq, callOptions ...callopt.Option) (r *auth.VerifyResp, err error)
	Logout(ctx context.Context, Req *auth.LogoutReq, callOptions ...callopt.Option) (r *auth.LogoutResp, err error)
	AddBlacklist(ctx context.Context, Req *auth.AddBlackListReq, callOptions ...callopt.Option) (r *auth.AddBlackListResp, err error)
	DeleteBlacklist(ctx context.Context, Req *auth.DeleteBlackListReq, callOptions ...callopt.Option) (r *auth.DeleteBlackListResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kAuthServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kAuthServiceClient struct {
	*kClient
}

func (p *kAuthServiceClient) DeliverTokenByRPC(ctx context.Context, Req *auth.DeliverTokenReq, callOptions ...callopt.Option) (r *auth.DeliveryResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeliverTokenByRPC(ctx, Req)
}

func (p *kAuthServiceClient) VerifyTokenByRPC(ctx context.Context, Req *auth.VerifyTokenReq, callOptions ...callopt.Option) (r *auth.VerifyResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VerifyTokenByRPC(ctx, Req)
}

func (p *kAuthServiceClient) Logout(ctx context.Context, Req *auth.LogoutReq, callOptions ...callopt.Option) (r *auth.LogoutResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Logout(ctx, Req)
}

func (p *kAuthServiceClient) AddBlacklist(ctx context.Context, Req *auth.AddBlackListReq, callOptions ...callopt.Option) (r *auth.AddBlackListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddBlacklist(ctx, Req)
}

func (p *kAuthServiceClient) DeleteBlacklist(ctx context.Context, Req *auth.DeleteBlackListReq, callOptions ...callopt.Option) (r *auth.DeleteBlackListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteBlacklist(ctx, Req)
}
