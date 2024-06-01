package logger

import (
	loggerzerolog "github.com/upikoth/starter-go/internal/pkg/logger/logger-zerolog"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
}

func New() Logger {
	return loggerzerolog.New()
}
