package xk6slog

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

type LoggerOpts struct {
	Output   string
	Format   string
	Level    string
	Filepath string
}

func (opts LoggerOpts) CreateSLogHandler() slog.Handler {
	output := opts.getOutputWriter()
	handlerOpts := &slog.HandlerOptions{Level: opts.getLevel()}

	switch strings.ToLower(opts.Format) {
	case "text":
		return slog.NewTextHandler(output, handlerOpts)
	case "json":
		return slog.NewJSONHandler(output, handlerOpts)
	default:
		panic(fmt.Errorf("undefined log format: %s", opts.Format))
	}
}

func (opts LoggerOpts) getLevel() slog.Level {
	switch strings.ToUpper(opts.Level) {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN", "WARNING":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		panic(fmt.Errorf("undefined log level: %s", opts.Level))
	}
}

func (opts LoggerOpts) getOutputWriter() io.Writer {
	switch strings.ToLower(opts.Output) {
	case "console":
		return os.Stdout
	case "file":
		file, err := os.OpenFile(opts.Filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(fmt.Errorf("error opening log file: %w", err))
		}
		return file
	default:
		panic(fmt.Errorf("undefined output type: %s", opts.Output))
	}
}
