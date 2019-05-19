package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/greatbridf/LiuliGo/liuli"
	"github.com/pkg/errors"
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
			liuli.PrintError(err)
			return
		}
		articles.Print()
		break
	case "content":
		id := query.Get("id")
		if id == "" {
			liuli.PrintError(errors.New("No content id given"))
			return
		}
		content, err := liuli.GetContent(id)
		if err != nil {
			liuli.PrintError(err)
		} else {
			fmt.Printf("Content-Type: text/html; charset=utf-8\n\n")
			fmt.Println(content)
		}
		break
	case "magnet":
		id := query.Get("id")
		if id == "" {
			liuli.PrintError(errors.New("No content id given"))
			return
		}
		magnet, err := liuli.GetMagnet(id)
		if err != nil {
			liuli.PrintError(err)
			return
		}
		magnet.Print()
		break
	case "delete":
		id := query.Get("id")
		if id == "" {
			liuli.PrintError(errors.New("No content id given"))
			return
		}
		result, err := liuli.DeleteResource(id)
		if err != nil {
			liuli.PrintError(err)
			return
		}
		result.Print()
		break
	default:
		liuli.PrintError(errors.New("Invalid query method"))
		break
	}
}
