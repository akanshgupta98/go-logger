package logger

import (
	"log/slog"
	"os"
)

var log *slog.Logger

type LogCfg struct {
	Env string
}

// Initialize the logger instance, and define handler as per cfg.
func Init(cfg LogCfg) {
	var handler slog.Handler
	if cfg.Env == "prod" {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		handler = NewDevHandler(os.Stdout, LOG_DEBUG)
	}

	log = slog.New(handler)

}

// Check whether logger is initialized or not.
func isInitialized() bool {
	return log != nil
}

// Log Info level log
func Info(msg string, args ...any) {
	if !isInitialized() {
		panic("logger called before Initialize")
	}
	log.Info(msg, args...)
}

// Log Debug level log
func Debug(msg string, args ...any) {
	if !isInitialized() {
		panic("logger called before Initialize")
	}
	log.Debug(msg, args...)
}

// Log Warning level log
func Warn(msg string, args ...any) {
	if !isInitialized() {
		panic("logger called before Initialize")
	}
	log.Warn(msg, args...)
}

// Log Error level log
func Error(msg string, args ...any) {
	if !isInitialized() {
		panic("logger called before Initialize")
	}
	log.Error(msg, args...)
}
