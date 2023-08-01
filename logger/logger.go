package logger

import (
	"fmt"
	"io"
)

type Logger interface {
	Printf(string, ...interface{})
	FPrintln(w io.Writer, a ...any) (n int, err error)
}

type DefaultLogger struct{}

func (d *DefaultLogger) Printf(msg string, args ...interface{}) {
	fmt.Printf(msg+"\n", args...)
}

func (d *DefaultLogger) FPrintln(w io.Writer, a ...any) (n int, err error) {
	fmt.Fprintln(w, a...)
	return
}

var (
	Log      Logger = &DefaultLogger{}
	ErrorLog Logger = &DefaultLogger{}
)

type LogFunc func(string, ...interface{})

func (f LogFunc) Printf(msg string, args ...interface{}) { f(msg, args...) }

func SetLogger(logger Logger) {
	Log = logger
}

func SetErrorLogger(logger Logger) {
	ErrorLog = logger
}
