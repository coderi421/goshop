package rpc

import (
	"context"
	"github.com/coderi421/gframework/gmicro/registry"
	"github.com/coderi421/gframework/gmicro/server/rpcserver"
	"github.com/coderi421/gframework/gmicro/server/rpcserver/clientinterceptors"
	"github.com/coderi421/gframework/pkg/errors"
	upbv1 "github.com/coderi421/goshop/api/user/v1"
	"github.com/coderi421/goshop/app/pkg/code"
	"github.com/coderi421/goshop/app/pkg/options"
	"github.com/coderi421/goshop/app/shop/custom/internal/data"
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

func NewUsers(uc upbv1.UserClient) *users {
	return &users{uc}
}

type users struct {
	uc upbv1.UserClient
}

var _ data.UserData = &users{}

func (u *users) Create(ctx context.Context, user *data.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *users) Update(ctx context.Context, user *data.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *users) Get(ctx context.Context, userID uint64) (data.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *users) GetByMobile(ctx context.Context, mobile string) (data.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *users) CheckPassWord(ctx context.Context, password, encryptedPwd string) error {
	//TODO implement me
	panic("implement me")
}
