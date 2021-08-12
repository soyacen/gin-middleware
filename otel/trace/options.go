package ginoteltrace

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type RequestHandlerFunc func(span trace.Span, c *gin.Context)

type options struct {
	tracer             trace.Tracer
	propagator         propagation.TextMapPropagator
	requestHandlerFunc RequestHandlerFunc
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(o *options)

func defaultClientOptions() *options {
	return &options{
		tracer:             otel.Tracer(""),
		propagator:         otel.GetTextMapPropagator(),
		requestHandlerFunc: func(span trace.Span, c *gin.Context) {},
	}
}

func WithTracer(tracer trace.Tracer) Option {
	return func(o *options) {
		o.tracer = tracer
	}
}

func WithPropagator(propagator propagation.TextMapPropagator) Option {
	return func(o *options) {
		o.propagator = propagator
	}
}

func WithRequestHandlerFunc(requestHandlerFunc RequestHandlerFunc) Option {
	return func(o *options) {
		o.requestHandlerFunc = requestHandlerFunc
	}
}
