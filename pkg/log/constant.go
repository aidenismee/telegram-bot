package log

import (
	"go.uber.org/zap/zapcore"
)

const (
	FieldCorrelationID = "cid"
	FieldError         = "error"
	FieldLogHeader     = "log_header"
)

const (
	LevelDebug = "DEBUG"
	LevelInfo  = "INFO"
	LevelWarn  = "WARN"
	LevelError = "ERROR"
	LevelPanic = "PANIC"
	LevelFatal = "FATAL"
)

var Levels = map[string]zapcore.Level{
	LevelDebug: zapcore.DebugLevel,
	LevelInfo:  zapcore.InfoLevel,
	LevelWarn:  zapcore.WarnLevel,
	LevelError: zapcore.ErrorLevel,
	LevelPanic: zapcore.PanicLevel,
	LevelFatal: zapcore.FatalLevel,
}
