package custom

import (
	"github.com/coderi421/gframework/gmicro/server/restserver"
	"github.com/coderi421/goshop/app/shop/custom/config"
	"github.com/coderi421/goshop/app/shop/custom/internal/controller/user/v1"
	"github.com/coderi421/goshop/app/shop/custom/internal/data/rpc"
	"github.com/coderi421/goshop/app/shop/custom/internal/service"
)

func initRouter(g *restserver.Server, conf *config.Config) {
	v1 := g.Group("/v1")
	ugroup := v1.Group("/user")

	data, err := rpc.GetDataFactoryOr(conf.Registry)
	if err != nil {
		panic(err)
	}

	serviceFactory := service.NewService(data, conf.Jwt, conf.Sms)
	uController := user.NewUserController(g.Translator(), serviceFactory)
	{
		ugroup.POST("pwd_login", uController.Login)
		ugroup.POST("register", uController.Register)

		jwtAuth := newJWTAuth(conf.Jwt)
		//第三方登录模式 暂不用
		//jwtStragy:=jwtAuth.(auth.JWTStrategy)
		//jwtStragy.LoginHandler()
		ugroup.GET("detail", jwtAuth.AuthFunc(), uController.GetUserDetail)
		ugroup.PATCH("update", jwtAuth.AuthFunc(), uController.UpdateUser)
	}
}
