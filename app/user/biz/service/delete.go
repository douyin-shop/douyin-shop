package service

import (
	"context"
<<<<<<< Updated upstream

=======
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
	
=======
	// Finish your business logic.

>>>>>>> Stashed changes
	return
}
