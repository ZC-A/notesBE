package log

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func withTraceID(ctx context.Context, format string, v ...any) string {
	str := fmt.Sprintf(format, v...)
	span := trace.SpanFromContext(ctx)
	traceID := span.SpanContext().TraceID()

	if traceID != [16]byte{0} {
		return fmt.Sprintf("[%s] %s", traceID, str)
	}
	return str
}

type Logger struct {
	logger *zap.Logger
}

func (l *Logger) Warnf(ctx context.Context, format string, v ...any) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.Warn(withTraceID(ctx, format, v...))
}

func (l *Logger) Infof(ctx context.Context, format string, v ...any) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.Info(withTraceID(ctx, format, v...))
}

func (l *Logger) Errorf(ctx context.Context, format string, v ...any) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.Error(withTraceID(ctx, format, v...))
}

func (l *Logger) Debugf(ctx context.Context, format string, v ...any) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.Debug(withTraceID(ctx, format, v...))
}

func (l *Logger) Panicf(ctx context.Context, format string, v ...any) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.Panic(withTraceID(ctx, format, v...))
}

func (l *Logger) Fatalf(ctx context.Context, format string, v ...any) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.Fatal(withTraceID(ctx, format, v...))
}

func Warnf(ctx context.Context, format string, v ...any) {
	DefaultLogger.Warnf(ctx, format, v...)
}

func Infof(ctx context.Context, format string, v ...any) {
	DefaultLogger.Infof(ctx, format, v...)
}

func Errorf(ctx context.Context, format string, v ...any) {
	DefaultLogger.Errorf(ctx, format, v...)
}

func Debugf(ctx context.Context, format string, v ...any) {
	DefaultLogger.Debugf(ctx, format, v...)
}

func Panicf(ctx context.Context, format string, v ...any) {
	DefaultLogger.Panicf(ctx, format, v...)
}

func Fatalf(ctx context.Context, format string, v ...any) {
	DefaultLogger.Fatalf(ctx, format, v...)
}
