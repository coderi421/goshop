package user

import (
	"github.com/coderi421/goshop/app/shop/custom/internal/service"
	ut "github.com/go-playground/universal-translator"
)

type userServer struct {
	trans ut.Translator
	srv   service.ServiceFactory
}

func NewUserController(trans ut.Translator, srv service.ServiceFactory) *userServer {
	return &userServer{trans, srv}
}
