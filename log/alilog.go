package log

import (
	cf "app/conf"
	"encoding/json"
	"fmt"
	"time"

	"github.com/laughmaker/go-pkg/conf"
	"github.com/laughmaker/go-pkg/mail"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

func Put(topic string, c *gin.Context) {
	PutLog(topic, c, "", "")
}

func PutLog(topic string, c *gin.Context, err string, stack string) {
	contents := []*sls.LogContent{}

	params, _ := json.Marshal(c.Params)
	basic, _ := json.Marshal(map[string]string{
		"method":     c.Request.Method,
		"params":     string(params),
		"host":       c.Request.Host,
		"clientIp":   c.ClientIP(),
		"remoteAddr": c.Request.RemoteAddr,
		"requestURI": c.Request.RequestURI,
	})
	content := &sls.LogContent{
		Key:   proto.String("basic"),
		Value: proto.String(string(basic)),
	}
	contents = append(contents, content)

	header, _ := json.Marshal(c.Request.Header)
	content = &sls.LogContent{
		Key:   proto.String("header"),
		Value: proto.String(string(header)),
	}
	contents = append(contents, content)

	requestBody, _ := c.Get(cf.ReqBodyKey)
	content = &sls.LogContent{
		Key:   proto.String("request"),
		Value: proto.String(requestBody.(string)),
	}
	contents = append(contents, content)

	respData, _ := c.Get(cf.RespDataKey)
	rd, _ := json.Marshal(respData)
	content = &sls.LogContent{
		Key:   proto.String("response"),
		Value: proto.String(string(rd)),
	}
	contents = append(contents, content)

	if err != "" {
		content = &sls.LogContent{
			Key:   proto.String("error"),
			Value: proto.String(err),
		}
		contents = append(contents, content)
	}

	if stack != "" {
		content = &sls.LogContent{
			Key:   proto.String("stack"),
			Value: proto.String(stack),
		}
		contents = append(contents, content)
	}

	logs := []*sls.Log{}
	log := &sls.Log{
		Time:     proto.Uint32(uint32(time.Now().Unix())),
		Contents: contents,
	}
	logs = append(logs, log)

	logGroup := &sls.LogGroup{
		Topic:  proto.String(topic),
		Source: proto.String(c.ClientIP()),
		Logs:   logs,
	}

	logConf := conf.Section("aliyunLog")
	client := &sls.Client{
		Endpoint:        logConf["Endpoint"],
		AccessKeyID:     logConf["AccessKeyID"],
		AccessKeySecret: logConf["AccessKeySecret"],
	}
	e := client.PutLogs(logConf["Project"], logConf["Logstore"], logGroup)
	if e != nil {
		mail.Send(conf.App.LogMail, "推送日志失败！", fmt.Sprintf("%v", e), "")
		fmt.Println(e)
	}
}
