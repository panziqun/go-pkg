package resp

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

const (
	// Success access successful
	Success = 2000

	// Failure access failed
	Failure = 2001

	// NetworkError network access failed
	NetworkError = 4000

	// TokenInvalid token invalid
	TokenInvalid = 4001

	// TokenExpired token has expired
	TokenExpired = 4002

	// TokenEmpty token cannot be empty
	TokenEmpty = 4003

	// DataNotExist data not exist
	DataNotExist = 4010

	// DataExist data has exist
	DataExist = 4011

	// UsernameOrPasswordError data has exist
	UsernameOrPasswordError = 4012
)

// CodeText code message struct
var codeText = map[int]string{
	Success: "请求成功！",
	Failure: "请求失败",

	NetworkError: "网络请求失败！",

	TokenInvalid: "无效的令牌！",
	TokenExpired: "过期的令牌！",
	TokenEmpty:   "空令牌！",

	DataNotExist:            "数据不存在！",
	DataExist:               "数据已存在！",
	UsernameOrPasswordError: "用户名或密码错误！",
}

var file *ini.File

// Setup load resp config
func Setup() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err = ini.Load(wd + "/conf/code.ini")
	if err != nil {
		fmt.Printf("fail to parse 'code.ini': %v", err)
		panic(err)
	}

	respCode := file.Section("respCode").KeysHash()
}

// CodeText access code message
func getCodeText(code int) string {
	message, ok := codeText[code]
	if ok {
		return message
	}

	respCode := file.Section("respCode").KeysHash()
	message = respCode[fmt.Sprintf("%d", code)]
	if message != "" {
		return message
	}

	return codeText[Failure]
}

// RespDataKey 存储每次返回数据到Context中
const RespDataKey = "RESP_DATA_KEY"

// ReqBodyKey 日志body备份
const ReqBodyKey = "REQ_BODY_KEY"

// DefaultDeletedAt default delete time
const DefaultDeletedAt = "2019-05-01 00:00:00"
