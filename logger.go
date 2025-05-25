package logger

import (
	"log/slog"
	"os"
)

var log *slog.Logger

type LogCfg struct {
	env string
}

// To implement a logger, we want to create a handler, and then create logger using New.
// For dev, we will create our own handler.
// For this, 3 methods are important.
// Enabled -> To log only till mentioned level.
// Handle -> Actually format the log.
// WithAttrs()
// WithGroup()

func Init(cfg LogCfg) {
	var handler slog.Handler
	if cfg.env == "prod" {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		handler = NewDevHandler(os.Stdout, LOG_DEBUG)
	}

	log = slog.New(handler)

}

func Info(msg string, args ...any) {
	log.Info(msg, args...)
}

func Debug(msg string, args ...any) {
	log.Debug(msg, args...)
}
func Warn(msg string, args ...any) {
	log.Warn(msg, args...)
}
func Error(msg string, args ...any) {
	log.Error(msg, args...)
}
