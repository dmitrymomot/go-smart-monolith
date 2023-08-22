package logx

import "log"

type Logger struct{}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Error(err error, kv ...interface{}) {
	log.Println(err, kv)
}

func (l *Logger) Fatal(err error) {
	log.Fatalln(err)
}
