package user

import (
	"github.com/coderi421/gframework/gmicro/server/restserver/middlewares"
	"github.com/coderi421/gframework/pkg/common/core"
	jtime "github.com/coderi421/gframework/pkg/common/time"
	"github.com/coderi421/goshop/app/pkg/translator/gin"
	"github.com/gin-gonic/gin"
	"time"
)

type UpdateUserForm struct {
	Name     string `form:"name" json:"name" binding:"required,min=3,max=10"`
	Gender   string `form:"gender" json:"gender" binding:"required,oneof=female male"`
	Birthday string `form:"birthday" json:"birthday" binding:"required,datetime=2006-01-02"`
}

func (us *userServer) UpdateUser(ctx *gin.Context) {
	updateForm := UpdateUserForm{}
	if err := ctx.ShouldBind(&updateForm); err != nil {
		translator.GinHandleValidatorError(ctx, err, us.trans)
		return
	}
	userID, _ := ctx.Get(middlewares.KeyUserID)
	userIDInt := int32(userID.(float64))
	userDTO, err := us.srv.User().Get(ctx, userIDInt)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	//将前端传递过来的日期格式转换成int
	loc, _ := time.LoadLocation("Local") //local的L必须大写
	birthDay, _ := time.ParseInLocation("2006-01-02", updateForm.Birthday, loc)
	userDTO.NickName = updateForm.Name
	userDTO.Birthday = jtime.Time{Time: birthDay}.Time
	userDTO.Gender = updateForm.Gender
	err = us.srv.User().Update(ctx, userDTO)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}
