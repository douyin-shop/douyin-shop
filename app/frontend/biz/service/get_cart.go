package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend/common"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp *common.Empty, err error) {
	userId := h.Context.Value("user_id")

	hlog.Info("GetCartService:", userId)
	return
}
