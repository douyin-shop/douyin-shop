package service

import (
	"context"
	auth "github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
)

type DeleteBlacklistService struct {
	ctx context.Context
} // NewDeleteBlacklistService new DeleteBlacklistService
func NewDeleteBlacklistService(ctx context.Context) *DeleteBlacklistService {
	return &DeleteBlacklistService{ctx: ctx}
}

// Run create note info
func (s *DeleteBlacklistService) Run(req *auth.DeleteBlackListReq) (resp *auth.DeleteBlackListResp, err error) {
	// Finish your business logic.

	return
}
