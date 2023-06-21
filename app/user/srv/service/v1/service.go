package v1

import v1 "github.com/coderi421/goshop/app/user/srv/data/v1"

type ServiceFactory interface {
	User() UserSrv
}

type service struct {
	data v1.DataFactory
}

func (s *service) User() UserSrv {
	return newUserService(s.data)
}

var _ ServiceFactory = &service{}

func NewService(store v1.DataFactory) *service {
	return &service{
		data: store,
	}
}
