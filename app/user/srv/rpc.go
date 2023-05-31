package srv

import (
	"fmt"

	"github.com/coderi421/gframework/pkg/log"
	"github.com/coderi421/goshop/app/user/srv/data/v1/db"

	"github.com/coderi421/gframework/gmicro/core/trace"

	upbv1 "github.com/coderi421/goshop/api/user/v1"
	srv1 "github.com/coderi421/goshop/app/user/srv/service/v1"

	"github.com/coderi421/gframework/gmicro/server/rpcserver"
	"github.com/coderi421/goshop/app/user/srv/config"
	"github.com/coderi421/goshop/app/user/srv/controller/user"
)

func NewUserRPCServer(cfg *config.Config) (*rpcserver.Server, error) {
	// 初始化 open-telemetry 的 exporter
	// 这里会根据 endpoint 为单元注册 trace 服务的实例
	trace.InitAgent(trace.Options{
		Name:     cfg.Telemetry.Name,
		Endpoint: cfg.Telemetry.Endpoint,
		Sampler:  cfg.Telemetry.Sampler,
		Batcher:  cfg.Telemetry.Batcher,
	})

	//data := mock.NewUsers()

	gormDB, err := db.GetDBFactoryOr(cfg.MySQLOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	data := db.NewUser(gormDB)
	srv := srv1.NewUserService(data)
	userver := user.NewUserServer(srv)

	rpcAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	urpcServer := rpcserver.NewServer(
		rpcserver.WithAddress(rpcAddr),
		rpcserver.WithEnableMetric(cfg.Server.EnableMetrics),
	)

	// 注册 user 模块的 rpc 服务
	upbv1.RegisterUserServer(urpcServer.Server, userver)

	return urpcServer, nil

}
