package loggerlog

import "log"

type LoggerLog struct{}

func New() *LoggerLog {
	return &LoggerLog{}
}

func (l *LoggerLog) Debug(args ...any) {
	log.Println(args...)
}

func (l *LoggerLog) Info(args ...any) {
	log.Println(args...)
}

func (l *LoggerLog) Warn(args ...any) {
	log.Println(args...)
}

func (l *LoggerLog) Error(args ...any) {
	log.Println(args...)
}

func (l *LoggerLog) Fatal(args ...any) {
	log.Fatal(args...)
}
