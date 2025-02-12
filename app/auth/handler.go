package main

import (
	"context"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/service"
	auth "github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	resp, err = service.NewDeliverTokenByRPCService(ctx).Run(req)

	return resp, err
}

// VerifyTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	resp, err = service.NewVerifyTokenByRPCService(ctx).Run(req)

	return resp, err
}

// AddBlacklist implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) AddBlacklist(ctx context.Context, req *auth.AddBlackListReq) (resp *auth.AddBlackListResp, err error) {
	resp, err = service.NewAddBlacklistService(ctx).Run(req)

	return resp, err
}

// DeleteBlacklist implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeleteBlacklist(ctx context.Context, req *auth.DeleteBlackListReq) (resp *auth.DeleteBlackListResp, err error) {
	resp, err = service.NewDeleteBlacklistService(ctx).Run(req)

	return resp, err
}

// Logout implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) Logout(ctx context.Context, req *auth.LogoutReq) (resp *auth.LogoutResp, err error) {
	resp, err = service.NewLogoutService(ctx).Run(req)

	return resp, err
}
