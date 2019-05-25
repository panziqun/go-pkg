package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-pkg/conf"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

type SmsResponse struct {
	Message   string
	RequestID string
	Code      string
	BizId     string
}

func SendCode(tel, code string) error {
	smsConf := conf.Section("aliyunSMS")
	client, err := sdk.NewClientWithAccessKey("default", smsConf["AppKey"], smsConf["AppSecret"])
	if err != nil {
		panic(err)
	}

	param := map[string]string{
		"code": code,
	}
	templateParam, _ := json.Marshal(param)

	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["PhoneNumbers"] = tel
	request.QueryParams["SignName"] = smsConf["SignName"]
	request.QueryParams["TemplateCode"] = smsConf["TemplateCode"]
	request.QueryParams["TemplateParam"] = string(templateParam)

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	smsR := SmsResponse{}
	json.Unmarshal([]byte(response.GetHttpContentString()), &smsR)
	if smsR.Code != "OK" {
		fmt.Println(response.GetHttpContentString())
		return errors.New(smsR.Message)
	}

	return nil
}
