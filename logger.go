package logger

import (
	"fmt"
	"io"
	"log/slog"
	"sync"
)

var (
	log *slog.Logger
	mu  sync.Mutex
)

type LogCfg struct {
	Env    string
	Writer io.Writer
}
type LOG_ENV int

const (
	PROD_ENV LOG_ENV = iota
	DEV_ENV
)

func (l LOG_ENV) String() string {

	switch l {
	case PROD_ENV:
		return "Production"
	case DEV_ENV:
		return "Development"
	}
	return "unknown env"

}

// Initialize the logger instance, and define handler as per cfg.
func Init(cfg LogCfg) error {
	mu.Lock()
	defer mu.Unlock()
	var err error
	var handler slog.Handler

	if isInitialized() {
		return nil
	}

	switch cfg.Env {
	case PROD_ENV.String():
		handler = slog.NewJSONHandler(cfg.Writer, nil)
	case DEV_ENV.String():
		handler = NewDevHandler(cfg.Writer, LOG_DEBUG)
	default:
		err = fmt.Errorf("unsupported logger environment")
		return err
	}

	log = slog.New(handler)
	return nil

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
