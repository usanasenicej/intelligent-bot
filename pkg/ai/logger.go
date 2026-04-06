package ai

import "fmt"

// Logger is a basic logging interface for the agent.
type Logger interface {
	Log(msg string)
	Errorf(format string, args ...interface{})
}

// DefaultLogger is a simple stdout logger.
type DefaultLogger struct{}

func (d DefaultLogger) Log(msg string) {
	fmt.Printf("[LOG]: %s\n", msg)
}

func (d DefaultLogger) Errorf(format string, args ...interface{}) {
	fmt.Printf("[ERR]: "+format+"\n", args...)
}
