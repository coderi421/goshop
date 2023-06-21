package service

import (
	"github.com/coderi421/goshop/app/pkg/options"
	"github.com/coderi421/goshop/app/shop/custom/internal/data"
	v1 "github.com/coderi421/goshop/app/shop/custom/internal/service/sms/v1"
	user "github.com/coderi421/goshop/app/shop/custom/internal/service/user/v1"
)

// ServiceFactory 工厂模式
type ServiceFactory interface {
	User() user.UserSrv
	Base() v1.SmsSrv
}

// 持久化
type service struct {
	data    data.DataFactory
	jwtOpts *options.JwtOptions
	smsOpts *options.SmsOptions
}

func (s service) User() user.UserSrv {
	return user.NewUser(s.data, s.jwtOpts)
}

func (s service) Base() v1.SmsSrv {
	return v1.NewSms(s.smsOpts)
}

var _ ServiceFactory = &service{}
