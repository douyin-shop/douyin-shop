package service

import (
	"context"
<<<<<<< HEAD
=======
<<<<<<< Updated upstream

=======
>>>>>>> Stashed changes
>>>>>>> ae6c4a5 (测试)
	user "github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user"
)

type DeleteService struct {
	ctx context.Context
} // NewDeleteService new DeleteService
func NewDeleteService(ctx context.Context) *DeleteService {
	return &DeleteService{ctx: ctx}
}

// Run create note info
func (s *DeleteService) Run(req *user.DeleteReq) (resp *user.DeleteResp, err error) {
<<<<<<< HEAD
	// Finish your business logic.

=======
<<<<<<< Updated upstream
	
=======
	// Finish your business logic.

>>>>>>> Stashed changes
>>>>>>> ae6c4a5 (测试)
	return
}
