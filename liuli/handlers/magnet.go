package handlers

import (
	"github.com/greatbridf/LiuliGo/liuli"
	"github.com/pkg/errors"
	"net/http"
)

func HandleMagnet(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tmp := req.Form["id"]
	if len(tmp) == 0 {
		liuli.PrintError(w, errors.New("no content id provided"))
		return
	}
	id := tmp[0]
	magnet, err := liuli.GetMagnet(id)
	if err != nil {
		liuli.PrintError(w, err)
		return
	}
	liuli.LiuliResp{
		200,
		"OK",
		&liuli.LiuliData{
			nil,
			magnet,
		},
	}.Print(w)
}
