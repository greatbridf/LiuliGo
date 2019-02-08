package liuli

import (
	"errors"
	"fmt"
	"os"

	logger "github.com/greatbridf/go-logger"
)

var (
	filename string = "LiuliGo.log"
	Log      *logger.Logger
)

const (
	ERR_CANNOT_GOQUERY = "Cannot get goquery documents"
)

func init() {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(errors.New("Cannot open log file"))
	}
	Log = logger.New(file, "LiuliGo")
}

// PrintError Print error message to stdout
func PrintError(msg string) {
	Log.E(msg)
	fmt.Printf("Content-Type: text/plain; charset=utf-8\n\n")
	fmt.Println(msg)
}

func PrintDebug(msg string) {
	Log.D(msg)
}
