package gincontext

import (
	"github.com/gin-gonic/gin"
)

func Middleware(opts ...Option) gin.HandlerFunc {
	o := defaultOptions()
	o.apply(opts...)
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = o.contextFunc(ctx)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
