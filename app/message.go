package app

var Message = map[int]string{
	SUCCESS:        "ok",
	FAILED:         "failed",
	INVALID_PARAMS: "请求参数错误",
	INVALID_TOKEN:  "无效的令牌",
	TOKEN_EXPIRED:  "无效的令牌",
	TOKEN_EMPTY:    "空令牌",
}

func GetMessage(code int) string {
	message, ok := Message[code]
	if ok {
		return message
	}

	return Message[FAILED]
}
