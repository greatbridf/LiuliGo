package liuli

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// GetMagnet get magnet link from content
// Returns array
func GetMagnet(id string) []string {
	link := "https://greatbridf.top/interface/LiuliGo.cgi?req=content&id=" + id
	doc, _ := goquery.NewDocument(link)
	doc.Find("a").Remove()
	doc.Find("img").Remove()
	content, _ := doc.Html()
	exp := regexp.MustCompile("[a-zA-Z0-9]{40}")
	magnet := exp.FindAllString(content, -1)
	for i := 0; i < len(magnet); i++ {
		magnet[i] = "magnet:?xt=urn:btih:" + magnet[i]
	}
	return magnet
}
