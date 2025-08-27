package dev

import (
	"fmt"
	"log"
)

func Info(msg string, a ...string) {
	fmt.Printf("[INFO] - %s", fmt.Sprintf(msg, a))
}

func Warning(msg string, a ...string) {
	fmt.Printf("[WARN] - %s", fmt.Sprintf(msg, a))
}

func Error(msg string, a ...string) {
	fmt.Printf("[ERROR] - %s", fmt.Sprintf(msg, a))
}

func LogInfo(msg string, a ...string) {
	log.Printf("[INFO] - %s", fmt.Sprintf(msg, a))
}

func LogWarning(msg string, a ...string) {
	log.Printf("[WARN] - %s", fmt.Sprintf(msg, a))
}

func LogError(msg string, a ...string) {
	log.Printf("[ERROR] - %s", fmt.Sprintf(msg, a))
}
