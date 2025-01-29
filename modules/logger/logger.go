package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
	With(fields ...Field) Logger
}

type Field struct {
	Key   string
	Value interface{}
}

type logger struct {
	zap *zap.Logger
}

type Params struct {
	fx.In

	Environment string `name:"environment"`
}

func NewLogger(p Params) (Logger, error) {
	config := zap.NewProductionConfig()

	// Set log level based on environment
	switch p.Environment {
	case "development":
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "staging":
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	default:
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	// Configure encoding
	config.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Create logger
	zapLogger, err := config.Build(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	if err != nil {
		return nil, err
	}

	return &logger{
		zap: zapLogger,
	}, nil
}

func (l *logger) Debug(msg string, fields ...Field) {
	l.zap.Debug(msg, l.convertFields(fields...)...)
}

func (l *logger) Info(msg string, fields ...Field) {
	l.zap.Info(msg, l.convertFields(fields...)...)
}

func (l *logger) Warn(msg string, fields ...Field) {
	l.zap.Warn(msg, l.convertFields(fields...)...)
}

func (l *logger) Error(msg string, fields ...Field) {
	l.zap.Error(msg, l.convertFields(fields...)...)
}

func (l *logger) Fatal(msg string, fields ...Field) {
	l.zap.Fatal(msg, l.convertFields(fields...)...)
}

func (l *logger) With(fields ...Field) Logger {
	return &logger{
		zap: l.zap.With(l.convertFields(fields...)...),
	}
}

func (l *logger) convertFields(fields ...Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, f := range fields {
		zapFields[i] = zap.Any(f.Key, f.Value)
	}
	return zapFields
}
