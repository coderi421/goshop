package rpc

import (
	"github.com/coderi421/gframework/gmicro/registry"
	"github.com/coderi421/gframework/gmicro/registry/consul"
	"github.com/coderi421/gframework/pkg/errors"
	upb "github.com/coderi421/goshop/api/user/v1"
	"github.com/coderi421/goshop/app/pkg/code"
	"github.com/coderi421/goshop/app/pkg/options"
	"github.com/coderi421/goshop/app/shop/custom/internal/data"
	consulAPI "github.com/hashicorp/consul/api"
	"sync"
)

type grpcData struct {
	uc upb.UserClient
}

func (g *grpcData) Users() data.UserData {
	return NewUsers(g.uc)
}

var (
	rpcFactory data.DataFactory
	once       sync.Once
)

// NewDiscovery 目前是基于consul实现的  以后想换成nocos etcd等  可以直接在这换
func NewDiscovery(opts *options.RegistryOptions) (registry.Discovery, error) {
	config := consulAPI.DefaultConfig()
	config.Address = opts.Address
	config.Scheme = opts.Scheme
	cli, err := consulAPI.NewClient(config)
	if err != nil {
		return nil, errors.WithCode(code.ErrInit, "init discovery error: %s", err.Error())
	}
	return consul.New(cli, consul.WithHealthCheck(true)), nil
}

// GetDataFactoryOr rpc的连接，基于服务发现
func GetDataFactoryOr(options *options.RegistryOptions) (data.DataFactory, error) {
	var (
		discovery  registry.Discovery
		userClient upb.UserClient
		errOnce    error
		once       sync.Once
	)

	if options == nil && rpcFactory == nil {
		return nil, errors.WithCode(code.ErrInit, "init Data Factory error: options is nil")
	}
	// 单利模式
	once.Do(func() {
		var err error
		//这里负责依赖所有的rpc连接
		discovery, err = NewDiscovery(options)
		if err != nil {
			errOnce = err
			return
		}

		//实例 user rpc client
		userClient, err = NewUserServiceClient(discovery)
		if err != nil {
			errOnce = err
			return
		}

		rpcFactory = &grpcData{
			uc: userClient,
		}
	})

	if rpcFactory == nil || errOnce != nil {
		return nil, errors.WithCode(code.ErrConnectGRPC, "failed to get rpc store factory %s", errOnce.Error())
	}

	return rpcFactory, nil
}
