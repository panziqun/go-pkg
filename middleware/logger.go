package middleware

import (
	"bytes"
	"io/ioutil"

	"github.com/laughmaker/go-pkg/log"
	"github.com/laughmaker/go-pkg/resp"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 捕获panic异常
		defer log.Try(c)

		// 写入访问日志
		// 把request的内容读取出来
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		// 把刚刚读出来的再写进去
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		c.Set(resp.ReqBodyKey, string(bodyBytes))
		defer log.Put("Info", c)

		c.Next()
	}
}
