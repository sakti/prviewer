package main

import (
	"fmt"
	"time"

	"github.com/andlabs/ui"
)

var (
	logBuffer = make(chan string, 100)
)

func processLog() {
	for {
		logLine := <-logBuffer
		if entryLog == nil {
			time.Sleep(1 * time.Second)
		}
		LogWrite(logLine)
	}
}

func LogWrite(content string) {
	text := entryLog.Text()
	text += content
	// Check: No append text api
	ui.QueueMain(func() {
		entryLog.SetText(text)
	})
}

func Info(content string) {
	now := time.Now().Format(time.RFC3339)
	logBuffer <- fmt.Sprintf("%s info: %s\n", now, content)
}

func Infof(format string, a ...interface{}) {
	now := time.Now().Format(time.RFC3339)
	content := fmt.Sprintf(format, a...)
	logBuffer <- fmt.Sprintf("%s info: %s\n", now, content)
}

func Error(content string) {
	now := time.Now().Format(time.RFC3339)
	logBuffer <- fmt.Sprintf("%s error: %s\n", now, content)
}

func Errorf(format string, a ...interface{}) {
	now := time.Now().Format(time.RFC3339)
	content := fmt.Sprintf(format, a...)
	logBuffer <- fmt.Sprintf("%s error: %s\n", now, content)
}
