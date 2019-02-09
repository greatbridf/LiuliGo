package liuli

import (
	"fmt"

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
	Log = logger.New(filename, "LiuliGo")
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
