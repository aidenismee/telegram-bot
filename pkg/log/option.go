package log

import (
	"github.com/google/uuid"
	"go.uber.org/zap/zapcore"
)

type Option func(*Logger)

func WithCid(cid string) Option {
	return func(l *Logger) {
		if cid == "" {
			cid = uuid.New().String()
		}

		l.correlationID = cid
	}
}

func WithTimeEncoder(encoder zapcore.TimeEncoder) Option {
	return func(l *Logger) {
		l.encodeTime = encoder
	}
}

func WithDurationEncoder(encoder zapcore.DurationEncoder) Option {
	return func(l *Logger) {
		l.encodeDuration = encoder
	}
}

func WithDevelopment(isDevelop bool) Option {
	return func(l *Logger) {
		l.isDevelopment = isDevelop
	}
}

func WithLevel(level string) Option {
	logLevel := zapcore.InfoLevel

	if lvl, exist := Levels[level]; exist {
		logLevel = lvl
	}

	return func(l *Logger) {
		l.level = logLevel
	}
}
