package custom

import (
	"context"
	gapp "github.com/coderi421/gframework/gmicro/app"
	"github.com/coderi421/gframework/gmicro/registry"
	"github.com/coderi421/gframework/gmicro/registry/consul"
	"github.com/coderi421/gframework/pkg/app"
	"github.com/coderi421/gframework/pkg/log"
	"github.com/coderi421/gframework/pkg/storage"
	"github.com/coderi421/goshop/app/pkg/options"
	"github.com/coderi421/goshop/app/shop/custom/config"
	"github.com/hashicorp/consul/api"
)

func NewApp(basename string) *app.App {
	cfg := config.New()
	appl := app.NewApp("custom",
		"shop",
		app.WithOptions(cfg),
		app.WithRunFunc(run(cfg)),
		//不读配置 使用命令行参数时使用
		//app.WithNoConfig(),
	)
	return appl
}
func NewRegistrar(registry *options.RegistryOptions) registry.Registrar {
	c := api.DefaultConfig()
	c.Address = registry.Address
	c.Scheme = registry.Scheme
	cli, err := api.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(true))
	return r
}
func NewUserApp(cfg *config.Config) (*gapp.App, error) {
	//初始化log
	log.Init(cfg.Log)
	defer log.Flush()
	//服务注册
	register := NewRegistrar(cfg.Registry)

	//连接redis
	redisConfig := &storage.Config{
		Host:                  cfg.Redis.Host,
		Port:                  cfg.Redis.Port,
		Addrs:                 cfg.Redis.Addrs,
		MasterName:            cfg.Redis.MasterName,
		Username:              cfg.Redis.Username,
		Password:              cfg.Redis.Password,
		Database:              cfg.Redis.Database,
		MaxIdle:               cfg.Redis.MaxIdle,
		MaxActive:             cfg.Redis.MaxActive,
		Timeout:               cfg.Redis.Timeout,
		EnableCluster:         cfg.Redis.EnableCluster,
		UseSSL:                cfg.Redis.UseSSL,
		SSLInsecureSkipVerify: cfg.Redis.SSLInsecureSkipVerify,
		EnableTracing:         cfg.Redis.EnableTracing,
	}
	go storage.ConnectToRedis(context.Background(), redisConfig)

	//生成http服务
	restServer, err := NewCustomHTTPServer(cfg)
	if err != nil {
		return nil, err
	}
	return gapp.New(
		gapp.WithName(cfg.Server.Name),
		gapp.WithRestServer(restServer),
		gapp.WithRegistrar(register),
	), nil

}

// controller(参数校验) ->service(具体的业务逻辑)->(数据库的接口)
func run(cfg *config.Config) app.RunFunc {
	return func(baseName string) error {
		userApp, err := NewUserApp(cfg)
		if err != nil {
			return err
		}
		//启动
		if err := userApp.Run(); err != nil {
			log.Errorf("run user app error: %s", err)
		}
		return nil
	}
}
