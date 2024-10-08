package log

import (
	"github.com/labstack/echo/v4"
	"log"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// A Logger provides fast, leveled, structured logging. All methods are safe
// for concurrent use, along with filter policy to synthesis logging data.
type Logger struct {
	*zap.Logger
	level          zapcore.Level
	encodeTime     zapcore.TimeEncoder
	encodeDuration zapcore.DurationEncoder

	prefix        string
	isDevelopment bool
	correlationID string
}

// NewLog create a new logger instance with default Zap logging production config
// and a logging scope based on the given name parameter. Custom logging option
// enables filter policy, correlationID and other configuration for logger.
// Logging is enabled at Info Level and above.
//
// For further logging function. please refer to: https://pkg.go.dev/go.uber.org/zap
//
// Example:
// Create a new logger with name "tenant-service", and filter policy
// to redact value from "email" fields.
// logger.NewLog("tenant-service", logger.WithFilters(filter.Field("email")))
func NewLog(name string, options ...Option) *Logger {
	result := &Logger{}
	for _, opt := range options {
		opt(result)
	}

	result.Logger = result.newZapLogger(name)

	return result
}

func (l *Logger) WithLogger(ctx echo.Context) echo.Context {
	ctx.Set("logger", l)

	return ctx
}

func (l *Logger) FromContext(ctx echo.Context) *Logger {
	if logger, ok := ctx.Get("logger").(*Logger); ok {
		return logger
	}

	return &Logger{}
}

func (l *Logger) newZapLogger(name string) *zap.Logger {
	zapConfig := zap.NewProductionConfig()
	if l.isDevelopment {
		zapConfig.Development = true
		zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	zapConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	if l.encodeTime != nil {
		zapConfig.EncoderConfig.EncodeTime = l.encodeTime
	}

	zapConfig.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	if l.encodeDuration != nil {
		zapConfig.EncoderConfig.EncodeDuration = l.encodeDuration
	}

	if l.level != 0 {
		zapConfig.Level = zap.NewAtomicLevelAt(l.level)
	}

	zapLogger, err := zapConfig.Build()
	if err != nil {
		log.Println(err)
		return nil
	}

	defer func() {
		_ = zapLogger.Sync()
	}()

	zapLogger = zapLogger.Named(name)
	if l.correlationID != "" {
		zapLogger = zapLogger.With(zap.String(FieldCorrelationID, l.correlationID))
	}

	return zapLogger
}

func (l *Logger) clone() *Logger {
	clonedLogger := *l
	return &clonedLogger
}

func (l *Logger) WithCid(cid string) *Logger {
	if cid == "" {
		cid = uuid.New().String()
	}

	cloned := l.clone()
	cloned.correlationID = cid
	cloned.Logger = cloned.With(zap.String(FieldCorrelationID, cid))

	return cloned
}

func (l *Logger) WithErr(err error) *Logger {
	if err == nil {
		return l
	}

	cloned := l.clone()
	errFields := make([]zap.Field, 0)
	errFields = append(errFields, zap.String(FieldError, err.Error()))

	cloned.Logger = cloned.With(errFields...)

	return cloned
}

func (l *Logger) WithFields(fields map[string]any) *Logger {
	cloned := l.clone()
	logFields := make([]zap.Field, 0)

	for key, value := range fields {
		logFields = append(logFields, zap.Any(key, value))
	}

	cloned.Logger = cloned.With(logFields...)

	return cloned
}

func (l *Logger) GetCid() string {
	return l.correlationID
}
