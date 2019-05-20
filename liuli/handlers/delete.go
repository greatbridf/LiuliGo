package handlers

import (
	"github.com/greatbridf/LiuliGo/liuli"
	"github.com/pkg/errors"
	"net/http"
)

func HandleDelete(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tmp := req.Form["id"]
	if len(tmp) == 0 {
		liuli.PrintError(w, errors.New("no resource id provided"))
		return
	}
	id := tmp[0]
	result, err := liuli.DeleteResource(id)
	if err != nil {
		liuli.PrintError(w, err)
		return
	}
	liuli.LiuliResp{
		result.Code,
		result.Msg,
		nil,
	}.Print(w)
}
