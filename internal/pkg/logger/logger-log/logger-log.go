package loggerlog

import "log"

type LoggerLog struct{}

func New() *LoggerLog {
	return &LoggerLog{}
}

func (l *LoggerLog) Debug(msg string) {
	log.Println(msg)
}

func (l *LoggerLog) Info(msg string) {
	log.Println(msg)
}

func (l *LoggerLog) Warn(msg string) {
	log.Println(msg)
}

func (l *LoggerLog) Error(msg string) {
	log.Println(msg)
}

func (l *LoggerLog) Fatal(msg string) {
	log.Fatal(msg)
}
