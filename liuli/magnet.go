package liuli

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// GetMagnet get magnet link from content
// Returns array
func GetMagnet(id string) []string {
	link := "https://www.hacg.me/wp/" + id + ".html"
	doc, err := goquery.NewDocument(link)
	if err != nil {
		panic(err)
	}
	content := GetContentNoStyle(doc)
	exp := regexp.MustCompile("[a-zA-Z0-9]{40}")
	return exp.FindAllString(content, -1)
}
