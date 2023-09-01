package Lestrade

import (
	"fmt"
	"log"
)

func LogInfo(message string) {
	log.Printf("INFO - %v", message)
}

func LogWarning(message string) {
	log.Printf("WARN - %v", message)
}

func LogError(message string) {
	log.Printf("ERROR - %v", message)
}

func PrintRed(message string) {
	fmt.Println("\033[1;31m" + message + "\033[0m")
}
