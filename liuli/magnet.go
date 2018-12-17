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
	magnet := exp.FindAllString(content, -1)
	for i := 0; i < len(magnet); i++ {
		magnet[i] = "magnet:?xt=urn:btih:" + magnet[i]
	}
	return magnet
}
