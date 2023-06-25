package sms

import (
	"github.com/coderi421/gframework/pkg/common/core"
	"github.com/coderi421/gframework/pkg/errors"
	"github.com/coderi421/gframework/pkg/storage"
	"github.com/coderi421/goshop/app/pkg/code"
	translator "github.com/coderi421/goshop/app/pkg/translator/gin"
	"github.com/coderi421/goshop/app/shop/custom/internal/service"
	v1 "github.com/coderi421/goshop/app/shop/custom/internal/service/sms/v1"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"time"
)

type SendSmsForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号码格式有规范可寻， 自定义validator
	Type   uint   `form:"type" json:"type" binding:"required,oneof=1 2"`
	//1. 注册发送短信验证码和动态验证码登录发送验证码
}
type smsController struct {
	//srv   sms.SmsSrv
	srv   service.ServiceFactory
	trans ut.Translator
}

func NewSmsController(srv service.ServiceFactory, trans ut.Translator) *smsController {
	return &smsController{srv, trans}
}

func (sc *smsController) SendSms(c *gin.Context) {
	sendSmsForm := SendSmsForm{}
	if err := c.ShouldBind(&sendSmsForm); err != nil {
		translator.GinHandleValidatorError(c, err, sc.trans)
	}
	smsCode := v1.GenerateSmsCode(6)
	err := sc.srv.Base().SendSms(c, sendSmsForm.Mobile, "SMS_181850725", "{\"code\":"+smsCode+"}")
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrSmsSend, err.Error()), nil)
		return
	}
	//将验证码保存起来
	rstore := storage.RedisCluster{}
	err = rstore.SetKey(c, sendSmsForm.Mobile, smsCode, 5*time.Minute)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrSmsSend, err.Error()), nil)
		return
	}
	core.WriteResponse(c, nil, nil)
}
