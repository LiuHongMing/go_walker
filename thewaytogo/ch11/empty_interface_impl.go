package main

import (
	"log"
	"os"
)

type Task struct {
	Command string
	*log.Logger
}

func NewTask(command string, logger *log.Logger) *Task {
	return &Task{command, logger}
}

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	t := NewTask("run", logger)
	t.Fatalf("fatal message ...")
}