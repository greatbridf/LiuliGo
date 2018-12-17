package liuli

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// RenderHTMLTag convert goquery.Selection to html string
func RenderHTMLTag(selection *goquery.Selection) string {
	node := selection.Get(0)
	buf := bytes.NewBuffer([]byte{})
	html.Render(buf, node)
	return buf.String()
}
