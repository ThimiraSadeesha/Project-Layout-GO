package logger

import (
	"fmt"
	"log"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func Info(format string) {
	message := fmt.Sprintf(format)
	log.Println(Green + "INFO: " + message)
}

func Warning(format string) {
	message := fmt.Sprintf(format)
	log.Println(Yellow + "WARNING: " + message)
}

func Error(format string) {
	message := fmt.Sprintf(format)
	log.Println(Red + "ERROR: " + message)
}

func Debug(format string) {
	message := fmt.Sprintf(format)
	log.Println(Blue + "DEBUG: " + message)
}
