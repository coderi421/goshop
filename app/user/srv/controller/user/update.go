package user

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/coderi421/gframework/pkg/log"
	upbv1 "github.com/coderi421/goshop/api/user/v1"
	v1 "github.com/coderi421/goshop/app/user/srv/data/v1"
	v12 "github.com/coderi421/goshop/app/user/srv/service/v1"
)

// controller层应该是很薄的一层，参数校验，日志打印，错误处理，调用service层
func (us *userServer) UpdateUser(ctx context.Context, request *upbv1.UpdateUserInfo) (*emptypb.Empty, error) {
	birthDay := time.Unix(int64(request.BirthDay), 0)
	userDO := v1.UserDO{
		BaseModel: v1.BaseModel{
			ID: request.Id,
		},
		NickName: request.NickName,
		Gender:   request.Gender,
		Birthday: &birthDay,
	}
	userDTO := v12.UserDTO{userDO}
	err := us.srv.Update(ctx, &userDTO)
	if err != nil {
		log.Errorf("update user: %v,error: %v", userDTO, err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
