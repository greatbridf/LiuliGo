package main

import (
	"fmt"
	"net/url"
	"os"

	"./liuli"
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
		articles := liuli.GetArticles(page)
		fmt.Printf("Content-Type: application/json; charset=utf-8\n\n")
		fmt.Println(liuli.GetArticlesJSON(articles))
		break
	case "content":
		id := query.Get("id")
		if id == "" {
			liuli.PrintError("No content id given")
			return
		}
		content := liuli.GetContent(id)
		if content == "" {
			liuli.PrintError("Unable to get content")
			return
		}
		fmt.Printf("Content-Type: text/html; charset=utf-8\n\n")
		fmt.Println(content)
		break
	case "magnet":
		id := query.Get("id")
		if id == "" {
			liuli.PrintError("No content id given")
			return
		}
		magnet := liuli.GetMagnet(id)
		if len(magnet) == 0 {
			liuli.PrintError("No magnet link found in " + id)
		}
		fmt.Printf("Content-Type: text/plain; charset=utf-8\n\n")
    for i := 0; i < len(magnet); i++ {
      fmt.Println(magnet[i])
    }
		break
	default:
		liuli.PrintError("Invalid query method")
		break
	}
}
