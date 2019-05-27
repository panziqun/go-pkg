package middleware

import (
	"app/conf"

	"github.com/laughmaker/go-pkg/resp"
	"github.com/laughmaker/go-pkg/util"

	"github.com/gin-gonic/gin"
)

var UserId int

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("token")
		if token == "" {
			abort(c, conf.TokenEmpty)
			return
		}
		var err error
		UserId, err = util.ParseToken(token)
		if err != nil {
			abort(c, conf.TokenInvalid)
			return
		}

		c.Next()
	}
}

func abort(c *gin.Context, code int) {
	r := resp.Resp{C: c}
	r.Failure(code)
	c.Abort()
}
