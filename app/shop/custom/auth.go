package custom

import (
	"github.com/coderi421/gframework/gmicro/server/restserver/middlewares"
	"github.com/coderi421/gframework/gmicro/server/restserver/middlewares/auth"
	"github.com/coderi421/goshop/app/pkg/options"
	"github.com/gin-gonic/gin"

	ginjwt "github.com/appleboy/gin-jwt/v2"
)

// 可换其他认证
func newJWTAuth(opts *options.JwtOptions) middlewares.AuthStrategy {
	gjwt, _ := ginjwt.New(&ginjwt.GinJWTMiddleware{
		Realm:            opts.Realm,
		SigningAlgorithm: "HS256",
		Key:              []byte(opts.Key),
		Timeout:          opts.Timeout,
		MaxRefresh:       opts.MaxRefresh,
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(code, nil)
		},
		//可以向gin.context中放入数据
		IdentityHandler: claimHandlerFun,

		IdentityKey: middlewares.KeyUserID,
		//还可以在哪找token
		TokenLookup: "header: Authorization:, query: token, cookie: jwt",
		//前端首字符改为Bearer
		TokenHeadName: "Bearer",
	})
	return auth.NewJWTStrategy(*gjwt)
}

func claimHandlerFun(c *gin.Context) interface{} {
	//取touken反解
	claims := ginjwt.ExtractClaims(c)
	c.Set(middlewares.KeyUserID, claims[middlewares.KeyUserID])
	return claims[ginjwt.IdentityKey]
}
