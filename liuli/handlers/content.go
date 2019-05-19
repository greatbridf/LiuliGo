package handlers

import (
	"github.com/greatbridf/LiuliGo/liuli"
	"github.com/pkg/errors"
	"net/http"
)

func HandleContent(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tmp := req.Form["id"]
	if len(tmp) == 0 {
		w.WriteHeader(400)
		liuli.PrintError(w, errors.New("no content id provided"))
		return
	}
	id := tmp[0]
	content, err := liuli.GetContent(id)
	if err != nil {
		w.WriteHeader(400)
		liuli.PrintError(w, err)
		return
	}
	w.Write([]byte(content))
}
