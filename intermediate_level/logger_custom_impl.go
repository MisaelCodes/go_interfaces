package intermediate_level

import "fmt"

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (cl ConsoleLogger) Log(message string) {
	fmt.Printf("[Console]: %s\n", message)
}

type FileLogger struct{}

func (fl FileLogger) Log(message string) {
	fmt.Printf("[File]: %s\n", message)
}

func LoggingStuff(l Logger, message string) {
	l.Log(message)
}
