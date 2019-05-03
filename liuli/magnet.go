package liuli

import (
	"errors"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

type Magnets []string

// GetMagnet get magnet link from content
// Returns array
func GetMagnet(id string) (Magnets, error) {
	link := "https://interface.greatbridf.top/liuli?req=content&id=" + id
	doc, _ := goquery.NewDocument(link)
	doc.Find("a").Remove()
	doc.Find("img").Remove()
	content, _ := doc.Html()
	exp := regexp.MustCompile("[a-zA-Z0-9]{40}|[a-zA-Z0-9]{32}")
	magnet := exp.FindAllString(content, -1)
	if len(magnet) == 0 {
		return nil, errors.New("No magnet links found")
	}
	for i := 0; i < len(magnet); i++ {
		magnet[i] = "magnet:?xt=urn:btih:" + magnet[i]
	}
	return magnet, nil
}
