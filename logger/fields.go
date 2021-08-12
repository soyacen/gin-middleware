package ginlogger

import (
	"context"
	"net/http"
	"time"
)

var (
	ServerField = map[string]interface{}{}
	ClientField = map[string]interface{}{"system": "http.client"}
)

type FieldBuilder struct {
	fields map[string]interface{}
}

func NewFieldBuilder() *FieldBuilder {
	return &FieldBuilder{
		fields: make(map[string]interface{}),
	}
}

func (f *FieldBuilder) System() *FieldBuilder {
	f.fields["system"] = "http.server"
	return f
}

func (f *FieldBuilder) StartTime(startTime time.Time) *FieldBuilder {
	f.fields["http.start_time"] = startTime.Format(time.RFC3339)
	return f
}

func (f *FieldBuilder) Deadline(ctx context.Context) *FieldBuilder {
	if d, ok := ctx.Deadline(); ok {
		f.fields["http.deadline"] = d.Format(time.RFC3339)
	}
	return f
}

func (f *FieldBuilder) Latency(duration time.Duration) *FieldBuilder {
	f.fields["http.latency"] = duration.String()
	return f
}

func (f *FieldBuilder) Method(method string) *FieldBuilder {
	f.fields["http.method"] = method
	return f
}

func (f *FieldBuilder) URI(uri string) *FieldBuilder {
	f.fields["http.uri"] = uri
	return f
}

func (f *FieldBuilder) Proto(proto string) *FieldBuilder {
	f.fields["http.proto"] = proto
	return f
}

func (f *FieldBuilder) Host(host string) *FieldBuilder {
	f.fields["http.host"] = host
	return f
}

func (f *FieldBuilder) RemoteAddress(address string) *FieldBuilder {
	f.fields["http.remote.address"] = address
	return f
}

func (f *FieldBuilder) Header(header http.Header) *FieldBuilder {
	for key, val := range header {
		f.fields["http.header."+key] = val
	}
	return f
}

func (f *FieldBuilder) Status(status int) *FieldBuilder {
	f.fields["http.response.status"] = status
	return f
}

func (f *FieldBuilder) Error(err error) *FieldBuilder {
	if err == nil {
		return f
	}
	f.fields["error"] = err
	return f
}

func (f *FieldBuilder) Build() map[string]interface{} {
	return f.fields
}
