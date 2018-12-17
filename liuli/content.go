package liuli

import (
	"github.com/PuerkitoBio/goquery"
)

// GetContent get page by link
// Both content and styles
func GetContent(id string) string {
	link := "https://www.hacg.me/wp/" + id + ".html"
	doc, err := goquery.NewDocument(link)
	if err != nil {
		panic(err)
	}

	content := GetContentNoStyle(doc)
	if content == "" {
		panic("No content")
	}
	style := GetStyle(doc)

	return style + content
}

// GetContentNoStyle get content from doc
func GetContentNoStyle(doc *goquery.Document) string {
	content := ""
	doc.Find(".entry-content").Each(func(_ int, selection *goquery.Selection) {
		tmp, err := selection.Html()
		if err != nil {
			panic(err)
		}
		content += tmp
	})
	return content
}

// GetStyle get css tags from doc
func GetStyle(doc *goquery.Document) string {
	style := ""
	doc.Find("link[rel='stylesheet']").Each(func(_ int, selection *goquery.Selection) {
		tmp := RenderHTMLTag(selection)
		style += (tmp + "\n")
	})
	return style
}