package user

import (
	"github.com/coderi421/gframework/pkg/common/core"
	translator "github.com/coderi421/goshop/app/pkg/translator/gin"
	v1 "github.com/coderi421/goshop/app/shop/custom/internal/service/user/v1"
	"github.com/gin-gonic/gin"
)

type RegisterForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号码格式有规范可寻， 自定义validator
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Code     string `form:"code" json:"code" binding:"required,min=6,max=6"`
}

// Register 注册
//
//	@Description:
//	@receiver us
//	@param ctx
func (us *userServer) Register(ctx *gin.Context) {
	regForm := RegisterForm{}
	if err := ctx.ShouldBind(&regForm); err != nil {
		translator.GinHandleValidatorError(ctx, err, us.trans)
		return
	}
	req := &v1.UserDTO{
		Mobile:   regForm.Mobile,
		PassWord: regForm.PassWord,
	}
	userDTO, err := us.srv.User().Register(ctx, req, regForm.Code)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, gin.H{
		"id":         userDTO.ID,
		"nick_name":  userDTO.NickName,
		"token":      userDTO.Token,
		"expired_at": userDTO.ExpiresAt,
	})

}
