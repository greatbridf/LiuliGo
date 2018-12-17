package liuli

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// GetContent get page by link
func GetContent(id string) string {
	link := "https://www.hacg.me/wp/" + id + ".html"
	res, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic(res.StatusCode)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	content, style := "", ""
	doc.Find(".entry-content").Each(func(_ int, selection *goquery.Selection) {
		tmp, err := selection.Html()
		if err != nil {
			panic(err)
		}
		content += tmp
	})
	if content == "" {
		return content
	}

	doc.Find("link[rel='stylesheet']").Each(func(_ int, selection *goquery.Selection) {
		tmp := RenderHTMLTag(selection)
		style += (tmp + "\n")
	})
	return style + content
}
