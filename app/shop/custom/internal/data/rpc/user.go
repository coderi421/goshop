package rpc

import (
	"context"
	v1 "github.com/coderi421/goshop/api/user/v1"

	"github.com/coderi421/gframework/gmicro/registry"
	"github.com/coderi421/gframework/gmicro/server/rpcserver"
	"github.com/coderi421/gframework/gmicro/server/rpcserver/clientinterceptors"
	itime "github.com/coderi421/gframework/pkg/common/time"
	"github.com/coderi421/gframework/pkg/errors"
	upbv1 "github.com/coderi421/goshop/api/user/v1"
	"github.com/coderi421/goshop/app/pkg/code"
	"github.com/coderi421/goshop/app/pkg/options"
	"github.com/coderi421/goshop/app/shop/custom/internal/data"
	"time"
)

const (
	serviceName = "discovery:///user-srv"
)

func NewUserServiceClient(r registry.Discovery) (upbv1.UserClient, error) {
	conn, err := rpcserver.DialInsecure(
		context.Background(),
		rpcserver.WithEndpoint(serviceName),
		rpcserver.WithDiscovery(r),
		rpcserver.WithClientUnaryInterceptor(clientinterceptors.UnaryTracingInterceptor(options.TracerNameNoGinCtx)),
	)
	if err != nil {
		return nil, errors.WithCode(code.ErrInit, err.Error())
	}
	c := upbv1.NewUserClient(conn)
	return c, nil
}

func NewUser(uc upbv1.UserClient) *user {
	return &user{uc}
}

type user struct {
	uc upbv1.UserClient
}

var _ data.UserData = &user{}

func (u *user) Create(ctx context.Context, user *data.User) (*v1.UserInfoResponse, error) {
	protoUser := &upbv1.CreateUserInfo{
		NickName: user.NickName,
		PassWord: user.PassWord,
		Mobile:   user.Mobile,
		//
		//Birthday: userDTO.Birthday,
		//Gender:   userDTO.Gender,
	}
	return u.uc.CreateUser(ctx, protoUser)
}

func (u *user) Update(ctx context.Context, user *data.User) error {
	protoUser := &upbv1.UpdateUserInfo{
		Id:       user.ID,
		NickName: user.NickName,
		Gender:   user.Gender,
		BirthDay: uint64(user.Birthday.Unix()),
	}
	_, err := u.uc.UpdateUser(ctx, protoUser)
	return err
}

func (u *user) Get(ctx context.Context, userID int32) (*data.User, error) {
	user, err := u.uc.GetUserById(ctx, &upbv1.IdRequest{Id: userID})
	if err != nil {
		return nil, err
	}
	return &data.User{
		ID:       user.Id,
		Mobile:   user.Mobile,
		NickName: user.NickName,
		Birthday: itime.Time{time.Unix(int64(user.BirthDay), 0)}.Time,
		Gender:   user.Gender,
		Role:     user.Role,
		PassWord: user.PassWord,
	}, nil
}

func (u *user) GetByMobile(ctx context.Context, mobile string) (*data.User, error) {
	user, err := u.uc.GetUserByMobile(ctx, &upbv1.MobileRequest{Mobile: mobile})
	if err != nil {
		return nil, err
	}
	return &data.User{
		ID:       user.Id,
		Mobile:   user.Mobile,
		NickName: user.NickName,
		Birthday: itime.Time{time.Unix(int64(user.BirthDay), 0)}.Time,
		Gender:   user.Gender,
		Role:     user.Role,
		PassWord: user.PassWord,
	}, nil
}

func (u *user) CheckPassWord(ctx context.Context, password, encryptedPwd string) (ok bool, err error) {
	cres, err := u.uc.CheckPassWord(ctx, &upbv1.PasswordCheckInfo{
		Password:          password,
		EncryptedPassword: encryptedPwd,
	})
	return cres.Success, err
}
