package v1

import (
	"context"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/coderi421/goshop/app/pkg/options"
	"math/rand"
	"strings"
	"time"
)

type SmsSrv interface {
	//
	// SendSms
	//  @Description: 发送短信验证码
	//  @param ctx
	//  @param movile: 手机号码
	//  @param tpc: teamplate code 消息模板编号
	//  @param tp: template param 消息参数
	//  @return error
	//
	SendSms(ctx context.Context, mobile string, tpc, tp string) error
}

func GenerateSmsCode(width int) string {
	//生成width长度的短信验证码
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	randSrc := rand.NewSource(time.Now().UnixNano())
	rand.New(randSrc)
	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()

}
func (s *smsService) SendSms(ctx context.Context, mobile string, tpc, tp string) error {
	client, err := dysmsapi.NewClientWithAccessKey("cn-beijing", s.smsOpts.APIKey, s.smsOpts.APISecret)
	if err != nil {
		panic(err)
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-beijing"
	request.QueryParams["PhoneNumbers"] = mobile //手机号
	request.QueryParams["SignName"] = "灿若烬星"     //阿里云验证过的项目名 自己设置
	request.QueryParams["TemplateCode"] = tpc    //阿里云的短信模板号 自己设置
	request.QueryParams["TemplateParam"] = tp    //短信模板中的验证码内容 自己生成   之前试过直接返回，但是失败，加上code成功。
	response, err := client.ProcessCommonRequest(request)
	fmt.Print(client.DoAction(request, response))
	if err != nil {
		return err
	}
	return nil
}

type smsService struct {
	smsOpts *options.SmsOptions
}

func NewSms(smsOpts *options.SmsOptions) SmsSrv {
	return &smsService{smsOpts: smsOpts}
}
