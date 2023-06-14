package config

import (
	"github.com/coderi421/gframework/pkg/app"
	cliflag "github.com/coderi421/gframework/pkg/common/cli/flag"
	"github.com/coderi421/gframework/pkg/log"
	"github.com/coderi421/goshop/app/pkg/options"
)

var _ app.CliOptions = &Config{}

func New() *Config {
	return &Config{
		Log:      log.NewOptions(),
		Server:   options.NewServerOptions(),
		Registry: options.NewRegistryOptions(),
		Jwt:      options.NewJwtOptions(),
		Redis:    options.NewRedisOptions(),
		Sms:      options.NewSmsOptions(),
	}
}

type Config struct {
	Log      *log.Options             `json:"log" mapstructure:"log"`
	Server   *options.ServerOptions   `json:"service" mapstructure:"service"`   // 服务发现
	Registry *options.RegistryOptions `json:"registry" mapstructure:"registry"` // 注册中心
	Jwt      *options.JwtOptions      `json:"jwt" mapstructure:"jwt"`           // jwt
	Redis    *options.RedisOptions    `json:"redis" mapstruct:"redis"`          // redis
	Sms      *options.SmsOptions      `json:"sms" mapstructure:"sms"`           // 短信
}

// Flags implements app.CliOptions interface.Add flags to the specified FlagSet object.
func (c *Config) Flags() (fss cliflag.NamedFlagSets) {
	// fss.FlagSet("logs") -> 创建一个FlagSet对象，命名为logs，做为专属的 logs 传递给 Config.Log
	c.Log.AddFlags(fss.FlagSet("logs"))
	c.Server.AddFlags(fss.FlagSet("service"))
	c.Registry.AddFlags(fss.FlagSet("registry"))
	c.Jwt.AddFlags(fss.FlagSet("jwt"))
	c.Redis.AddFlags(fss.FlagSet("redis"))
	c.Sms.AddFlags(fss.FlagSet("sms"))
	return fss
}

// Validate 将配置中的所有校验子逻辑 注册到当前实例的校验中
func (c *Config) Validate() (errors []error) {
	// 将 Log 中的校验，注册到 user 服务的，校验逻辑中
	errors = append(errors, c.Log.Validate()...)
	errors = append(errors, c.Server.Validate()...)
	errors = append(errors, c.Registry.Validate()...)
	errors = append(errors, c.Jwt.Validate()...)
	errors = append(errors, c.Redis.Validate()...)
	errors = append(errors, c.Sms.Validate()...)
	return
}
