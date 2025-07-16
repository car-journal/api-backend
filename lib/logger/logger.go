// Package logger handles logging system
package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var LoggerInterface Interface

type Interface interface {
	Log(logMessage string)
}

type Logger struct {
	featureFlag bool
}

func Init(featureFlag bool) {
	LoggerInterface = &Logger{
		featureFlag: featureFlag,
	}
}

func (l Logger) Log(logMessage string) {
	if !l.featureFlag {
		return
	}

	now := time.Now()
	fileName := fmt.Sprintf("%s.log", now.Format("2006-01-02"))
	fileDirectory := fmt.Sprintf("logs/%s", fileName)

	logFile, err := os.OpenFile(fileDirectory, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()

	// Create a logger that writes to the file
	timestamp := now.Format("2006-01-02 15:04:05 -0700")
	logger := log.New(logFile, "", 0)
	logger.Printf("%s %s", timestamp, logMessage)
}
