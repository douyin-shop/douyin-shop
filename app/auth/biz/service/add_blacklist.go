package service

import (
	"context"
	auth "github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
)

type AddBlacklistService struct {
	ctx context.Context
} // NewAddBlacklistService new AddBlacklistService
func NewAddBlacklistService(ctx context.Context) *AddBlacklistService {
	return &AddBlacklistService{ctx: ctx}
}

// Run create note info
func (s *AddBlacklistService) Run(req *auth.AddBlackListReq) (resp *auth.AddBlackListResp, err error) {
	// Finish your business logic.

	return
}
