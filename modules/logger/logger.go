package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger() *Logger {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return &Logger{zapLogger.Sugar()}
}
