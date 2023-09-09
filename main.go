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
// PrintRed prints text in red color.
func PrintRed(message string) {
	fmt.Println("\033[1;31m" + message + "\033[0m")
}

// PrintOrange prints text in orange color.
func PrintOrange(message string) {
	fmt.Println("\033[1;38;5;208m" + message + "\033[0m")
}

// PrintYellow prints text in yellow color.
func PrintYellow(message string) {
	fmt.Println("\033[1;33m" + message + "\033[0m")
}

// PrintGreen prints text in green color.
func PrintGreen(message string) {
	fmt.Println("\033[1;32m" + message + "\033[0m")
}

// PrintBlue prints text in blue color.
func PrintBlue(message string) {
	fmt.Println("\033[1;34m" + message + "\033[0m")
}

// PrintPurple prints text in purple color.
func PrintPurple(message string) {
	fmt.Println("\033[1;35m" + message + "\033[0m")
}

// PrintRainbow prints text in a rainbow of colors.
func PrintRainbow(message string) {
	PrintRed(message)
	PrintOrange(message)
	PrintYellow(message)
	PrintGreen(message)
	PrintBlue(message)
	PrintPurple(message)
}
