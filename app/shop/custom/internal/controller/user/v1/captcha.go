package user

import (
	"github.com/coderi421/gframework/pkg/common/core"
	"github.com/coderi421/gframework/pkg/errors"
	"github.com/coderi421/gframework/pkg/log"
	"github.com/coderi421/goshop/app/pkg/code"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

// GetCaptcha
//
//	@Description: 获取验证码
//	@param ctx
func GetCaptcha(ctx *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	//通过设置的driver放到store里面
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		log.Errorf("生成验证码错误,:", err.Error())
		core.WriteResponse(ctx, errors.WithCode(code.ErrSmsIncorrect, "生成验证码错误"), nil)
		return
	}
	core.WriteResponse(ctx, nil, gin.H{
		"captchaId": id,
		"picPath":   b64s,
	})
	//store.Verify(id,b64s,true)
}
