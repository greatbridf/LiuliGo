package liuli

import (
	"bytes"

	"crypto/md5"
	"encoding/hex"
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

func Hash(input []byte) string {
	m := md5.New()
	m.Write(input)
	output := hex.EncodeToString(m.Sum(nil))
	return output
}
