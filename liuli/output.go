package liuli

import (
	"errors"
	"os"

	logger "github.com/greatbridf/go-logger"
)

var (
	filename string = "LiuliGo.log"
	log      *logger.Logger
)

func init() {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(errors.New("Cannot open log file"))
	}
	log = logger.New(file, "LiuliGo")
}

// PrintError Print error message to stdout
func PrintError(msg string) {
	log.E(msg)
}

func PrintDebug(msg string) {
	log.D(msg)
}
