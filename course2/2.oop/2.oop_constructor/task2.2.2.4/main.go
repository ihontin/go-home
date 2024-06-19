package main

import (
	"fmt"
	"os"
)

// Logger interface
type Logger interface {
	Log(message string)
}

type FileLogger struct {
	file *os.File
}

func (f FileLogger) Log(message string) {
	_, _ = fmt.Fprintln(f.file, message)
}

type LogSystem struct {
	Logger Logger
}

func (l LogSystem) Log(message string) {
	l.Logger.Log(message)
}

// LogOption functional option type
type LogOption func(*LogSystem)

func NewLogSystem(option LogOption) *LogSystem {
	out := &LogSystem{}
	option(out)
	return out
}

func WithLogger(l Logger) LogOption {
	return func(s *LogSystem) {
		s.Logger = l
	}
}

func main() {
	file, _ := os.Create("log.txt")
	defer file.Close()

	fileLogger := FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(fileLogger))

	logSystem.Log("Hello, world!")
	fmt.Println("all wright")
}
