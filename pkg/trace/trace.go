package trace

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
)

const (
	TracerName = "notes/notesBE"
)

type Span struct {
	name string
	span oteltrace.Span
}

// NewSpan 新建一个 span
func NewSpan(ctx context.Context, name string) (context.Context, *Span) {
	var span oteltrace.Span
	// 向trace context中添加trace
	tracer := otel.Tracer(TracerName)
	ctx, span = tracer.Start(ctx, name)

	return ctx, &Span{
		name: name,
		span: span,
	}
}

// TraceID 获取 traceid
func (s *Span) TraceID() string {
	if s.span == nil {
		return ""
	}

	traceID := s.span.SpanContext().TraceID()
	if !traceID.IsValid() {
		return ""
	}

	return traceID.String()
}

// Set attribute 打点
func (s *Span) Set(key string, value any) {
	if s.span == nil {
		return
	}
	var attr attribute.KeyValue
	switch v := value.(type) {
	case bool:
		attr = attribute.Bool(key, v)
	case int:
		attr = attribute.Int(key, v)
	case int64:
		attr = attribute.Int64(key, v)
	case []int64:
		attr = attribute.Int64Slice(key, v)
	case float64:
		attr = attribute.Float64(key, v)
	case []float64:
		attr = attribute.Float64Slice(key, v)
	case []byte:
		attr = attribute.String(key, string(v))
	case string:
		attr = attribute.String(key, v)
	case []string:
		attr = attribute.StringSlice(key, v)
	case time.Time:
		location, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			return
		}
		t := value.(time.Time)
		attr = attribute.String(key, t.In(location).Format("2006-01-02 15:04:05"))
	case time.Duration:
		attr = attribute.String(key, value.(time.Duration).String())
	default:
		attr = attribute.String(key, fmt.Sprintf("%+v", value))
	}

	s.span.SetAttributes(attr)
}

// End span end 增加错误异常判断
func (s *Span) End(errPoint *error) {
	if *errPoint != nil {
		s.span.SetStatus(codes.Error, fmt.Sprintf("%v", *errPoint))
		s.span.RecordError(*errPoint)
	}
	s.span.End()
}
