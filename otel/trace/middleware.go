package ginoteltrace

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

func Middleware(opts ...Option) gin.HandlerFunc {
	o := defaultClientOptions()
	o.apply(opts...)
	return func(c *gin.Context) {
		ctx := o.propagator.Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))
		spanName := c.Request.Method
		ctx, span := o.tracer.Start(
			ctx,
			spanName,
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(
				semconv.RPCSystemKey.String("http"),
				semconv.HTTPMethodKey.String(c.Request.Method),
				semconv.HTTPURLKey.String(c.Request.URL.String()),
			),
		)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
		span.SetAttributes(semconv.HTTPStatusCodeKey.Int(c.Writer.Status()))
		span.End()
	}
}
