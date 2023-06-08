package service

import user "github.com/coderi421/goshop/app/shop/custom/internal/service/user/v1"

type ServiceFactory interface {
	User() user.UserSrv
}
