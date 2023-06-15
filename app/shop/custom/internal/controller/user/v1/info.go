package user

import (
	"github.com/coderi421/gframework/gmicro/server/restserver/middlewares"
	"github.com/coderi421/gframework/pkg/common/core"
	"github.com/coderi421/gframework/pkg/errors"
	"github.com/coderi421/goshop/app/pkg/code"
	"github.com/gin-gonic/gin"
)

// GetUserDetail
//
//	@Description: 获取用户详情
//	@receiver us
//	@param ctx
func (us *userServer) GetUserDetail(ctx *gin.Context) {
	userID, exists := ctx.Get(middlewares.KeyUserID)
	if !exists {
		core.WriteResponse(ctx, errors.WithCode(code.ErrInvalidAuthHeader, "未登录"), nil)
		return
	}

	//最好debug看一下是什么类型  再进行断言
	userDTO, err := us.srv.User().Get(ctx, int32(userID.(float64)))
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	core.WriteResponse(ctx, nil, gin.H{
		"name":     userDTO.NickName,
		"birthday": userDTO.Birthday.Format("2006-01-02"),
		"gender":   userDTO.Gender,
		"mobile":   userDTO.Mobile,
	})
}
