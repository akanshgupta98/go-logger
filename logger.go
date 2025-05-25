package logger

import (
	"fmt"
	"io"
	"log/slog"
	"sync"
)

var (
	log  *slog.Logger
	once sync.Once
)

type LogCfg struct {
	Env    string
	Writer io.Writer
}

const (
	PROD_ENV = "Production"
	DEV_ENV  = "Development"
)

// Initialize the logger instance, and define handler as per cfg.
func Init(cfg LogCfg) {

	once.Do(func() {
		var handler slog.Handler

		if cfg.Env == PROD_ENV {
			handler = slog.NewJSONHandler(cfg.Writer, nil)
		} else {
			handler = NewDevHandler(cfg.Writer, LOG_DEBUG)
		}

		log = slog.New(handler)
	})

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

// Log Info level log
func Infof(msg string, args ...any) {
	if !isInitialized() {
		panic("logger called before Initialize")
	}

	logMsg := fmt.Sprintf(msg, args...)
	log.Info(logMsg)

}

// Log Debug level log
func Debug(msg string, args ...any) {
	if !isInitialized() {
		panic("logger called before Initialize")
	}
	log.Debug(msg, args...)
}

// Log Debug level log
func Debugf(msg string, args ...any) {
	if !isInitialized() {
		panic("logger called before Initialize")
	}

	logMsg := fmt.Sprintf(msg, args...)
	log.Debug(logMsg)

}

// Log Warning level log
func Warn(msg string, args ...any) {
	if !isInitialized() {
		panic("logger called before Initialize")
	}
	log.Warn(msg, args...)
}

// Log Warn level log
func Warnf(msg string, args ...any) {
	if !isInitialized() {
		panic("logger called before Initialize")
	}

	logMsg := fmt.Sprintf(msg, args...)
	log.Warn(logMsg)

}

// Log Error level log
func Error(msg string, args ...any) {
	if !isInitialized() {
		panic("logger called before Initialize")
	}
	log.Error(msg, args...)
}

// Log Error level log
func Errorf(msg string, args ...any) {
	if !isInitialized() {
		panic("logger called before Initialize")
	}

	logMsg := fmt.Sprintf(msg, args...)
	log.Error(logMsg)

}
