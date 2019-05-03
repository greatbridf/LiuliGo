package liuli

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"

	logger "github.com/greatbridf/go-logger"
)

type Strings []string

type LiuliResp struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data LiuliData `json:"data"`
}

type LiuliData struct {
	Articles Articles `json:"articles,omitempty"`
	Magnets  Strings  `json:"magnets,omitempty"`
}

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
	fmt.Printf("Content-Type: application/json; charset=utf-8\n\n")
	resp := LiuliResp{
		400,
		msg,
		LiuliData{},
	}
	json, err := respStringify(resp)
	if err != nil {
		Log.E(err.Error())
		return
	}
	fmt.Println(json)
}

func (data Articles) Print() error {
	resp := LiuliResp{
		200,
		"OK",
		LiuliData{
			data,
			nil,
		},
	}
	json, err := respStringify(resp)
	if err != nil {
		Log.E(err.Error())
		return errors.Wrap(err, "cannot print data")
	}
	fmt.Printf("Content-Type: application/json; charset=utf-8\n\n")
	fmt.Println(json)
	return nil
}

func (data Strings) Print() error {
	resp := LiuliResp{
		200,
		"OK",
		LiuliData{
			nil,
			data,
		},
	}
	json, err := respStringify(resp)
	if err != nil {
		Log.E(err.Error())
		return errors.Wrap(err, "cannot print data")
	}
	fmt.Printf("Content-Type: application/json; charset=utf-8\n\n")
	fmt.Println(json)
	return nil
}

func respStringify(resp LiuliResp) (string, error) {
	out, err := json.Marshal(resp)
	if err != nil {
		Log.E(err.Error())
		return "", errors.Wrap(err, "cannot stringify json")
	}
	return string(out), nil
}

func PrintDebug(msg string) {
	Log.D(msg)
}
