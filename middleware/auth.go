package middleware

import (
	"github.com/laughmaker/go-pkg/resp"
	"github.com/laughmaker/go-pkg/util"

	"github.com/gin-gonic/gin"
)

// Auth auth middleware
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("token")
		if token == "" {
			abort(c, resp.TokenEmpty)
			return
		}
		var err error
		_, err = util.ParseToken(token)
		if err != nil {
			abort(c, resp.TokenInvalid)
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
