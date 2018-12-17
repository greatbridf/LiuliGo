package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	"./liuli"
)

func main() {
	uri, _ := url.Parse(os.Getenv("REQUEST_URI"))
	query, _ := url.ParseQuery(uri.RawQuery)
	req := query.Get("req")
	switch req {
	case "articles":
		fmt.Printf("Content-Type: application/json; charset=utf-8\n\n")
		page, err := strconv.Atoi(query.Get("page"))
		if err != nil {
			page = 1
		}
		articles := liuli.GetArticles(page)
		fmt.Println(liuli.GetArticlesJSON(articles))
		break
	case "content":
		fmt.Printf("Content-Type: text/plain; charset=utf-8\n\n")
		fmt.Printf("OK!")
		break
	case "magnet":
		fmt.Printf("Content-Type: text/plain; charset=utf-8\n\n")
		fmt.Printf("OK!")
		break
	default:
		fmt.Printf("Content-Type: text/plain; charset=utf-8\n\n")
		fmt.Printf("NO!")
		break
	}
}
