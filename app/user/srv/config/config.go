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
		Log:          log.NewOptions(),
		Server:       options.NewServerOptions(),
		Registry:     options.NewRegistryOptions(),
		Telemetry:    options.NewTelemetryOptions(),
		MySQLOptions: options.NewMySQLOptions(),
	}
}

type Config struct {
	Log *log.Options `json:"log" mapstructure:"log"`
	// 服务发现
	Server *options.ServerOptions `json:"server" mapstructure:"server"`
	// 注册中心
	Registry *options.RegistryOptions `json:"registry" mapstructure:"registry"`
	//	链路追踪
	Telemetry    *options.TelemetryOptions `json:"telemetry" mapstructure:"telemetry"`
	MySQLOptions *options.MySQLOptions     `json:"mysql" mapstructure:"mysql"`
}

// Flags implements app.CliOptions interface.Add flags to the specified FlagSet object.
func (c *Config) Flags() (fss cliflag.NamedFlagSets) {
	// fss.FlagSet("logs") -> 创建一个FlagSet对象，命名为logs，做为专属的 logs 传递给 Config.Log
	c.Log.AddFlags(fss.FlagSet("logs"))
	c.Server.AddFlags(fss.FlagSet("server"))
	c.Registry.AddFlags(fss.FlagSet("registry"))
	c.Telemetry.AddFlags(fss.FlagSet("telemetry"))
	c.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	return fss
}

// Validate 将配置中的所有校验子逻辑 注册到当前实例的校验中
func (c *Config) Validate() (errors []error) {
	// 将 Log 中的校验，注册到 admin 服务的，校验逻辑中
	errors = append(errors, c.Log.Validate()...)
	errors = append(errors, c.Server.Validate()...)
	errors = append(errors, c.Registry.Validate()...)
	errors = append(errors, c.Telemetry.Validate()...)
	errors = append(errors, c.MySQLOptions.Validate()...)
	return
}
