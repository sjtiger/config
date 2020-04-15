package log

import (
	"context"
	"io"

	"github.com/sirupsen/logrus"
)

type contextKey int

// Logger for app
type Logger interface {
	logrus.FieldLogger
	WriterLevel(logrus.Level) *io.PipeWriter
}

const (
	loggerKey contextKey = iota
)

var (
	mainLogger Logger
)

func init() {
	mainLogger = logrus.StandardLogger()
}

// SetLogger set the logger
func SetLogger(l Logger) {
	mainLogger = l
}

// SetOutput set logger
func SetOutput(o io.Writer) {
	logrus.SetOutput(o)
}

// SetLevel set logger level
func SetLevel(l logrus.Level) {
	logrus.SetLevel(l)
}

// GetLevel return logger level
func GetLevel() logrus.Level {
	return logrus.GetLevel()
}

// WithoutContext Gets the main logger
func WithoutContext() Logger {
	return mainLogger
}

// Str adds a string field
func Str(key, value string) func(logrus.Fields) {
	return func(fields logrus.Fields) {
		fields[key] = value
	}
}

// With Adds fields
func With(ctx context.Context, opts ...func(logrus.Fields)) context.Context {
	logger := FromContext(ctx)

	fields := make(logrus.Fields)
	for _, opt := range opts {
		opt(fields)
	}
	logger = logger.WithFields(fields)

	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext return logger from context
func FromContext(ctx context.Context) Logger {
	if ctx == nil {
		panic("nil context")
	}

	logger, ok := ctx.Value(loggerKey).(Logger)
	if !ok {
		logger = mainLogger
	}

	return logger
}
