package logger

import loggerlog "github.com/upikoth/starter-go/internal/pkg/logger/logger-log"

type Logger interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)
}

func New() Logger {
	return loggerlog.New()
}
