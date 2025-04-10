package advanced_level

import "fmt"

type EventListener interface {
	OnEvent(data string)
}

type LoggerListener struct{}

func (l *LoggerListener) OnEvent(data string) {
	fmt.Printf("a log was issued with: %s\n", data)
}

type EmailListener struct{}

func (l *EmailListener) OnEvent(data string) {
	fmt.Printf("an email was issued with: %s\n", data)
}

func ListenSomething(e EventListener) {
	e.OnEvent("data")
}
