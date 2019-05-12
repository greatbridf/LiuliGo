package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/greatbridf/LiuliGo/liuli"
)

func main() {
	uri, _ := url.Parse(os.Getenv("REQUEST_URI"))
	query, _ := url.ParseQuery(uri.RawQuery)
	req := query.Get("req")
	switch req {
	case "articles":
		page := query.Get("page")
		if page == "" {
			page = "1"
		}
		articles, err := liuli.GetArticles(page)
		if err != nil {
			liuli.PrintError(err.Error())
			return
		}
		err = articles.Print()
		if err != nil {
			liuli.PrintError(err.Error())
			return
		}
		break
	case "content":
		id := query.Get("id")
		if id == "" {
			liuli.PrintError("No content id given")
			return
		}
		content, err := liuli.GetContent(id)
		if err != nil {
			liuli.PrintError(err.Error())
		} else {
			fmt.Printf("Content-Type: text/html; charset=utf-8\n\n")
			fmt.Println(content)
		}
		break
	case "magnet":
		id := query.Get("id")
		if id == "" {
			liuli.PrintError("No content id given")
			return
		}
		magnet, err := liuli.GetMagnet(id)
		if err != nil {
			liuli.PrintError(err.Error())
			return
		}
		err = magnet.Print()
		if err != nil {
			liuli.PrintError(err.Error())
		}
		break
	default:
		liuli.PrintError("Invalid query method")
		break
	}
}
