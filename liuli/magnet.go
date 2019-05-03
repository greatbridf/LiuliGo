package liuli

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

type Magnets []string

// GetMagnet get magnet link from content
// Returns array
func GetMagnet(id string) Magnets {
	link := "https://interface.greatbridf.top/liuli?req=content&id=" + id
	doc, _ := goquery.NewDocument(link)
	doc.Find("a").Remove()
	doc.Find("img").Remove()
	content, _ := doc.Html()
	exp := regexp.MustCompile("[a-zA-Z0-9]{40}|[a-zA-Z0-9]{32}")
	magnet := exp.FindAllString(content, -1)
	for i := 0; i < len(magnet); i++ {
		magnet[i] = "magnet:?xt=urn:btih:" + magnet[i]
	}
	return magnet
}
