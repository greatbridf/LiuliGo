package main

import (
	"fmt"
	"io/ioutil"
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
		} else {
			fmt.Printf("Content-Type: application/json; charset=utf-8\n\n")
			fmt.Println(liuli.GetArticlesJSON(articles))
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
		magnet := liuli.GetMagnet(id)
		fmt.Printf("Content-Type: text/plain; charset=utf-8\n\n")
		for i := 0; i < len(magnet); i++ {
			fmt.Println(magnet[i])
		}
		break
	case "resource":
		hash := query.Get("hash")
		data, err := liuli.GetResource(hash)
		if err != nil {
			liuli.PrintError(err.Error())
		} else {
			fmt.Println("Content-Length: " + fmt.Sprintf("%d", len(data)))
			fmt.Printf("Content-Type: image/jpeg\n\n")
			err := ioutil.WriteFile("/dev/stdout", data, 0666)
			if err != nil {
				liuli.Log.E(err.Error())
			}
		}
		break
	default:
		liuli.PrintError("Invalid query method")
		break
	}
}
