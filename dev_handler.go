package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"strings"
)

type LogLevel int

const (
	LOG_INFO LogLevel = iota
	LOG_DEBUG
	LOG_WARN
	LOG_ERROR
)

type DevHandler struct {
	level  slog.Level
	writer io.Writer
}

var h slog.Handler

func NewDevHandler(w io.Writer, level LogLevel) *DevHandler {
	var logLevel slog.Level
	switch level {
	case LOG_INFO:
		logLevel = slog.LevelInfo
	case LOG_DEBUG:
		logLevel = slog.LevelDebug
	case LOG_WARN:
		logLevel = slog.LevelWarn
	case LOG_ERROR:
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo

	}
	return &DevHandler{
		level:  logLevel,
		writer: w,
	}

}

func (d *DevHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= d.level

}

func (d *DevHandler) Handle(_ context.Context, r slog.Record) error {
	var msg string
	ts := r.Time.Format("12-06-2025 14:00:00")
	level := strings.ToUpper(r.Level.String())

	if r.Level == slog.LevelError {
		_, file, line, _ := runtime.Caller(4)
		msg = fmt.Sprintf("%s [%s] %s - %s:%s", ts, level, r.Message, file, line)
	} else {
		msg = fmt.Sprintf("%s [%s] %s", ts, level, r.Message)
	}

	_, err := fmt.Fprintf(d.writer, "%s %s", ts, msg)
	if err != nil {
		return err
	}
	return nil
}

func (d *DevHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return d
}

func (d *DevHandler) WithGroup(name string) slog.Handler {
	return d
}
