package liuli

import "fmt"

// PrintError Print error message to stdout
func PrintError(msg string) {
	fmt.Printf("Content-Type: text/plain; charset=utf-8\n\n")
	fmt.Println("Error:", msg)
}
