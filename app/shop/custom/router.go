package custom

import (
	"github.com/coderi421/gframework/gmicro/server/restserver"
	"github.com/coderi421/goshop/app/shop/custom/internal/controller"
)

func initRouter(g *restserver.Server) {
	v1 := g.Group("/v1")
	ugroup := v1.Group("/user")

	{
		ucontroller := controller.NewUserController()
		ugroup.GET("list", ucontroller.List)
	}
}
