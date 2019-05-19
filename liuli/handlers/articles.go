package handlers

import (
	"net/http"

	"github.com/greatbridf/LiuliGo/liuli"
	"github.com/pkg/errors"
)

func HandleArticles(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	tmp := req.Form["id"]
	page := "1"
	if len(tmp) != 0 {
		page = tmp[0]
	}
	articles, err := liuli.GetArticles(page)
	if err != nil {
		w.WriteHeader(400)
		liuli.PrintError(w, err)
		return
	}
	liuli.LiuliResp{
		200,
		"OK",
		&liuli.LiuliData{
			articles,
			nil,
		},
	}.Print(w)
}
