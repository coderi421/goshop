package user

import (
	"github.com/coderi421/gframework/pkg/common/core"
	"github.com/coderi421/gframework/pkg/errors"
	"github.com/coderi421/goshop/app/pkg/code"
	"github.com/coderi421/goshop/app/pkg/translator/gin"
	"github.com/gin-gonic/gin"
)

type PassWordLoginForm struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号码格式有规范可寻
	PassWord  string `form:"password" json:"password" binding:"required,min=3,max=10"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}

// Login
//
//	@Description: 登录
//	@receiver us
//	@param ctx
func (us *userServer) Login(ctx *gin.Context) {
	//表单验证
	passwordLoginForm := PassWordLoginForm{}
	if err := ctx.ShouldBind(&passwordLoginForm); err != nil {
		translator.GinHandleValidatorError(ctx, err, us.trans)
		return
	}
	//验证码验证
	if !store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, true) {
		core.WriteResponse(ctx, errors.WithCode(code.ErrSmsIncorrect, "验证码错误"), nil)
		return
	}
	//登录验证
	userDTO, err := us.srv.User().MobileLogin(ctx, passwordLoginForm.Mobile, passwordLoginForm.PassWord)
	if err != nil {
		core.WriteResponse(ctx, errors.WithCode(code.ErrUserPasswordIncorrect, "登陆失败"), nil)
		return
	}

	core.WriteResponse(ctx, nil, gin.H{
		"id":          userDTO.ID,
		"nick_name":   userDTO.NickName,
		"token":       userDTO.Token,
		"expiresd_at": userDTO.ExpiresAt,
	})
}
