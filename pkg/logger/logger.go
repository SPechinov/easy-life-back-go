package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"runtime"
)

type contextKey string

const logContextKey contextKey = "logrus-context"

func Get(ctx context.Context) *logrus.Entry {
	if ctx == nil {
		return logrus.WithFields(logrus.Fields{})
	}

	if entry, ok := ctx.Value(logContextKey).(*logrus.Entry); ok {
		return entry
	}

	return logrus.WithFields(logrus.Fields{})
}

func Set(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, logContextKey, logger)
}

func Trace(ctx context.Context, args ...any) {
	l := Get(ctx)
	l.Trace(args)
}

func Debug(ctx context.Context, args ...any) {
	l := Get(ctx)
	l.Debug(args)
}

func Print(ctx context.Context, args ...any) {
	l := Get(ctx)
	l.Print(args)
}

func Info(ctx context.Context, args ...any) {
	l := Get(ctx)
	l.Info(args)
}

func Warn(ctx context.Context, args ...any) {
	l := Get(ctx)
	l.Warn(args)
}

func Warning(ctx context.Context, args ...any) {
	l := Get(ctx)
	l.Warning(args)
}

func Error(ctx context.Context, args ...any) {
	l := Get(ctx)

	_, file, line, ok := runtime.Caller(1)
	if ok {
		l = l.WithFields(logrus.Fields{
			"File": file,
			"Line": line,
		})
	}

	l.Error(args)
}

func Fatal(ctx context.Context, args ...any) {
	l := Get(ctx)
	l.Fatal(args)
}

func Panic(ctx context.Context, args ...any) {
	l := Get(ctx)
	l.Panic(args)
}
