package service

import (
	"context"
	user "github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user"
)

type GetService struct {
	ctx context.Context
} // NewGetService new GetService
func NewGetService(ctx context.Context) *GetService {
	return &GetService{ctx: ctx}
}

// Run create note info
func (s *GetService) Run(req *user.GetReq) (resp *user.GetResp, err error) {
	// Finish your business logic.

	return
}
