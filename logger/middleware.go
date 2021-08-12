package ginlogger

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func Middleware(opts ...Option) gin.HandlerFunc {
	o := defaultOptions()
	o.apply(opts...)
	return func(c *gin.Context) {
		if o.loggerFactory == nil {
			c.Next()
			return
		}
		logger := o.loggerFactory(c.Request.Context())
		// Start timer
		startTime := time.Now()
		c.Next()
		builder := NewFieldBuilder().
			System().
			StartTime(startTime).
			Deadline(c.Request.Context()).
			Method(c.Request.Method).
			URI(c.Request.RequestURI).
			Proto(c.Request.Proto).
			Host(c.Request.Host).
			RemoteAddress(c.Request.RemoteAddr).
			//Header(c.Request.Header).
			Status(c.Writer.Status()).
			Error(errors.New(c.Errors.String())).
			Latency(time.Since(startTime))
		logger.Log(builder.Build())
	}
}
