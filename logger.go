package xk6slog

import (
	"context"
	"log/slog"
)

type Logger struct {
	logger *slog.Logger
}

func NewLogger(opts LoggerOpts) *Logger {
	return &Logger{logger: slog.New(opts.CreateSLogHandler())}
}

func (l *Logger) Debug(msg string, fields map[string]interface{}) {
	l.log(msg, slog.LevelDebug, fields)
}

func (l *Logger) Info(msg string, fields map[string]interface{}) {
	l.log(msg, slog.LevelInfo, fields)
}

func (l *Logger) Warn(msg string, fields map[string]interface{}) {
	l.log(msg, slog.LevelWarn, fields)
}

func (l *Logger) Error(msg string, fields map[string]interface{}) {
	l.log(msg, slog.LevelError, fields)
}

func (l *Logger) log(msg string, level slog.Level, fields map[string]interface{}) {
	attrs := make([]slog.Attr, 0, len(fields))

	for k, v := range fields {
		attrs = append(attrs, slog.Any(k, v))
	}

	l.logger.LogAttrs(context.Background(), level, msg, attrs...)
}
