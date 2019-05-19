package liuli

import (
	"encoding/json"
	"io"

	logger "github.com/greatbridf/go-logger"
)

type LiuliResp struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data *LiuliData `json:"data,omitempty"`
}

type LiuliData struct {
	Articles Articles `json:"articles,omitempty"`
	Magnets  Magnets  `json:"magnets,omitempty"`
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
func PrintError(w io.Writer, e error) {
	Log.Err(e)
	resp := LiuliResp{
		400,
		e.Error(),
		nil,
	}
	json, _ := json.Marshal(resp)
	w.Write(json)
}

func (resp LiuliResp) Print(w io.Writer) {
	json, _ := json.Marshal(resp)
	w.Write(json)
}

func PrintDebug(msg string) {
	Log.D(msg)
}
