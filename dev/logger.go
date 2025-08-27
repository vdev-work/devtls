package dev

import (
	"fmt"
	"log"
)

func Info(msg string, a ...string) {
	fmt.Printf("[INFO] - %s\n", fmt.Sprintf(msg, a))
}

func Warning(msg string, a ...string) {
	fmt.Printf("[WARN] - %s\n", fmt.Sprintf(msg, a))
}

func Error(msg string, a ...string) {
	fmt.Printf("[ERROR] - %s\n", fmt.Sprintf(msg, a))
}

func LogInfo(msg string, a ...string) {
	log.Printf("[INFO] - %s\n", fmt.Sprintf(msg, a))
}

func LogWarning(msg string, a ...string) {
	log.Printf("[WARN] - %s\n", fmt.Sprintf(msg, a))
}

func LogError(msg string, a ...string) {
	log.Printf("[ERROR] - %s\n", fmt.Sprintf(msg, a))
}
