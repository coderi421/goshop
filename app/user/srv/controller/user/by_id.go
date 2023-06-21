package user

import (
	"context"

	upbv1 "github.com/coderi421/goshop/api/user/v1"

	"github.com/coderi421/gframework/pkg/log"
)

func (us *userServer) GetUserById(ctx context.Context, request *upbv1.IdRequest) (*upbv1.UserInfoResponse, error) {
	user, err := us.srv.User().GetByID(ctx, uint64(request.Id))
	if err != nil {
		log.Errorf("get user by id: %s,error: %v", request.Id, err)
	}
	userInfoRsp := DTOToResponse(*user)
	return userInfoRsp, nil
}
