package util

import "fmt"

func PrintError(message string) {
	fmt.Printf(" ❌️ ERROR: %v\n", message)
}

func PrintSuccess(message string) {
	fmt.Printf("✅️ SUCCESS: %v\n", message)
}

func PrintWarning(message string) {
	fmt.Printf(" ‼️ WARNING: %v\n", message)
}
