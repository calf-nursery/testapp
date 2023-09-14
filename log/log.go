package log

import "fmt"

type Logger interface {
	Info(message string)
	Debug(message string)
}

func New() Logger {
	return &fmtLogger{}
}

type fmtLogger struct{}

func (f *fmtLogger) Info(message string) {
	fmt.Printf("INFO - %s\n", message)
}

func (f *fmtLogger) Debug(message string) {
	fmt.Printf("DEBUG - %s\n", message)
}
