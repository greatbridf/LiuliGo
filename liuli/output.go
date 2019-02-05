package liuli

import (
    "os"
    "fmt"
)

var filename = "LiuliGo.log"

// PrintError Print error message to stdout
func PrintError(msg string) {
    file, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
    file.WriteString("[ERROR] " + msg + "\n")
    file.Close()
	fmt.Printf("Content-Type: text/plain; charset=utf-8\n\n")
	fmt.Println("Error:", msg)
}

func PrintDebug(msg string) {
    file, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
    file.WriteString("[DEBUG] " + msg + "\n")
    file.Close()
}

