package user

import (
	"context"

	"github.com/coderi421/gframework/pkg/log"
	upbv1 "github.com/coderi421/goshop/api/user/v1"
	v12 "github.com/coderi421/goshop/app/user/srv/data/v1"
	v1 "github.com/coderi421/goshop/app/user/srv/service/v1"
)

// CreateUser controller层应该是很薄的一层，参数校验，日志打印，错误处理，调用service层
func (us *userServer) CreateUser(ctx context.Context, request *upbv1.CreateUserInfo) (*upbv1.UserInfoResponse, error) {
	userDO := v12.UserDO{
		Mobile:   request.Mobile,
		NickName: request.NickName,
		Password: request.PassWord,
	}
	userDTO := v1.UserDTO{userDO}

	err := us.srv.User().Create(ctx, &userDTO)
	if err != nil {
		log.Errorf("get user by mobile: %s,error: %v", request.Mobile, err)
	}
	userInfoRsp := DTOToResponse(userDTO)
	return userInfoRsp, nil
}
